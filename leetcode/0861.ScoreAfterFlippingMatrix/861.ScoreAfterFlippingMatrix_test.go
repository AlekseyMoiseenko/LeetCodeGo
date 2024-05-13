package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatrixScore(t *testing.T) {
	testCases := []struct {
		list     [][]int
		expected int
	}{
		{
			list:     [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}},
			expected: 39,
		},
		{
			list:     [][]int{{0}},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		res := matrixScore(tc.list)

		require.Equal(t, res, tc.expected)
	}
}
