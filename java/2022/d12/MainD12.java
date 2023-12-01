import org.javatuples.Pair;
import org.jgrapht.alg.shortestpath.DijkstraShortestPath;
import org.jgrapht.graph.DefaultDirectedWeightedGraph;
import org.jgrapht.graph.DefaultWeightedEdge;

import java.nio.file.Files;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static java.nio.file.Paths.get;

public class MainD12 {

	public static void main(String[] args) throws Exception {
		new MainD12().run();
	}

	List<Integer> pos = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".chars().boxed().toList();
	int start = pos.indexOf((int) 'S');
	int end = pos.indexOf((int) 'E');

	List<List<Integer>> grid;

	void run() throws Exception {
		grid = Files.readAllLines((get("andre/d12/input.txt"))).stream().map(s -> s.chars().boxed().map(pos::indexOf).toList()).toList();
		// x, y
		DefaultDirectedWeightedGraph<Pair<Integer, Integer>, DefaultWeightedEdge> graph = new DefaultDirectedWeightedGraph<>(DefaultWeightedEdge.class);

		for (int i = 0; i < grid.size(); i++) {
			List<Integer> row = grid.get(i);
			for (int k = 0; k < row.size(); k++) {
				Pair<Integer, Integer> from = Pair.with(i, k);
				graph.addVertex(from);

				int raw = row.get(k);
				List<Pair<Integer, Integer>> moves = Arrays.asList(Pair.with(i - 1, k), Pair.with(i + 1, k), Pair.with(i, k - 1), Pair.with(i, k + 1));
				moves = moves.stream().filter(m -> m.getValue0() >= 0 && m.getValue0() < grid.size() && m.getValue1() >= 0 && m.getValue1() < row.size()).toList();

				for (Pair<Integer, Integer> m : moves) {
					int weight = Integer.MAX_VALUE;
					int x1 = m.getValue0();
					int y1 = m.getValue1();
					Pair<Integer, Integer> to = Pair.with(x1, y1);
					graph.addVertex(to);

					int rawNext = grid.get(x1).get(y1);

					raw = raw == start ? pos.indexOf((int) 'a') : raw;
					rawNext = rawNext == end ? pos.indexOf((int) 'z') : rawNext;

					if ((Math.abs(raw - rawNext) <= 1) || (rawNext < raw)) {
						//weight = rawNext;
						weight = 1;
					}

					if (weight != Integer.MAX_VALUE) {
						DefaultWeightedEdge edge = graph.addEdge(from, to);
						if (edge == null) throw new IllegalStateException("edge null");
						graph.setEdgeWeight(edge, weight);
					}

				}
			}
		}

		// example: start 0, 0
		// example: end 2, 5
		// start 20, 0
		// end 20, 149

//		Pair<Integer, Integer> from = Pair.with(0, 0);
//		Pair<Integer, Integer> to = Pair.with(2, 5);
		Pair<Integer, Integer> from = Pair.with(20, 0);
		Pair<Integer, Integer> to = Pair.with(20, 149);

		int shortest = Integer.MAX_VALUE;
		for (var p : findPosOf('a')) {
			var pathBetween = DijkstraShortestPath.findPathBetween(graph, p, to);
			if (pathBetween == null) continue;
			if (pathBetween.getLength() < shortest) {
				shortest = pathBetween.getLength();
			}
		}
		print("points: \n" + shortest);
	}

	List<Pair<Integer, Integer>> findPosOf(char c) {
		int v = pos.indexOf((int) c);
		List<Pair<Integer, Integer>> p = new ArrayList<>();
		for (int i = 0; i < grid.size(); i++) {
			List<Integer> row = grid.get(i);
			for (int k = 0; k < grid.get(0).size(); k++) {
				if (row.get(k) == v) {
					p.add(Pair.with(i, k));
				}
			}
		}
		return p;
	}

	static void print(String s) {
		System.out.println(s);
	}


}

