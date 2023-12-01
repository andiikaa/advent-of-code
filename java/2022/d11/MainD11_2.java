import java.io.IOException;
import java.math.BigInteger;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.function.Function;

import static java.nio.file.Paths.get;

public class MainD11_2 {

	Map<Integer, Monkey> collect = new HashMap<>();
	List<Integer> primes;
	static int mult = 1;

	public static void main(String[] args) throws IOException {
		new MainD11_2().run();
	}

	void run() throws IOException {
		Arrays.stream(Files.readString((get("andre/d11/input.txt"))).split("\n\n")).map(this::toM).forEach(monkey -> collect.put(monkey.id, monkey));
		List<Monkey> monkeys = collect.values().stream().sorted(Comparator.comparing(Monkey::getId)).toList();
		primes = monkeys.stream().map(Monkey::getDevisor).toList();
		primes.stream().forEach(p -> mult = mult * p);
		print("mult " + mult);

		for (int i = 0; i < 10000; i++) {
			monkeys.forEach(Monkey::testAll);
		}

		List<BigInteger> all = monkeys.stream().map(Monkey::getInspected).toList();
		List<BigInteger> active = all.stream().sorted(Comparator.reverseOrder()).limit(2).toList();
		print("all: " + all);
		print("active: " + active);
		print("active: " + (active.get(0).multiply(active.get(1))));
	}

	Monkey toM(String s) {
		Monkey m = new Monkey();
		String[] split = s.split("\n");
		m.id = Integer.parseInt(split[0].substring(split[0].indexOf(" "), split[0].lastIndexOf(":")).trim());
		m.starting = new ArrayList<>(Arrays.stream(split[1].split(":")[1].split(",")).map(v -> Integer.parseInt(v.trim())).toList());
		m.devisor = Integer.parseInt(split[3].split("by")[1].trim());
		m.whenTrue = Integer.parseInt(split[4].substring(split[4].lastIndexOf(" ")).trim());
		m.whenFalse = Integer.parseInt(split[5].substring(split[5].lastIndexOf(" ")).trim());
		m.op = toF(split[2]);
		return m;
	}

	class Monkey {
		int id;
		int devisor;
		int whenTrue;
		int whenFalse;
		Function<Integer, Integer> op;
		List<Integer> starting;

		BigInteger inspected = BigInteger.ZERO;

		int getDevisor() {
			return devisor;
		}

		BigInteger getInspected() {
			return inspected;
		}

		void addW(int w) {
			//print("" + id + " received " + w);
			starting.add(w);
		}

		int getId() {
			return id;
		}

		void testAll() {
			inspected = inspected.add(BigInteger.valueOf(starting.size()));
			starting.sort(Comparator.naturalOrder());
			starting.forEach(this::test);
			starting.clear();
		}

		void test(int i) {
			//print("\n" + id + " inspects " + i);

			int r = op.apply(i % mult) % mult;

			if (r % devisor == 0) {
				Monkey m = collect.get(whenTrue);
				m.addW(r);
			} else {
				Monkey m = collect.get(whenFalse);
				m.addW(r);
			}
		}
	}

	static Function<Integer, Integer> toF(String s) {
		boolean isPlus;
		String[] split;
		s = s.substring(s.lastIndexOf("=") + 1);

		if (s.contains("*")) {
			isPlus = false;
			split = s.split("\\*");
		} else if (s.contains("+")) {
			isPlus = true;
			split = s.split("\\+");
		} else {
			throw new IllegalArgumentException(s);
		}

		Integer nL = parse(split[0]);
		Integer nR = parse(split[1]);
		return old -> {
			Integer left = nL == null ? old : nL;
			Integer right = nR == null ? old : nR;
			if (isPlus)
				return BigInteger.valueOf(left).add(BigInteger.valueOf(right)).mod(BigInteger.valueOf(mult)).intValue();
			return BigInteger.valueOf(left).multiply(BigInteger.valueOf(right)).mod(BigInteger.valueOf(mult)).intValue();
		};
	}

	static Integer parse(String s) {
		try {
			return Integer.parseInt(s.trim());
		} catch (NumberFormatException e) {
			return null;
		}
	}

	static void print(String s) {
		System.out.println(s);
	}


}

