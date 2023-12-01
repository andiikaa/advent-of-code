import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.Collections;

public class Main2 {

	public static void main(String[] args) throws IOException {
		int max = Arrays.stream(Files.readString(Paths.get("andre/01/input.txt")).split("\n\n"))
				.map(s -> Arrays.asList(s.split("\n")))
				.map(l -> l.stream().mapToInt(Integer::parseInt).sum())
				.sorted(Collections.reverseOrder())
				.mapToInt(Integer::intValue)
				.limit(3)
				.sum();
		System.out.println("max: " + max);
	}
}