package hash

// Time complexity: O(n)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		if idx, exist := m[target-v]; exist {
			return []int{idx, i}
		}

		m[v] = i
	}

	return nil
}
