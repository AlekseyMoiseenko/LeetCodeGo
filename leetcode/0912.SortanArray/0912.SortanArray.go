package leetcode

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}

	numsCount := make([]int, maxVal-minVal+1)
	for _, num := range nums {
		numsCount[num-minVal]++
	}

	index := 0
	for v, count := range numsCount {
		for count > 0 {
			nums[index] = v + minVal
			index++
			count--
		}
	}

	return nums
}
