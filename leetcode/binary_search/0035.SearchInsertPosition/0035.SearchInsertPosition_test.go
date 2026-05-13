package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= nums.length <= 10^4
// -10^4 <= nums[i] <= 10^4
// nums contains distinct values sorted in ascending order.
// -10^4 <= target <= 10^4
func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums     []int
		val      int
		expected int
	}{
		{
			nums:     []int{1},
			val:      2,
			expected: 1,
		},
		{
			nums:     []int{1},
			val:      -1,
			expected: 0,
		},
		{
			nums:     []int{1, 3, 5, 6},
			val:      5,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 5, 6},
			val:      2,
			expected: 1,
		},
		{
			nums:     []int{1, 3, 5, 6},
			val:      7,
			expected: 4,
		},
		{
			nums:     []int{-6, -5, -3, -1, 0, 1, 3, 5, 6},
			val:      -4,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, searchInsert(tc.nums, tc.val))
	}
}
