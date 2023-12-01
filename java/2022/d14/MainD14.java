import org.javatuples.Pair;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.IntStream;

public class MainD14 {

	public static void main(String[] args) throws Exception {
		new MainD14().run();
	}

	static int maxX = 0;
	static int maxY = 0;
	static int minX = Integer.MAX_VALUE;
	static int minY = Integer.MAX_VALUE;

	List<List<String>> grid;

	int falls = 0;

	void run() throws IOException {
		List<List<Pair<Integer, Integer>>> pairs = Files.readAllLines(Paths.get("andre/d14/input.txt")).stream().map(MainD14::split).toList();

		int xNr = maxX + 1;
		int yNr = maxY + 1;

		grid = new ArrayList<>();
		IntStream.range(0, yNr + 5).forEach(i ->
				grid.add(new ArrayList<>(IntStream.range(0, xNr).mapToObj(k -> ".").toList()))
		);

		for (List<Pair<Integer, Integer>> lp : pairs) {
			int lastX = -1;
			int lastY = -1;
			for (Pair<Integer, Integer> p : lp) {
				int currX = p.getValue0();
				int curry = p.getValue1();

				if (lastX < 0) lastX = currX;
				if (lastY < 0) lastY = curry;

				if (lastX == currX) {
					// walk y
					if (lastY < curry || lastY == curry)
						IntStream.rangeClosed(lastY, curry).forEach(k -> grid.get(k).set(currX, "#"));
					else
						IntStream.rangeClosed(curry, lastY).forEach(k -> grid.get(k).set(currX, "#"));
				} else if (lastY == curry) {
					// walk x
					if (lastX < currX || lastX == currX)
						IntStream.rangeClosed(lastX, currX).forEach(k -> grid.get(curry).set(k, "#"));
					else
						IntStream.rangeClosed(currX, lastX).forEach(k -> grid.get(curry).set(k, "#"));
				}

				lastX = currX;
				lastY = curry;

			}
		}

		while (fall2()) {
			falls++;
		}

		printGrid(grid);

		// 210 :(
		// 592
		System.out.println("falls: " + falls);
	}

	boolean fall2() {
		int x = 500;
		int yOld = 0;
		int xOld = x;

		for (int y = 0; y < grid.size(); y++) {
			if (get(xOld, y).equals(".")) {
				yOld = y;
			} else if (get(xOld, y).equals("~")) {
				// endless
				return false;
			} else {
				// block or something
				if (get(xOld - 1, y).equals(".")) {
					xOld--;
					yOld = y;
				} else if (get(xOld + 1, y).equals(".")) {
					xOld++;
					yOld = y;
				}

				String left = get(xOld - 1, yOld + 1);
				String middle = get(xOld, yOld + 1);
				String right = get(xOld + 1, yOld + 1);

				if (middle.equals(".")) {
					yOld++;
				} else if (left.equals(".")) {
					xOld--;
					yOld++;
				} else if (right.equals(".")) {
					xOld++;
					yOld++;
				}

				if (!left.equals(".") && !left.equals("~") && !right.equals(".") && !right.equals("~") && !middle.equals(".") && !middle.equals("~")) {
					set(xOld, yOld, "o");
					return true;
				}
			}
		}
		return false;
	}

	void set(int x, int y, String v) {
		grid.get(y).set(x, v);
	}

	String get(int x, int y) {
		if (x < grid.get(0).size() && x >= 0 && y >= 0 && y < grid.size()) {
			return grid.get(y).get(x);
		}
		return "~";
	}

	void printGrid(List<List<String>> grid) {
		for (int i = 0; i < grid.size(); i++) {
			List<String> row = grid.get(i);
			for (int k = minX - 10; k < grid.get(0).size(); k++) {
				System.out.print(row.get(k));
			}
			System.out.println();
		}
	}

	static List<Pair<Integer, Integer>> split(String s) {
		return Arrays.stream(s.split(" -> ")).map(v -> {
			String[] split = v.split(",");
			int x = Integer.parseInt(split[0]);
			int y = Integer.parseInt(split[1]);
			maxX = Math.max(x, maxX);
			maxY = Math.max(y, maxY);
			minX = Math.min(x, minX);
			minY = Math.min(y, minY);
			return Pair.with(x, y);
		}).toList();
	}
}