package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 2 <= nums.length <= 10^4
// -10^9 <= nums[i] <= 10^9
// -10^9 <= target <= 10^9
// Only one valid answer exists.
func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			nums:     []int{1, 2, 3, 4},
			target:   15,
			expected: nil,
		},
		{
			nums:     []int{-1000000000, -10, 0, 3, 1000000000},
			target:   0,
			expected: []int{0, 4},
		},
		{
			nums:     []int{-1000000000, -10, 0, 3, 1000000000},
			target:   -10,
			expected: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, twoSum(tc.nums, tc.target))
	}
}
