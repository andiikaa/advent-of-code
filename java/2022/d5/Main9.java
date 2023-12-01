import java.io.IOException;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

import static com.google.common.base.Strings.padEnd;
import static com.google.common.collect.Lists.reverse;
import static java.lang.System.out;
import static java.nio.file.Files.readString;
import static java.nio.file.Paths.get;
import static java.util.Arrays.stream;
import static java.util.stream.IntStream.range;

public class Main9 {

	public static void main(String[] args) throws IOException {
		String data = readString(get("andre/d5/input.txt"));
		String[] split = data.split("\n\n");
		String[] rows = split[0].split("\n");
		int rowLength = rows[rows.length - 1].length() + 2;

		HashMap<Integer, Deque<String>> stacks = stream(rows)
				.limit(rows.length - 1)
				.map(r -> padEnd(r, rowLength, ' '))
				.map(r -> range(0, rowLength).filter(i -> i % 4 == 0).mapToObj(i -> r.substring(i, i + 3)).toList())
				.collect(HashMap::new, (s1, s2) -> {
					for (int i = 0; i < s2.size(); i++)
						if (!s2.get(i).trim().isEmpty()) s1.computeIfAbsent(i, ArrayDeque::new).addLast(s2.get(i));
				}, (s1, s2) -> {
				});

		List<List<String>> moves = stream(split[1].split("\n"))
				.map(s -> {
							List<String> finds = new ArrayList<>();
							Matcher matcher = Pattern.compile("\\d{1,6}").matcher(s);
							while (matcher.find()) finds.add(matcher.group());
							return finds;
						}
				).toList();

		for (List<String> move : moves) {
			int count = Integer.parseInt(move.get(0));
			int from = Integer.parseInt(move.get(1));
			int to = Integer.parseInt(move.get(2));

			Deque<String> fromS = stacks.get(from - 1);
			Deque<String> toS = stacks.get(to - 1);

			List<String> l = new ArrayList<>();
			for (int i = 0; i < count; i++) l.add(fromS.pop());
			reverse(l).forEach(toS::push);
		}

		String result = stacks.entrySet().stream().sorted(Map.Entry.comparingByKey())
				.map(Map.Entry::getValue)
				.map(Deque::pop)
				.collect(Collectors.joining())
				.replace("[", "")
				.replace("]", "");

		out.println("result  : " + result);
	}
}
