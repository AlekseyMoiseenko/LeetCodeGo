package hash

import "sort"

func containsDuplicate(nums []int) bool {
	// Time Complexity: O(n log n)
	// Space Complexity: O(log n)
	sort.Ints(nums)

	// Time Complexity: O(n)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

// first approach
// Time Complexity: O(n)
// Space Complexity: O(n)
func containsDuplicate1(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > 1 {
			return true
		}
	}

	return false
}

// first approach improvement
// Time Complexity: O(n)
// Space Complexity: O(n)
func containsDuplicate2(nums []int) bool {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, exist := m[v]; exist {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}
