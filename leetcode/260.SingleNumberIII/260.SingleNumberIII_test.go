package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{1, 2, 1, 3, 2, 5},
			expected: []int{3, 5},
		},
		{
			list:     []int{-1, 0},
			expected: []int{-1, 0},
		},
		{
			list:     []int{0, 1},
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		res := singleNumber(tc.list)

		require.Equal(t, tc.expected, res)
	}
}
