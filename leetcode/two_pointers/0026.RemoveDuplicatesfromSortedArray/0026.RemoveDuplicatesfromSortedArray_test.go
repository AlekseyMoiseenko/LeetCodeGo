package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= nums.length <= 3 * 10^4
// -100 <= nums[i] <= 100
// nums is sorted in non-decreasing order.
func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums         []int
		expected     int
		expectedNums []int
	}{
		{
			nums:         []int{5},
			expected:     1,
			expectedNums: []int{5},
		},
		{
			nums:         []int{1, 2, 3},
			expected:     3,
			expectedNums: []int{1, 2, 3},
		},
		{
			nums:         []int{1, 1, 1, 1, 5, 5, 5},
			expected:     2,
			expectedNums: []int{1, 5},
		},
		{
			nums:         []int{-100, 100},
			expected:     2,
			expectedNums: []int{-100, 100},
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected:     5,
			expectedNums: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		k := removeDuplicates(tc.nums)
		require.Equal(t, tc.expected, k)
		require.Equal(t, tc.expectedNums, tc.nums[:k])
	}
}
