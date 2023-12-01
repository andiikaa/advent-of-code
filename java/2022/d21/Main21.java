import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.Map;
import java.util.function.BinaryOperator;

public class Main21 {

	public static void main(String[] args) throws IOException {
		new Main21().run();
	}

	void run() throws IOException {
		Map<String, Monkey> input = Files.readAllLines(Paths.get("andre/d21/input.txt")).stream()
				.map(this::toMonkey)
				.collect(HashMap::new, (c, m) -> c.putIfAbsent(m.id, m), (c1, c2) -> {
				});
		Monkey root = input.get("root");
		long rootVal = root.executeAll(input);
		root.operation = (l, r) -> (long) l.compareTo(r);
		Monkey human = input.get("humn");
		human.value = 0L;

		long humanValue = input.get("root").predictValue(input, 0L);

		// 232974643455000
		System.out.println("root yelling:        " + rootVal);
		// 3740214169961
		System.out.println("human needs to yell: " + humanValue);
	}

	Monkey toMonkey(String s) {
		String[] split = s.split(":");
		Monkey m = new Monkey(split[0]);

		String[] op = null;
		BinaryOperator<Long> operation = null;
		BinaryOperator<Long> reverseLeft = null;
		BinaryOperator<Long> reverseRight = null;
		if (split[1].contains("+")) {
			op = split[1].split("\\+");
			operation = Long::sum;
			reverseLeft = (y, r) -> y - r;
			reverseRight = (y, l) -> y - l;
		} else if (split[1].contains("-")) {
			op = split[1].split("-");
			operation = (l, r) -> l - r;
			reverseLeft = Long::sum;
			reverseRight = (y, l) -> l - y;
		} else if (split[1].contains("*")) {
			op = split[1].split("\\*");
			operation = (l, r) -> l * r;
			reverseLeft = (y, r) -> y / r;
			reverseRight = (y, l) -> y / l;
		} else if (split[1].contains("/")) {
			op = split[1].split("/");
			operation = (l, r) -> l / r;
			reverseLeft = (y, r) -> y * r;
			reverseRight = (y, l) -> l / y;
		}

		if (op == null) {
			m.value = Long.parseLong(split[1].trim());
		} else {
			m.operation = operation;
			m.leftMonkey = op[0].trim();
			m.rightMonkey = op[1].trim();
			m.reverseRight = reverseRight;
			m.reverseLeft = reverseLeft;
		}

		return m;
	}

	class Monkey {
		private final String id;

		Long value;
		Long left;
		Long right;
		BinaryOperator<Long> operation;
		BinaryOperator<Long> reverseLeft;
		BinaryOperator<Long> reverseRight;
		String leftMonkey;
		String rightMonkey;

		Monkey(String id) {
			this.id = id;
		}

		long predictValue(Map<String, Monkey> input, long valueNeeded) {
			if (id.equals("humn")) return valueNeeded;
			if (value != null) return value;

			if (id.equals("root")) {
				valueNeeded = input.get(rightMonkey).executeAll(input);
				return input.get(leftMonkey).predictValue(input, valueNeeded);
			}

			Monkey human = input.get("humn");
			long leftTmp = input.get(leftMonkey).executeAll(input);
			long rightTmp = input.get(rightMonkey).executeAll(input);

			human.value = 33L;

			left = input.get(leftMonkey).executeAll(input);
			right = input.get(rightMonkey).executeAll(input);

			human.value = 0L;

			if (left != leftTmp) {
				valueNeeded = reverseLeft.apply(valueNeeded, right);
				valueNeeded = input.get(leftMonkey).predictValue(input, valueNeeded);
			} else if (right != rightTmp) {
				valueNeeded = reverseRight.apply(valueNeeded, left);
				valueNeeded = input.get(rightMonkey).predictValue(input, valueNeeded);
			} else {
				throw new IllegalStateException("could not detect the influenced branch");
			}

			return valueNeeded;
		}

		long executeAll(Map<String, Monkey> input) {
			if (value != null) return value;

			left = input.get(leftMonkey).executeAll(input);
			right = input.get(rightMonkey).executeAll(input);
			return operation.apply(left, right);
		}
	}
}
