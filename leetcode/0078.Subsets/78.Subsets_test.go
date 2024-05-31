package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubsets(t *testing.T) {
	testCases := []struct {
		list     []int
		expected [][]int
	}{
		{
			list:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			list:     []int{0},
			expected: [][]int{{}, {0}},
		},
	}

	for _, tc := range testCases {
		res := subsets(tc.list)

		require.Equal(t, tc.expected, res)
	}
}
