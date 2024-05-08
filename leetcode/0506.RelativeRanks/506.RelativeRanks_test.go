package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindRelativeRanks(t *testing.T) {
	testCases := []struct {
		list     []int
		expected []string
	}{
		{
			list:     []int{10, 3, 8, 9, 4},
			expected: []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
		{
			list:     []int{5, 4, 3, 2, 1},
			expected: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
	}

	for _, tc := range testCases {
		res := findRelativeRanks(tc.list)
		require.Equal(t, res, tc.expected)
	}
}
