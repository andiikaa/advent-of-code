import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

import static java.lang.System.out;

public class Main3 {

	public static void main(String[] args) throws IOException {
		int[] v = {4, 1, 7, 0, 1, 2};
		int sum = Files.readAllLines(Paths.get("andre/03/input.txt"))
				.stream()
				.map(s -> s.split(" "))
				.mapToInt(s -> {
					int he = s[0].getBytes()[0] - 65;
					int me = s[1].getBytes()[0] - 88;
					return v[((he + me + v[me + 3]) % 3)] + v[me + 3];
				}).sum();
		out.println("max: " + sum);
	}

	// a x stein 1
	// b y papier 2
	// c z schere 3
	// win: 6 draw: 3: loose: 0

	// offset 65 -> a 0
	// offset 88 -> x 0

	// 00 -> draw  3 1 4 0
	// 10 -> loose 0 1 1 1
	// 20 -> win   6 1 7 2

	// 01 -> win   6 2 8 1
	// 11 -> draw  3 2 5 2
	// 21 -> loose 0 2 2 3

	// 02 -> loose 0 3 3 2
	// 12 -> win   6 3 9 3
	// 22 -> draw  3 3 6 4

	// 0 + 0 = 0 -> 4
	// 1 + 0 = 1 -> 1
	// 2 + 0 = 2 -> 7

	// 1 + 1 = 2 -> 8
	// 2 + 1 = 3 -> 5
	// 3 + 1 = 4 -> 2

	// 2 + 2 = 4 -> 3
	// 3 + 2 = 5 -> 9
	// 4 + 2 = 6 -> 6
}