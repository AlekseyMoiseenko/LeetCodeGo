package leetcode

// nums[i] is either 0, 1, or 2.
const (
	max = 2
	min = 0
)

// Time Complexity: O(n)
// Space Complexity: O(1)
func sortColors(nums []int) {
	counter := make([]int, max-min+1)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]-min]++
	}

	index := 0
	for i, v := range counter {
		for j := 0; j < v; j++ {
			nums[index] = i + min
			index++
		}
	}
}
