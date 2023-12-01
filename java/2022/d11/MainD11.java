import org.apache.commons.math3.primes.Primes;

import java.io.IOException;
import java.math.BigInteger;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.Function;

import static java.nio.file.Paths.get;

public class MainD11 {

	Map<Integer, Monkey> collect = new HashMap<>();
	List<Integer> primes;

	public static void main(String[] args) throws IOException {
		new MainD11().run();
	}

	void run() throws IOException {
		Arrays.stream(Files.readString((get("andre/d11/example.txt"))).split("\n\n")).map(this::toM).forEach(monkey -> collect.put(monkey.id, monkey));
		List<Monkey> monkeys = collect.values().stream().sorted(Comparator.comparing(Monkey::getId)).toList();
		primes = monkeys.stream().map(Monkey::getDevisor).toList();

		for (int i = 0; i < 10000; i++) {
			monkeys.forEach(Monkey::testAll);
			System.out.print("" + i + "\r");
			//monkeys.forEach(mo -> print("" + mo.getId() + " " + mo.starting));
		}

		List<BigInteger> active = monkeys.stream().map(Monkey::getInspected).sorted(Comparator.reverseOrder()).limit(2).toList();
		print("active: " + active);
		print("active: " + (active.get(0).multiply(active.get(1))));
	}

	Monkey toM(String s) {
		Monkey m = new Monkey();
		String[] split = s.split("\n");
		m.id = Integer.parseInt(split[0].substring(split[0].indexOf(" "), split[0].lastIndexOf(":")).trim());
		m.starting = new ArrayList<>(Arrays.stream(split[1].split(":")[1].split(",")).map(v -> BigInteger.valueOf(Integer.parseInt(v.trim()))).toList());
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
		Function<BigInteger, BigInteger> op;
		List<BigInteger> starting;

		BigInteger inspected = BigInteger.ZERO;

		int getDevisor() {
			return devisor;
		}

		BigInteger getInspected() {
			return inspected;
		}

		void addW(BigInteger w) {
			//print("" + id + " received " + w);
			starting.add(w);
		}

		int getId() {
			return id;
		}

		void refit() {

		}

		void testAll() {
			inspected = inspected.add(BigInteger.valueOf(starting.size()));
			starting.sort(Comparator.naturalOrder());
			starting.forEach(this::test);
			starting.clear();
		}

		boolean test1(BigInteger i) {
			BigInteger r = op.apply(i);
			return r.mod(BigInteger.valueOf(devisor)).equals(BigInteger.ZERO);
		}

		void test(BigInteger i) {
			//print("\n" + id + " inspects " + i);
			BigInteger r = op.apply(i);
			//r = r.divide(BigInteger.valueOf(3));
			//r = toPrimeMult(r);

			if (r.mod(BigInteger.valueOf(devisor)).equals(BigInteger.ZERO)) {
				Monkey m = collect.get(whenTrue);
				//m.test1(r);
				m.addW(r);
			} else {
				Monkey m = collect.get(whenFalse);
				m.addW(r);
			}
		}

		BigInteger toPrimeMult(BigInteger i) {
			int n = i.intValue();
			List<Integer> p = Primes.primeFactors(n);
			AtomicInteger c = new AtomicInteger(1);
			p.stream().filter(primes::contains).forEach(v -> c.set(c.get() * v));
			return BigInteger.valueOf(c.get());
		}
	}

	static Function<BigInteger, BigInteger> toF(String s) {
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
			BigInteger left = nL == null ? old : BigInteger.valueOf(nL);
			BigInteger right = nR == null ? old : BigInteger.valueOf(nR);
			if (isPlus) return left.add(right);
			return left.multiply(right);
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

