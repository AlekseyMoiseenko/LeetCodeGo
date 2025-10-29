package leetcode

func smallestNumber(n int) int {
	bitSet := 1
	for bitSet < n {
		bitSet = bitSet*2 + 1
	}
	return bitSet
}
