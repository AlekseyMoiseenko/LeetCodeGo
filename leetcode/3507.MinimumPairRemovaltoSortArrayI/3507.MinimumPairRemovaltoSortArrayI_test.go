package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// 1 <= nums.length <= 50
// -1000 <= nums[i] <= 1000
func TestMinimumPairRemoval(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{5, 2, 3, 1},
			expected: 2,
		},
		{
			nums:     []int{1, 2, 2},
			expected: 0,
		},
		{
			nums:     []int{1, 2, 1},
			expected: 2,
		},
		{
			nums:     []int{1, 5, 3, 2},
			expected: 1,
		},
		{
			nums:     []int{1, 5, 3, 2, 5},
			expected: 1,
		},
		{
			nums:     []int{1, 5, 3, 2, 5, 1, 1, 3},
			expected: 3,
		},
		{
			nums:     []int{1, 1, 4, 4, 2, -4, -1},
			expected: 5,
		},
		{
			nums:     []int{2, 2, -1, 3, -2, 2, 1, 1, 1, 0, -1},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, minimumPairRemoval(tc.nums))
	}
}
