package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortArray(t *testing.T) {
	testCases := []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{5, 2, 3, 1},
			expected: []int{1, 2, 3, 5},
		},
		{
			list:     []int{5, 1, 1, 2, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tc := range testCases {
		res := sortArray(tc.list)

		require.Equal(t, tc.expected, res)
	}
}
