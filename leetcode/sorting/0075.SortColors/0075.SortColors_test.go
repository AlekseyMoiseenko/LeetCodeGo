package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// n == nums.length
// 1 <= n <= 300
// nums[i] is either 0, 1, or 2.
func TestSortColors(t *testing.T) {
	testCases := []struct {
		colors   []int
		expected []int
	}{
		{
			colors:   []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			colors:   []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			colors:   []int{1},
			expected: []int{1},
		},
		{
			colors:   []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
			expected: []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			colors:   []int{2, 2, 2, 0, 0, 0, 1, 1, 1},
			expected: []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
	}

	for _, tc := range testCases {
		sortColors(tc.colors)
		require.Equal(t, tc.expected, tc.colors)
	}
}
