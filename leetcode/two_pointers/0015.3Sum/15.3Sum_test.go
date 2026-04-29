package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 3 <= nums.length <= 3000
// -10^5 <= nums[i] <= 10^5
func TestThreeSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{-10, -1, 0, 1, 11},
			expected: [][]int{{-10, -1, 11}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 1, 11, -10, -1},
			expected: [][]int{{-10, -1, 11}, {-1, 0, 1}},
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, threeSum(tc.nums))
	}
}
