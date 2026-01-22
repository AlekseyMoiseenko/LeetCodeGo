package main

const (
	maxVal = int(1e5)
)

func minimumPairRemoval(nums []int) int {
	opCount := 0

	for !isSorted(nums) {
		minSum := maxVal
		idx := 0
		for i := 0; i+1 < len(nums); i++ {
			if nums[i]+nums[i+1] < minSum {
				minSum = nums[i] + nums[i+1]
				idx = i
			}
		}
		nums[idx] = minSum
		nums = append(nums[:idx+1], nums[idx+2:]...)
		opCount++
	}
	return opCount
}

func isSorted(x []int) bool {
	for i := len(x) - 1; i > 0; i-- {
		if x[i] < x[i-1] {
			return false
		}
	}
	return true
}
