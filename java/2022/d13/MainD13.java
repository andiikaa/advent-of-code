import com.google.common.collect.Lists;

import java.io.IOException;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import static java.nio.file.Paths.get;

public class MainD13 {

	public static void main(String[] args) throws Exception {
		new MainD13().run();
	}

	void run() throws IOException {
		List<String> strings = Lists.newArrayList(Files.readString((get("andre/d13/input.txt"))).split("\n\n"));
		List<Integer> alreadyOrdered = new ArrayList<>();

		for (int number = 0; number < strings.size(); number++) {
			String[] split = strings.get(number).split("\n");
			List<Object> left = new ArrayList<>();
			List<Object> right = new ArrayList<>();
			parseL2(split[0].substring(1, split[0].length() - 1), left, 0);
			parseL2(split[1].substring(1, split[1].length() - 1), right, 0);
			int order = getOrder(left, right);
			if (order < 0) {
				alreadyOrdered.add(number + 1);
			}
		}
		// 5913 :(
		// 5580
		System.out.println("ordered " + alreadyOrdered.stream().mapToInt(i -> i).sum());
	}

	// -1 left is smaller, 0 means equal, 1 means greater (bad)
	int getOrder(Object left, Object right) {
		if (left instanceof Integer leftI && right instanceof Integer rightI) {
			return leftI.compareTo(rightI);
		}

		if (left instanceof Integer) {
			return getOrder(Collections.singletonList(left), right);
		}

		if (right instanceof Integer) {
			return getOrder(left, Collections.singletonList(right));
		}

		if (left instanceof List leftL && right instanceof List rightL) {
			if (leftL.isEmpty() && !rightL.isEmpty()) {
				return -1;
			}

			for (int i = 0; i < leftL.size(); i++) {
				if (i >= rightL.size()) {
					return 1;
				}
				Object oL = leftL.get(i);
				Object oR = rightL.get(i);
				int order = getOrder(oL, oR);
				if (order != 0) return order;
			}
			return leftL.size() == rightL.size() ? 0 : -1;
		}

		return 1;
	}

	int parseL2(String input, List<Object> parsed, int currentIndex) {
		while (currentIndex < input.length()) {
			char c = input.charAt(currentIndex);
			Integer number = null;
			if (c != ',' && c != '[' && c != ']') {
				// integer parsing only works for 2 digits ;P
				int next = currentIndex + 2;
				if (next <= input.length()) {
					number = parseInt(input.substring(currentIndex, next));
				}
				if (number == null) {
					number = parseInt(String.valueOf(c));
					if (number == null) throw new IllegalStateException("failed to parse to int " + c);
				} else {
					currentIndex++;
				}
			}

			if (number != null) parsed.add(number);
			if (c == '[') {
				List<Object> list = new ArrayList<>();
				currentIndex = parseL2(input, list, currentIndex + 1);
				parsed.add(list);
			}
			if (c == ']') return currentIndex;
			currentIndex++;
		}
		return currentIndex;
	}


	static Integer parseInt(String s) {
		try {
			return Integer.parseInt(s);
		} catch (NumberFormatException e) {
			return null;
		}
	}
}