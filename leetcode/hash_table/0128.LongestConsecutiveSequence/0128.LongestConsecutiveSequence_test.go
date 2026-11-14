package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "empty slice",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "nil slice",
			nums:     nil,
			expected: 0,
		},
		{
			name:     "single element",
			nums:     []int{42},
			expected: 1,
		},
		{
			name:     "classic example",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "with duplicates",
			nums:     []int{1, 2, 0, 1},
			expected: 3,
		},
		{
			name:     "only duplicates",
			nums:     []int{5, 5, 5, 5},
			expected: 1,
		},
		{
			name:     "no consecutive",
			nums:     []int{1, 3, 5, 7, 9},
			expected: 1,
		},
		{
			name:     "all negative consecutive",
			nums:     []int{-3, -2, -1, -100},
			expected: 3,
		},
		{
			name:     "crosses zero",
			nums:     []int{-1, 0, 1, 2},
			expected: 4,
		},
		{
			name:     "full sequence",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "two separate sequences, pick longer",
			nums:     []int{1, 2, 10, 11, 12, 13},
			expected: 4,
		},
		{
			name:     "min and max constraint values",
			nums:     []int{-1_000_000_000, 1_000_000_000},
			expected: 1,
		},
		{
			name:     "sequence at min boundary",
			nums:     []int{-1_000_000_000, -999_999_999, -999_999_998},
			expected: 3,
		},
		{
			name:     "sequence at max boundary",
			nums:     []int{999_999_998, 999_999_999, 1_000_000_000},
			expected: 3,
		},
		{
			name:     "unsorted with duplicates and gaps",
			nums:     []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := longestConsecutive(tc.nums)
			require.Equal(t, tc.expected, res,
				"longestConsecutive(%v) = %d, want %d", tc.nums, res, tc.expected)
		})
	}
}
