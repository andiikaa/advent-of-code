import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Map;

import static java.lang.System.out;

public class Main4 {

	public static void main(String[] args) throws IOException {
		Map<String, Integer> map = Map.of("A X", 3, "B X", 1, "C X", 2, "A Y", 4, "B Y", 5, "C Y", 6, "A Z", 8, "B Z", 9, "C Z", 7);
		out.println("max: " + Files.readAllLines(Paths.get("andre/03/input.txt")).stream().mapToInt(map::get).sum());
	}
}