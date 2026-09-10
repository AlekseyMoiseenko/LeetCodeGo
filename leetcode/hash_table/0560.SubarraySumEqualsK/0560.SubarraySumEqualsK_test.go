package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= nums.length <= 2 * 10^4
// -1000 <= nums[i] <= 1000
// -10^7 <= k <= 10^7
func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "SubarraySum test 1",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "SubarraySum test 2",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			name:     "SubarraySum test 3",
			nums:     []int{-1, -1, -1},
			k:        1,
			expected: 0,
		},
		{
			name:     "SubarraySum test 4",
			nums:     []int{-1, 0, 1},
			k:        0,
			expected: 2,
		},
		{
			name:     "SubarraySum test 5",
			nums:     []int{-10, -5, 1, 1, 1, 1, 1},
			k:        0,
			expected: 1,
		},
		{
			name:     "SubarraySum test 6",
			nums:     []int{3, 4, 7, -3, -3, 2},
			k:        7,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := subarraySum(tc.nums, tc.k)
			require.Equal(t, tc.expected, res)
		})
	}
}
