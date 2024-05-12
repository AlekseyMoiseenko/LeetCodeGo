package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestLocal(t *testing.T) {
	testCases := []struct {
		list     [][]int
		expected [][]int
	}{
		{
			list:     [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}},
			expected: [][]int{{9, 9}, {8, 6}},
		},
		{
			list:     [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
			expected: [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
		},
	}

	for _, tc := range testCases {
		res := largestLocal(tc.list)

		require.Equal(t, res, tc.expected)
	}
}
