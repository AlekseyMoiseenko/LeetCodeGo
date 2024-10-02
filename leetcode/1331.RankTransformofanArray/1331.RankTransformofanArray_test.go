package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayRankTransform(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{40, 10, 20, 30},
			expected: []int{4, 1, 2, 3},
		},
		{
			arr:      []int{100, 100, 100},
			expected: []int{1, 1, 1},
		},
		{
			arr:      []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			expected: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
		{
			arr:      []int{40, 10, 20, 30, 1, 1, 1, 10, 20, 30, 100, 1000, 1000000000},
			expected: []int{5, 2, 3, 4, 1, 1, 1, 2, 3, 4, 6, 7, 8},
		},
		{
			arr:      []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, arrayRankTransform(tc.arr))
	}
}
