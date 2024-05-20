package leetcode

func subsetXORSum(nums []int) int {
	return subset(nums, 0)
}

func subset(nums []int, n int) int {
	result := n

	for i, v := range nums {
		result += subset(nums[i+1:], n^v)
	}

	return result
}
