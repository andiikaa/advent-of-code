import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;

public class Main {

	public static void main(String[] args) throws IOException {
		int max = Arrays.stream(Files.readString(Paths.get("andre/01/input.txt")).split("\n\n"))
				.map(s -> Arrays.asList(s.split("\n")))
				.mapToInt(l -> l.stream().mapToInt(Integer::parseInt).sum())
				.max()
				.getAsInt();
		System.out.println("max: " + max);
	}
}