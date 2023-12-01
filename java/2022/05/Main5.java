import java.io.IOException;
import java.nio.file.Paths;
import java.util.List;

import static java.lang.System.out;
import static java.nio.file.Files.readAllLines;

public class Main5 {

	public static void main(String[] args) throws IOException {
		List<Integer> i = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".chars().boxed().toList();
		int sum = readAllLines(Paths.get("andre/05/input.txt"))
				.stream()
				.map(s -> new String[]{s.substring(0, s.length() / 2), s.substring(s.length() / 2)})
				.mapToInt(s -> s[0].chars().boxed().distinct().filter(n -> s[1].indexOf(n) > -1).mapToInt(i::indexOf).sum())
				.sum();
		out.println("sum:" + sum);
	}
}