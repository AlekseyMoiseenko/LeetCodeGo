package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinIncrementForUnique(t *testing.T) {
	testCases := []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{2, 2, 2, 2, 2},
			expected: 10,
		},
		{
			list:     []int{1, 2, 2},
			expected: 1,
		},
		{
			list:     []int{3, 2, 1, 2, 1, 7},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		res := minIncrementForUnique(tc.list)

		require.Equal(t, tc.expected, res)
	}
}
