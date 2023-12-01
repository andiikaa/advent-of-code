import org.javatuples.Pair;

import java.math.BigInteger;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.BiConsumer;
import java.util.function.Consumer;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.IntStream;
import java.util.stream.Stream;

// TODO not working
public class MainD15_2 {

	public static void main(String[] args) throws Exception {
		List<Device> devices = Files.readAllLines(Paths.get("andre/d15/input.txt")).stream().flatMap(MainD15_2::split).toList();
		AtomicBoolean found = new AtomicBoolean(false);
		IntStream.rangeClosed(0, 4000000).parallel().forEach(row -> {
			if (found.get()) return;
			BigInteger tuning = new MainD15_2(row, devices).findTuning();
			found.set(!tuning.equals(BigInteger.ZERO));
		});
		System.out.println("done");
	}

	static volatile int maxX = 0;
	static volatile int maxY = 0;
	static volatile int minX = Integer.MAX_VALUE;
	static volatile int minY = Integer.MAX_VALUE;

	List<Device> devices;
	int row;
	List<String> selectedRow;

	MainD15_2(int row, List<Device> devices) {
		this.row = row;
		this.devices = devices;
		selectedRow = new ArrayList<>(IntStream.rangeClosed(minX, maxX).mapToObj(k -> ".").toList());
	}

	BigInteger findTuning() {
		System.out.println("checking row " + row);
		BiConsumer<Pair<Integer, Integer>, String> valueSet = (p, v) -> set(p.getValue0(), p.getValue0(), v);
		Consumer<Pair<Integer, Integer>> radarSet = p -> setRadar(p.getValue0(), p.getValue0());

		selectedRow = new ArrayList<>(IntStream.rangeClosed(minX, maxX).mapToObj(k -> ".").toList());
		devices.forEach(d -> d.add(row, valueSet, radarSet));
		long sum = selectedRow.stream().filter(s -> s.equals(".")).count();
		if (sum != 1) {
			return BigInteger.ZERO;
		}
		Device d = devices.get(0);
		BigInteger tuning = BigInteger.valueOf(4000000).multiply(BigInteger.valueOf(d.x)).add(BigInteger.valueOf(d.y));

//		5073496
		System.out.println("sum: " + sum);
		System.out.println("add: " + tuning);
		return tuning;
	}

	void setRadar(int x, int y) {
		y = Math.abs(minY) + y;
		x = Math.abs(minX) + x;
		int finalRow = Math.abs(minY) + row;
		if (y == finalRow && x >= 0 && x < selectedRow.size()) {
			if (!selectedRow.get(x).equals("B") && !selectedRow.get(x).equals("S")) {
				selectedRow.set(x, "#");
			}
		}
	}

	void set(int x, int y, String v) {
		int finalRow = Math.abs(minY) + row;
		x = Math.abs(minX) + x;
		y = Math.abs(minY) + y;
		if (y == finalRow && x > 0 && x < selectedRow.size()) {
			selectedRow.set(x, v);
		}
	}

	static Stream<Device> split(String s) {
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

	abstract static class Device {
		int x;
		int y;

		Device(int x, int y) {
			this.x = x;
			this.y = y;
		}

		void add(int row, BiConsumer<Pair<Integer, Integer>, String> valueSet, Consumer<Pair<Integer, Integer>> radarSet) {
			valueSet.accept(Pair.with(x, y), toString());
		}

	}


	static class Beacon extends Device {
		Beacon(int x, int y) {
			super(x, y);
		}

		@Override
		public String toString() {
			return "B";
		}
	}

	static class Sensor extends Device {
		Beacon closest;

		Sensor(int x, int y) {
			super(x, y);
		}

		@Override
		public String toString() {
			return "S";
		}

		@Override
		void add(int row, BiConsumer<Pair<Integer, Integer>, String> valueSet, Consumer<Pair<Integer, Integer>> radarSet) {
			super.add(row, valueSet, radarSet);

			int distance = distance();
			int finalRow = Math.abs(minY) + row;

			// TODO only one loop needed here
			for (int i = 0; i < distance + 1; i++) {
				int row1 = y - distance + i;
				int row2 = y + distance - i;

				int tempRow = Math.abs(minY) + row1;
				int tempRow2 = Math.abs(minY) + row2;
				if (tempRow2 != finalRow && tempRow != finalRow) {
					continue;
				}

				for (int col = x - i; col <= x + i; col++) {
					radarSet.accept(Pair.with(col, row1));
					if (i != distance) radarSet.accept(Pair.with(col, row2));
				}
			}
		}

		int distance() {
			return Math.abs(x - closest.x) + Math.abs(y - closest.y);
		}

	}
}