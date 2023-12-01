import org.apache.commons.lang3.Range;
import org.javatuples.Pair;

import java.io.IOException;
import java.math.BigInteger;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.LongStream;
import java.util.stream.Stream;

// TODO not working
public class MainD15_3 {

	public static void main(String[] args) throws Exception {
//		new MainD15_3().run();
		new MainD15_3().run2();
	}

	//		static final int ROW = 10;
	static final long ROW = 2000000;
	static final long MAX = 4000000;
//	static final long MAX = 20;

	static final int ROWS = 1000;
	static final int COLS = (int) MAX;

	static final Range<Long> X_RANGE = Range.between(0L, MAX - 1);


	byte[][] grid;
	long shift = 0;

	void run2() throws IOException {
		// might crash if not enough memory -Xmx12G
		grid = new byte[ROWS][COLS];

		List<Device> devices = Files.readAllLines(Paths.get("andre/d15/input.txt")).stream().flatMap(this::split).toList();

		Pair<Integer, Integer> pos = null;

		// 4000 times
		for (int i = 1; i <= MAX / ROWS; i++) {
			System.out.println("round " + i + "/" + MAX / ROWS);
			shift = i;
			devices.forEach(Device::addToGrid);

			line_loop:
			for (int k = 0; k < grid.length; k++) {
				int sum = 0;

				row_loop:
				for (int h = 0; h < grid[k].length; h++) {
					if (grid[k][h] == 0) {
						sum++;
						pos = Pair.with(h, k);
					}

					if (sum > 1) {
						sum = 0;
						pos = null;
						break row_loop;
					}
				}
				if (sum == 1) break line_loop;
			}

			if (pos != null) {
				BigInteger result = BigInteger.valueOf(4000000).multiply(BigInteger.valueOf(pos.getValue0())).add(BigInteger.valueOf(pos.getValue1()));
				System.out.println("possible beacon at " + pos.getValue0() + ", " + pos.getValue1() + " result: " + result);
				break;
			}
		}

		System.out.println("done");
	}

	void run() throws IOException {
		List<Device> devices = Files.readAllLines(Paths.get("andre/d15/input.txt")).stream().flatMap(this::split).toList();

		Set<Long> xPositions = new HashSet<>();
		devices.stream()
				.filter(d -> d.radarHitsRow(ROW))
				.map(d -> ((Sensor) d).radarForRow(ROW))
				.forEach(r -> LongStream.rangeClosed(r.getMinimum(), r.getMaximum()).forEach(xPositions::add));

		long devicesOnRow = devices.stream()
				.filter(d -> d.y == ROW && xPositions.contains(d.x))
				.map(d -> Pair.with(d.x, d.y))
				.distinct().count();

		long sum = xPositions.size() - devicesOnRow;
//		5073496
		System.out.println("sum: " + sum);
	}

	Stream<Device> split(String s) {
		List<Integer> finds = new ArrayList<>();
		Matcher matcher = Pattern.compile("-?\\d{1,20}").matcher(s);
		while (matcher.find()) finds.add(Integer.parseInt(matcher.group()));

		Sensor sensor = new Sensor(finds.get(0), finds.get(1));
		Beacon beacon = new Beacon(finds.get(2), finds.get(3));
		sensor.closest = beacon;
		return Stream.of(sensor, beacon);
	}

	abstract class Device {
		long x;
		long y;

		Device(int x, int y) {
			this.x = x;
			this.y = y;
		}

		abstract byte toByte();

		abstract boolean radarHitsRow(long row);

		void addToGrid() {
			long upper = shift * ROWS;
			long lower = upper - ROWS;
			Range<Long> yRange = Range.between(lower, upper - 1);
			if (X_RANGE.contains(x) && yRange.contains(y)) {
				grid[(int) y][(int) x] = toByte();
			}
		}
	}

	class Beacon extends Device {
		Beacon(int x, int y) {
			super(x, y);
		}

		@Override
		byte toByte() {
			return 2;
		}

		@Override
		public String toString() {
			return "B";
		}

		@Override
		boolean radarHitsRow(long row) {
			return false;
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
		byte toByte() {
			return 1;
		}

		long distance() {
			return Math.abs(x - closest.x) + Math.abs(y - closest.y);
		}

		@Override
		boolean radarHitsRow(long row) {
			long distance = distance();
			return Range.between(y - distance, y + distance).contains(row);
		}

		@Override
		void addToGrid() {
			super.addToGrid();
			addRadar();
		}

		void addRadar() {
			long distance = distance();

			long upper = shift * ROWS;
			long lower = upper - ROWS;
			Range<Long> yRange = Range.between(lower, upper - 1);

			for (long i = 0; i < distance + 1; i++) {
				long row1 = y - distance + i;
				long row2 = y + distance - i;

				if (yRange.contains(row1) || yRange.contains(row2)) {
					LongStream.rangeClosed(X_RANGE.fit(x - i), X_RANGE.fit(x + i)).forEach(v -> {
						if (yRange.contains(row1)) setRadar((int) v, (int) row1);
						if (yRange.contains(row2)) setRadar((int) v, (int) row2);
					});
				}
			}

		}

		void setRadar(int x, int y) {
			if (grid[y][x] == 0) grid[y][x] = 3;
		}

		Range<Long> radarForRow(long row) {
			long distance = distance();

			for (long i = 0; i < distance + 1; i++) {
				long row1 = y - distance + i;
				long row2 = y + distance - i;
				if ((row1 == row) || (row2 == row)) {
					return Range.between(x - i, x + i);
				}
			}
			throw new IllegalArgumentException("radar not in row " + row);
		}
	}
}