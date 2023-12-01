import java.io.IOException;
import java.util.List;

import static java.lang.Math.max;
import static java.lang.System.out;
import static java.nio.file.Files.readAllLines;
import static java.nio.file.Paths.get;

public class MainD8_2 {

	static List<String> s;

	public static void main(String[] args) throws IOException {
		s = readAllLines((get("andre/d8/input.txt")));

		int count = 0;
		for (int row = 1; row < s.size() - 1; row++) {
			String r = s.get(row);
			for (int col = 1; col < r.length() - 1; col++) {
				count = max(sonic(row, col), count);
			}
		}

		out.println("max: " + count);
	}

	static int sonic(int row, int col) {
		String s1 = s.get(row);
		int height = parse(s1.charAt(col));
		int top = 0;
		int down = 0;
		int left = 0;
		int right = 0;

		for (int i = row - 1; i > -1; i--)
			if (parse(s.get(i).charAt(col)) >= height) {
				top++;
				break;
			} else top++;

		for (int i = row + 1; i < s.size(); i++)
			if (parse(s.get(i).charAt(col)) >= height) {
				down++;
				break;
			} else down++;

		for (int i = col - 1; i > -1; i--)
			if (parse(s1.charAt(i)) >= height) {
				left++;
				break;
			} else left++;

		for (int i = col + 1; i < s1.length(); i++)
			if (parse(s1.charAt(i)) >= height) {
				right++;
				break;
			} else right++;

		return top * down * left * right;
	}

	static int parse(char c) {
		return Integer.parseInt(String.valueOf(c));
	}

}

