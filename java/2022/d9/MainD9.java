import org.javatuples.Pair;

import java.io.IOException;
import java.nio.file.Files;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import static java.nio.file.Paths.get;

public class MainD9 {

	List<Move> m;
	Set<Pair<Integer, Integer>> fields = new HashSet<>();
	final int start = 1000;
	int row = start;
	int col = start;

	int rowT = start;
	int colT = start;

	public static void main(String[] args) throws IOException {
		new MainD9().run();
	}

	void run() throws IOException {
		m = Files.readAllLines((get("andre/d9/example.txt")))
				.stream().map(this::toMove).toList();
		fields.add(Pair.with(rowT, colT));
		for (Move v : m) {
			print("\n" + v);
			v.move();
		}

		// 6311
		print("count: " + fields.size());
	}

	Move toMove(String s) {
		int c = Integer.parseInt(s.split(" ")[1]);
		if (s.startsWith("U")) return new Up(c);
		if (s.startsWith("D")) return new Down(c);
		if (s.startsWith("L")) return new Left(c);
		if (s.startsWith("R")) return new Right(c);
		throw new IllegalArgumentException(s);
	}

	abstract class Move {
		final int steps;

		Move(int steps) {
			this.steps = steps;
		}

		abstract void move();

		void moveT() {
			if (row == rowT && col == colT) {
				print("same pos");
			} else if (row == rowT && colT > col + 1) {
				print("left");
				colT--;
			} else if (row == rowT && colT < col - 1) {
				print("right");
				colT++;
			} else if (col == colT && rowT > row + 1) {
				print("top");
				rowT--;
			} else if (col == colT && rowT < row - 1) {
				print("down");
				rowT++;
			} else if ((col > colT && row == rowT - 2) || (col == colT + 2 && row == rowT - 1)) {
				print("top right");
				rowT--;
				colT++;
			} else if ((col < colT && row == rowT - 2) || (col == colT - 2 && row == rowT - 1)) {
				print("top left");
				rowT--;
				colT--;
			} else if ((col > colT && row == rowT + 2) || (col == colT + 2 && row == rowT + 1)) {
				print("down right");
				rowT++;
				colT++;
			} else if ((col < colT && row == rowT + 2) || (col == colT - 2 && row == rowT + 1)) {
				print("down left");
				rowT++;
				colT--;
			} else {
				print("close by");
			}

			fields.add(Pair.with(rowT, colT));
		}

	}

	class Up extends Move {

		Up(int steps) {
			super(steps);
		}

		void move() {
			for (int i = 0; i < steps; i++) {
				row--;
				moveT();
			}
		}

		@Override
		public String toString() {
			return "up " + steps;
		}
	}

	class Down extends Move {
		Down(int steps) {
			super(steps);
		}

		void move() {
			for (int i = 0; i < steps; i++) {
				row++;
				moveT();
			}
		}

		@Override
		public String toString() {
			return "down " + steps;
		}
	}

	class Left extends Move {
		Left(int steps) {
			super(steps);
		}

		void move() {
			for (int i = 0; i < steps; i++) {
				col--;
				moveT();
			}
		}

		@Override
		public String toString() {
			return "left " + steps;
		}
	}

	class Right extends Move {
		Right(int steps) {
			super(steps);
		}

		void move() {
			for (int i = 0; i < steps; i++) {
				col++;
				moveT();
			}
		}

		@Override
		public String toString() {
			return "right " + steps;
		}
	}

	static void print(String s) {
		System.out.println(s);

	}


	static int parse(char c) {
		return Integer.parseInt(String.valueOf(c));
	}

}

