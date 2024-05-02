package leetcode

func findMaxK(nums []int) int {
	max := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] > max && containsNegative(nums, nums[i]) {
			max = nums[i]
		}
	}

	return max
}

func containsNegative(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == -num {
			return true
		}
	}

	return false
}
