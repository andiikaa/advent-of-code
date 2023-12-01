import java.io.IOException;
import java.nio.file.Files;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

import static java.lang.System.out;
import static java.nio.file.Paths.get;

public class MainD7_2 {

	static Map<String, Integer> size = new TreeMap<>();

	public static void main(String[] args) throws IOException {
		List<String> s = Files.readAllLines((get("andre/d7/input.txt")));
		String current = "";
		for (String c : s) {
			if (c.startsWith("$")) {
				String cmd = c.substring(2);
				if (cmd.startsWith("cd")) {
					cmd = cmd.substring(3);
					if (cmd.equals("/")) current = cmd;
					else if (cmd.equals("..") && current.lastIndexOf("/") > 0)
						current = current.substring(0, current.lastIndexOf("/"));
					else if (cmd.equals("..")) current = "/";
					else current += current.length() == 1 ? cmd : "/" + cmd;
				}
			} else if (!c.startsWith("dir")) {
				size.merge(current, Integer.parseInt(c.split(" ")[0]), Integer::sum);
			} else {
				String[] s1 = c.split(" ");
				String currentTmp = current.endsWith("/") ? current + s1[1] : current + "/" + s1[1];
				size.putIfAbsent(currentTmp, 0);
			}
		}

		int totalSpace = size.values().stream().mapToInt(e -> e).sum();
		int needToFree = Math.abs(40000000 - totalSpace);

		String deleteF = "";
		int cSize = Integer.MAX_VALUE;

		for (String f : size.keySet()) {
			if (f.equals("/")) continue;
			int tmp = size(f);
			if (tmp >= needToFree && tmp < cSize) {
				deleteF = f;
				cSize = tmp;
			}
		}

		// 1300850
		out.println("deleting: " + deleteF + " frees: " + cSize);
	}

	static int size(String key) {
		int tmp = 0;
		for (Map.Entry<String, Integer> e : size.entrySet())
			if (e.getKey().startsWith(key))
				tmp += e.getValue();
		return tmp;
	}


}

