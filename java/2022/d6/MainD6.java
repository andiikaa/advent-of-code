import java.io.IOException;

import static java.lang.System.out;
import static java.nio.file.Files.readString;
import static java.nio.file.Paths.get;

public class MainD6 {

	public static void main(String[] args) throws IOException {
		String s = readString(get("andre/d6/input.txt"));
		int count = 0;
		for (int i = 0; i < s.length(); i++)
			if (s.substring(i, i + 14).chars().distinct().count() == 14) {
				count = i + 14;
				break;
			}
		out.println("result: " + count);
	}


}

