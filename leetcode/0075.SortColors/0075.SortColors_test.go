package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			list:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			list:     []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		sortColors(tc.list)
		require.Equal(t, tc.expected, tc.list)
	}
}
