import java.io.IOException;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.List;

import static java.nio.file.Paths.get;

public class MainD10 {

	List<Op> m;
	List<Integer> i = new ArrayList<>();
	int x = 1;
	int cycle = 0;
	int cycleSum = 0;

	public static void main(String[] args) throws IOException {
		new MainD10().run();
	}

	void run() throws IOException {
		m = Files.readAllLines((get("andre/d10/input.txt")))
				.stream().map(this::toOp).toList();
		for (Op v : m) {
			print("" + v);
			v.exec();
		}

		System.out.println(i);
		print("count: " + i.stream().mapToInt(v -> v).sum());
	}

	Op toOp(String s) {
		if (s.startsWith("noop")) return new Noop();
		if (s.startsWith("addx")) return new Add(Integer.parseInt(s.split(" ")[1]));
		throw new IllegalArgumentException(s);
	}

	abstract class Op {

		abstract void exec();

	}

	class Noop extends Op {

		void exec() {
			cycle += 1;
			sum();
		}

		@Override
		public String toString() {
			return "noop";
		}
	}

	class Add extends Op {
		final int c;

		Add(int c) {
			this.c = c;
		}

		void exec() {
			cycle += 1;
			sum();
			cycle += 1;
			sum();
			x += c;
		}

		@Override
		public String toString() {
			return "addx " + c;
		}
	}

	void sum() {
		if (cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220) {
			i.add(cycle * x);
			print("#######################     cycle " + cycle + " sum " + cycleSum);
		}
	}

	static void print(String s) {
		System.out.println(s);

	}


	static int parse(char c) {
		return Integer.parseInt(String.valueOf(c));
	}

}

