package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeightChecker(t *testing.T) {
	testCases := []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{1, 1, 4, 2, 1, 3},
			expected: 3,
		},
		{
			list:     []int{5, 1, 2, 3, 4},
			expected: 5,
		},
		{
			list:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		res := heightChecker(tc.list)

		require.Equal(t, tc.expected, res)
	}
}
