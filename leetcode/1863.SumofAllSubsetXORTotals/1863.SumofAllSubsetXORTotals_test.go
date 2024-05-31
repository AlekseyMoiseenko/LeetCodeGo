package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubsetXORSum(t *testing.T) {
	testCases := []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{1, 3},
			expected: 6,
		},
		{
			list:     []int{5, 1, 6},
			expected: 28,
		},
	}

	for _, tc := range testCases {
		res := subsetXORSum(tc.list)

		require.Equal(t, tc.expected, res)
	}
}
