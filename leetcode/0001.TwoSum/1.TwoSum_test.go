package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		list     []int
		target   int
		expected []int
	}{
		{
			list:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			list:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			list:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			list:     []int{1, 2, 3, 4},
			target:   15,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		indices := twoSum(tc.list, tc.target)
		require.Equal(t, tc.expected, indices)
	}
}
