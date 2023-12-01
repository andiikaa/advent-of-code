import org.javatuples.Pair;
import org.jgrapht.Graph;
import org.jgrapht.GraphPath;
import org.jgrapht.alg.shortestpath.DijkstraShortestPath;
import org.jgrapht.graph.DefaultUndirectedWeightedGraph;
import org.jgrapht.graph.DefaultWeightedEdge;
import org.jgrapht.nio.Attribute;
import org.jgrapht.nio.DefaultAttribute;
import org.jgrapht.nio.dot.DOTExporter;

import java.io.FileWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Set;
import java.util.stream.Collectors;

public class Main16 {

	public static void main(String[] args) throws IOException {
		new Main16().run();
	}

	int maxFlowRate = 0;
	int maxPathSize = 0;
	Map<String, Integer> distanceMap;
	DefaultUndirectedWeightedGraph<Valve, ExtendedWeightedEdge> graph;
	List<Valve> remainingValves;

	void run() throws IOException {
		String name = "input";
		List<String> data = Files.readAllLines(Paths.get("andre/d16/" + name + ".txt"));
		Map<String, Valve> valves = data.stream()
				.map(this::toValve)
				.collect(HashMap::new, (map, valve) -> map.putIfAbsent(valve.name, valve), HashMap::putAll);
		data.forEach(s -> addTunnels(s, valves));
		graph = toGraph(valves);
		maxFlowRate = valves.values().stream().mapToInt(v -> v.flowRate).sum();
		Valve start = valves.get("AA");

//		printGraph(undirectedGraph, name + ".dot");
		simplifyGraph(graph, start);
//		printGraph(undirectedGraph, name + "-simple.dot");

		distanceMap = calculateDistances(graph);
		remainingValves = graph.vertexSet().stream().filter(v -> !v.equals(start)).toList();
		maxPathSize = remainingValves.size();

		Set<Valve> startValves = new HashSet<>();
		addStartValues(start, startValves, 0, 2);
		startValves.remove(start);

		Set<List<Valve>> paths = dfs2(start, new ArrayList<>(), 1, 0);
		Pair<Integer, List<Valve>> best = findBest(start, paths);

		int max = best.getValue0();
		List<Valve> bestPath = best.getValue1();

		// 2183: [IF, IE, WQ, GU, UN, RQ, BT, CQ, MU]
		// Example: found best with 1651: [DD, BB, JJ, HH, EE, CC]
		System.out.println("found best with " + max + ": " + bestPath);
	}

	Pair<Integer, List<Valve>> findBest(Valve start, Set<List<Valve>> paths) {
		int max = 0;
		List<Valve> bestPath = null;
		for (List<Valve> next : paths) {
			int current = pathMax(start, next, distanceMap);
			if (current > max) {
				max = current;
				bestPath = new ArrayList<>(next);
				System.out.println("found new best with " + max + ": " + bestPath);
			}
		}
		return Pair.with(max, bestPath);
	}

	void addStartValues(Valve root, Set<Valve> startValves, int currentDept, final int maxDept) {
		currentDept++;
		for (var e : graph.edgesOf(root)) {
			var child = e.getSource().equals(root) ? e.getTarget() : e.getSource();
			startValves.add(child);
			if (currentDept < maxDept) {
				addStartValues(child, startValves, currentDept, maxDept);
			} else {
				return;
			}
		}
	}

	int pathMax(Valve start, List<Valve> path, Map<String, Integer> distanceMap) {
		Valve last = start;
		int minutes = 1;
		int sum = 0;
		for (var v : path) {
			String id = last.name + v.name;
			// +1 time needed for opening the valve
			minutes += distanceMap.get(id) + 1;
			if (minutes >= 30) break;
			sum += ((30 - minutes + 1) * v.flowRate);
			last = v;
		}
		return sum;
	}

	Set<List<Valve>> dfs2(final Valve node, final List<Valve> path, int minutes, int flowRate) {
		Set<List<Valve>> allPaths = new HashSet<>();
		if (!node.name.equals("AA") && !path.contains(node)) {
			// open a valve
			minutes++;
			flowRate += node.flowRate;
			path.add(node);
		}

		if (minutes >= 30 || flowRate == maxFlowRate || path.size() >= maxPathSize) {
			allPaths.add(path);
			return allPaths;
		}

		for (var e : graph.edgesOf(node)) {
			var child = e.getSource().equals(node) ? e.getTarget() : e.getSource();
			int distance = e.getWeightAsInt();
			if (path.contains(child)) continue;
			var tmpPath = new ArrayList<>(path);
			allPaths.addAll(dfs2(child, tmpPath, minutes + distance, flowRate));
		}
		return allPaths;
	}

