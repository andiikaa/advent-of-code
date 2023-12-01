import java.io.IOException;
import java.util.List;

import static com.google.common.collect.Lists.partition;
import static java.lang.System.out;
import static java.nio.file.Files.readAllLines;
import static java.nio.file.Paths.get;

public class Main6 {

	public static void main(String[] args) throws IOException {
		List<Integer> i = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".chars().boxed().toList();
		int sum = partition(readAllLines(get("andre/05/input.txt")), 3).stream()
				.map(l -> l.get(0).chars().distinct().filter(c -> l.get(1).indexOf(c) > -1 && l.get(2).indexOf(c) > -1)
						.findFirst()
						.getAsInt())
				.mapToInt(i::indexOf).sum();
		out.println("sum:" + sum);
	}
}
