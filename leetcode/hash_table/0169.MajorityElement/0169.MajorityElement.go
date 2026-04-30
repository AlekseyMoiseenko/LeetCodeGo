package hash

import (
	"sort"
)

// Time complexity: O(n log n). (O(n log n) - sort, O(n) - array iterating)
// Space complexity: O(1)
func majorityElement(nums []int) int {
	sort.Ints(nums)

	key, maxCount, count := nums[0], 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count > maxCount {
			key = nums[i]
			maxCount = count
		}
	}

	return key
}

// first approach
// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	for _, n := range nums {
// 		m[n]++
// 	}

// 	maxK, maxV := 0, 0
// 	for k, v := range m {
// 		if maxV < v {
// 			maxK = k
// 			maxV = v
// 		}
// 	}

// 	return maxK
// }
