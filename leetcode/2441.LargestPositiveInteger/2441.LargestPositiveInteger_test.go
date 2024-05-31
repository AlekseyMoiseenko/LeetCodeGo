package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMaxK(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{-1, 10, 6, 7, -7, 1},
			expected: 7,
		},
		{
			nums:     []int{-1, 2, -3, 3},
			expected: 3,
		},
		{
			nums:     []int{-10, 8, 6, 7, -2, -3},
			expected: -1,
		},
	}
	for _, tc := range testCases {
		res := findMaxK(tc.nums)
		require.Equal(t, tc.expected, res)
	}
}
