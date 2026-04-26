package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			nums:     []int{1},
			expected: false,
		},
		{
			nums:     []int{-100000000},
			expected: false,
		},
		{
			nums:     []int{100000000},
			expected: false,
		},
		{
			nums:     []int{-100000000, 0, 100000000},
			expected: false,
		},
		{
			nums:     []int{-100000000, 0, 100000000, 9999999, 99998, 99888, -100000000},
			expected: true,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, containsDuplicate(tc.nums))
		require.Equal(t, tc.expected, containsDuplicate1(tc.nums))
		require.Equal(t, tc.expected, containsDuplicate2(tc.nums))
	}
}
