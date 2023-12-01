import java.io.IOException;
import java.math.BigInteger;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Objects;
import java.util.UUID;

public class Main20 {

	public static void main(String[] args) throws IOException {
		new Main20().run();
	}

	void run() throws IOException {
		List<String> input = Files.readAllLines(Paths.get("andre/d20/input.txt"));
		List<Entry> data = input.stream()
				.map(Integer::parseInt).map(Entry::new).toList();
//		System.out.println("before: " + data);
		List<Entry> modified = shiftValues(data, 1);
//		System.out.println("after:  " + modified);

		Entry zero = modified.stream().filter(v -> v.value == 0).findFirst().get();
		long sum = getSum(modified, zero);
		// 5962
		System.out.println("part 1 sum: " + sum);

		modified = shiftValues(data, 10);
		zero = modified.stream().filter(v -> v.value == 0).findFirst().get();
		BigInteger bigSum = getBigSum(modified, zero);
		// 9862431387256
		System.out.println("part 2 sum: " + bigSum);
	}

	private long getSum(List<Entry> modified, Entry zero) {
		int zeroIndex = modified.indexOf(zero) + 1;
		long one = getEntry(1000, zeroIndex, modified).value;
		long two = getEntry(2000, zeroIndex, modified).value;
		long three = getEntry(3000, zeroIndex, modified).value;
		return one + two + three;
	}

	private BigInteger getBigSum(List<Entry> modified, Entry zero) {
		int zeroIndex = modified.indexOf(zero) + 1;
		BigInteger one = getEntry(1000, zeroIndex, modified).getBig();
		BigInteger two = getEntry(2000, zeroIndex, modified).getBig();
		BigInteger three = getEntry(3000, zeroIndex, modified).getBig();
		return one.add(two).add(three);
	}

	private Entry getEntry(int position, int zeroIndex, List<Entry> modified) {
		int index = (position + zeroIndex) % modified.size();
		return modified.get(index - 1);
	}

	private List<Entry> shiftValues(List<Entry> data, int times) {
		List<Entry> modify = new ArrayList<>(data);
		for (int i = 0; i < times; i++) {
			for (var e : data) {
				int index = modify.indexOf(e);
				int newIndex = times > 1 ? e.getNewIndex(index, modify.size() - 1) : Math.floorMod(index + e.value, modify.size() - 1);
				if (newIndex == 0) newIndex = modify.size() - 1;
				int min = Math.min(index, newIndex);
				int max = Math.max(index, newIndex);
				List<Entry> entry = modify.subList(min, max + 1);
				Collections.rotate(entry, index < newIndex ? -1 : 1);
			}
		}
		return modify;
	}

	// this is only needed because values are not unique
	class Entry {
		final long value;
		final UUID id;
		final BigInteger big;

		Entry(long value) {
			this.value = value;
			this.id = UUID.randomUUID();
			this.big = BigInteger.valueOf(value).multiply(BigInteger.valueOf(811589153));
		}

		@Override
		public boolean equals(Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;
			Entry entry = (Entry) o;
			return id.equals(entry.id);
		}

		@Override
		public int hashCode() {
			return Objects.hash(id);
		}

		@Override
		public String toString() {
			return "" + value;
		}

		public int getNewIndex(int index, int modValue) {
			return big.add(BigInteger.valueOf(index)).mod(BigInteger.valueOf(modValue)).intValue();
		}

		public BigInteger getBig() {
			return big;
		}
	}
}
