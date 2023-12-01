import java.io.IOException;
import java.nio.file.Files;
import java.util.List;

import static java.lang.System.out;
import static java.nio.file.Paths.get;

public class MainD8 {

	static List<String> s;

	public static void main(String[] args) throws IOException {
		s = Files.readAllLines((get("andre/d8/input.txt")));

		int count = 0;
		for (int row = 1; row < s.size() - 1; row++) {
			String r = s.get(row);
			for (int col = 1; col < r.length() - 1; col++) {
				if (visible(row, col)) count++;
			}
		}

		count += s.size() * 2 + (s.get(0).length() - 2) * 2;
		out.println("sum: " + count);
	}

	static boolean visible(int row, int col) {
		String s1 = s.get(row);
		int height = parse(s1.charAt(col));
		int tmp = 0;

		for (int i = row - 1; i > -1; i--)
			if (parse(s.get(i).charAt(col)) >= height) {
				tmp++;
				break;
			}

		for (int i = row + 1; i < s.size(); i++)
			if (parse(s.get(i).charAt(col)) >= height) {
				tmp++;
				break;
			}

		for (int i = col - 1; i > -1; i--)
			if (parse(s1.charAt(i)) >= height) {
				tmp++;
				break;
			}

		for (int i = col + 1; i < s1.length(); i++)
			if (parse(s1.charAt(i)) >= height) {
				tmp++;
				break;
			}

		return tmp < 4;
	}

	static int parse(char c) {
		return Integer.parseInt(String.valueOf(c));
	}

}

