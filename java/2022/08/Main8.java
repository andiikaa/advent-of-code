import java.io.IOException;

import static java.lang.System.out;
import static java.nio.file.Files.readAllLines;
import static java.nio.file.Paths.get;

public class Main8 {

	public static void main(String[] args) throws IOException {
		int sum = readAllLines(get("andre/07/input.txt"))
				.stream().map(s -> {
					String[] l = s.split(",");
					String[] s1 = l[0].split("-");
					String[] s2 = l[1].split("-");
					return new String[][]{s1, s2};
				}).mapToInt(s -> {
					int i1 = Integer.parseInt(s[0][0]);
					int i2 = Integer.parseInt(s[0][1]);
					int i3 = Integer.parseInt(s[1][0]);
					int i4 = Integer.parseInt(s[1][1]);
					return (i1 <= i3 && i2 >= i3) || (i3 <= i1 && i4 >= i1) ? 1 : 0;
				}).sum();
		out.println("sum:" + sum);
	}
}
