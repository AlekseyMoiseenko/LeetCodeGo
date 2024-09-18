package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestNumber(t *testing.T) {
	testCases := []struct {
		list     []int
		expected string
	}{
		{
			list:     []int{5, 30, 34, 3, 9},
			expected: "9534330",
		},
		{
			list:     []int{10, 2},
			expected: "210",
		},
		{
			list:     []int{0, 0, 0, 0, 0, 0},
			expected: "0",
		},
		{
			list:     []int{0, 0, 0, 1, 0, 0},
			expected: "100000",
		},
		{
			list:     []int{0, 1000000000, 0, 1, 0, 0},
			expected: "110000000000000",
		},
	}
	for _, tc := range testCases {
		require.Equal(t, tc.expected, largestNumber(tc.list))
	}
}