	// clearing all vertices with flowrate 0
	void simplifyGraph(DefaultUndirectedWeightedGraph<Valve, ExtendedWeightedEdge> graph, Valve start) {
		for (var v : new ArrayList<>(graph.vertexSet())) {
			if (v.flowRate > 0 || v.equals(start)) continue;
			Set<ExtendedWeightedEdge> edges = graph.edgesOf(v);
			for (var e : edges) {
				var source = e.getSource().equals(v) ? e.getTarget() : e.getSource();
				for (var e2 : edges) {
					if (e.equals(e2)) continue;
					var target = e2.getSource().equals(v) ? e2.getTarget() : e2.getSource();
					ExtendedWeightedEdge newEdge = graph.addEdge(source, target);
					if (newEdge != null) graph.setEdgeWeight(newEdge, e2.getWeight() + e.getWeight());
				}
			}
			graph.removeVertex(v);
			graph.vertexSet().forEach(o -> o.leadsTo.remove(v));
		}
	}

	Map<String, Integer> calculateDistances(DefaultUndirectedWeightedGraph<Valve, ExtendedWeightedEdge> graph) {
		Map<String, Integer> distances = new HashMap<>();
		DijkstraShortestPath<Valve, ExtendedWeightedEdge> dij = new DijkstraShortestPath<>(graph);
		for (var s : graph.vertexSet()) {
			for (var t : graph.vertexSet()) {
				if (t.equals(s)) continue;
				GraphPath<Valve, ExtendedWeightedEdge> path = dij.getPath(s, t);
				if (path != null) distances.putIfAbsent("" + s + t, (int) path.getWeight());
			}
		}
		return distances;
	}

	void printGraph(Graph<Valve, ExtendedWeightedEdge> g, String name) throws IOException {
		DOTExporter<Valve, ExtendedWeightedEdge> exporter = new DOTExporter<>(v -> v.name);
		exporter.setVertexAttributeProvider(v -> {
			Map<String, Attribute> map = new HashMap<>();
			map.put("label", DefaultAttribute.createAttribute(v.name + "\n" + v.flowRate));
			return map;
		});

		exporter.setEdgeAttributeProvider(v -> {
			Map<String, Attribute> map = new HashMap<>();
			map.put("label", DefaultAttribute.createAttribute(v.getWeightAsInt()));
			return map;
		});
		FileWriter f = new FileWriter(name);
		exporter.exportGraph(g, f);
	}

	DefaultUndirectedWeightedGraph<Valve, ExtendedWeightedEdge> toGraph(Map<String, Valve> valves) {
		DefaultUndirectedWeightedGraph<Valve, ExtendedWeightedEdge> g = new DefaultUndirectedWeightedGraph<>(ExtendedWeightedEdge.class);
		g.setEdgeSupplier(ExtendedWeightedEdge::new);
		for (var v : valves.values()) {
			g.addVertex(v);
			for (var c : v.leadsTo) {
				if (!g.containsVertex(c)) g.addVertex(c);
				ExtendedWeightedEdge edge = g.addEdge(v, c);
				if (edge != null) g.setEdgeWeight(edge, 1);
			}
		}
		return g;
	}

	void addTunnels(String in, Map<String, Valve> valves) {
		String[] split = in.split(";");
		String s = split[0];
		String name = s.split(" ")[1];
		List<Valve> leadsTo = valves.get(name).leadsTo;
		String mapping = split[1];
		String otherValves = mapping.substring(mapping.lastIndexOf("to valve") + 9);
		if (otherValves.contains(",")) {
			Arrays.stream(otherValves.split(","))
					.map(v -> valves.get(v.trim()))
					.forEach(leadsTo::add);
		} else {
			leadsTo.add(valves.get(otherValves.trim()));
		}
	}

	Valve toValve(String in) {
		String s = in.split(";")[0];
		int rate = Integer.parseInt(s.substring(s.lastIndexOf("rate=") + 5));
		String name = s.split(" ")[1];
		return new Valve(name, rate);
	}

	class ExtendedWeightedEdge extends DefaultWeightedEdge {
		@Override
		protected double getWeight() {
			return super.getWeight();
		}

		protected int getWeightAsInt() {
			return (int) super.getWeight();
		}

		@Override
		protected Valve getSource() {
			return (Valve) super.getSource();
		}

		@Override
		protected Valve getTarget() {
			return (Valve) super.getTarget();
		}

	}

	class Valve {
		final String name;
		final int flowRate;

		List<Valve> leadsTo = new ArrayList<>();

		Valve(String name, int flowRate) {
			this.name = name;
			this.flowRate = flowRate;
		}

		@Override
		public boolean equals(Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;
			Valve valve = (Valve) o;
			return flowRate == valve.flowRate && name.equals(valve.name) && leadsTo.equals(valve.leadsTo);
		}

		@Override
		public int hashCode() {
			return Objects.hash(name, flowRate);
		}

		@Override
		public String toString() {
			return name;
		}
	}
}
