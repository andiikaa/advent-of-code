import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.IntStream;
import java.util.stream.Stream;

public class MainD15 {

	public static void main(String[] args) throws Exception {
		new MainD15().run();
	}

	int maxX = 0;
	int maxY = 0;
	int minX = Integer.MAX_VALUE;
	int minY = Integer.MAX_VALUE;

	//	static final int ROW = 10;
	static final int ROW = 2000000;

	List<String> selectedRow;

	void run() throws IOException {
		List<Device> devices = Files.readAllLines(Paths.get("andre/d15/input.txt")).stream().flatMap(this::split).toList();
		selectedRow = new ArrayList<>(IntStream.rangeClosed(minX, maxX).mapToObj(k -> ".").toList());
		devices.forEach(Device::add);
		long sum = selectedRow.stream().filter(s -> s.equals("#")).count();
//		5073496
		System.out.println("sum: " + sum);
	}

	void setRadar(int x, int y) {
		y = Math.abs(minY) + y;
		x = Math.abs(minX) + x;
		int finalRow = Math.abs(minY) + ROW;
		if (y == finalRow && x >= 0 && x < selectedRow.size()) {
			if (!selectedRow.get(x).equals("B") && !selectedRow.get(x).equals("S")) {
				selectedRow.set(x, "#");
			}
		}
	}

	void set(int x, int y, String v) {
		int finalRow = Math.abs(minY) + ROW;
		x = Math.abs(minX) + x;
		y = Math.abs(minY) + y;
		if (y == finalRow && x > 0 && x < selectedRow.size()) {
			selectedRow.set(x, v);
		}
	}

	Stream<Device> split(String s) {
		List<Integer> finds = new ArrayList<>();
		Matcher matcher = Pattern.compile("-?\\d{1,20}").matcher(s);
		while (matcher.find()) finds.add(Integer.parseInt(matcher.group()));

		Sensor sensor = new Sensor(finds.get(0), finds.get(1));
		Beacon beacon = new Beacon(finds.get(2), finds.get(3));
		sensor.closest = beacon;
		maxX = Math.max(sensor.x + sensor.distance(), maxX);
		maxY = Math.max(sensor.y + sensor.distance(), maxY);
		minX = Math.min(sensor.x - sensor.distance(), minX);
		minY = Math.min(sensor.y - sensor.distance(), minY);
		maxX = Math.max(beacon.x, maxX);
		maxY = Math.max(beacon.y, maxY);
		minX = Math.min(beacon.x, minX);
		minY = Math.min(beacon.y, minY);
		return Stream.of(sensor, beacon);
	}

	abstract class Device {
		int x;
		int y;

		Device(int x, int y) {
			this.x = x;
			this.y = y;
		}

		void add() {
			set(x, y, toString());
		}

	}

	class Beacon extends Device {
		Beacon(int x, int y) {
			super(x, y);
		}

		@Override
		public String toString() {
			return "B";
		}
	}

	class Sensor extends Device {
		Beacon closest;

		Sensor(int x, int y) {
			super(x, y);
		}

		@Override
		public String toString() {
			return "S";
		}

		@Override
		void add() {
			super.add();
			addRadar();
		}

		int distance() {
			return Math.abs(x - closest.x) + Math.abs(y - closest.y);
		}

		void addRadar() {
			int distance = distance();
			int finalRow = Math.abs(minY) + ROW;

			// TODO only one loop needed here
			for (int i = 0; i < distance + 1; i++) {
				int row = y - distance + i;
				int row2 = y + distance - i;

				int tempRow = Math.abs(minY) + row;
				int tempRow2 = Math.abs(minY) + row2;
				if (tempRow2 != finalRow && tempRow != finalRow) {
					continue;
				}

				for (int col = x - i; col <= x + i; col++) {
					setRadar(col, row);
					if (i != distance) setRadar(col, row2);
				}
			}
		}
	}
}