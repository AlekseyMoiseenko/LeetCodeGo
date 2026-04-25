package hash

// Time complexity: O(k log n)
// where k is the number of iterations until we find a cycle or reach 1, and log n represents the number of digits in n. Each iteration processes all digits of the current number.

// Space complexity: O(k)
// where k is the number of unique numbers we encounter before finding a cycle or reaching 1. We store all these numbers in the set.
func isHappy(n int) bool {
	repeatMap := make(map[int]struct{})

	for {
		if n == 1 {
			return true
		}

		repeatMap[n] = struct{}{}

		digits := splitNumber(n)
		n = 0
		for _, v := range digits {
			n += v * v
		}

		if _, exist := repeatMap[n]; exist {
			return false
		}
	}
}

func splitNumber(n int) []int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	return digits
}
