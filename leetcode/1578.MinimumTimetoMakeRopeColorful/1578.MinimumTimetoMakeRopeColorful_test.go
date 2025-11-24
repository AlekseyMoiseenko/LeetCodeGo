package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinCost(t *testing.T) {
	testCases := []struct {
		colors     string
		neededTime []int
		expected   int
	}{
		{
			colors:     "abaac",
			neededTime: []int{1, 2, 3, 4, 5},
			expected:   3,
		},
		{
			colors:     "abc",
			neededTime: []int{1, 2, 3},
			expected:   0,
		},
		{
			colors:     "aabaa",
			neededTime: []int{1, 2, 3, 4, 1},
			expected:   2,
		},
		{
			colors:     "a",
			neededTime: []int{10},
			expected:   0,
		},
		{
			colors:     "aaaaa",
			neededTime: []int{5, 4, 6, 3, 7},
			expected:   18,
		},
	}

	for _, tc := range testCases {
		res := minCost(tc.colors, tc.neededTime)

		require.Equal(t, tc.expected, res)
	}
}
