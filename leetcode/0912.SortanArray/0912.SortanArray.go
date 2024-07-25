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
	for v, c := range numsCount {
		for c > 0 {
			nums[index] = v + minVal
			index++
			c--
		}
	}

	return nums
}

// time limit
// func sortArray(nums []int) []int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] > nums[i+1] {
// 			nums[i+1], nums[i] = nums[i], nums[i+1]
// 			if i >= 1 {
// 				i -= 2
// 			}
// 		}
// 	}
// 	return nums
// }
