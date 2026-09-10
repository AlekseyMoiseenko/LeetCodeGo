package main

// Time complexity: O(n)
// Space complexity: O(n)
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	prefixSum := 0
	m[prefixSum]++

	res := 0
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		res += m[prefixSum-k]
		m[prefixSum]++
	}

	return res
}
