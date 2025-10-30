package leetcode

func minNumberOperations(target []int) int {
	if len(target) == 0 {
		return 0
	}

	res := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			res += target[i] - target[i-1]
		}
	}

	return res
}
