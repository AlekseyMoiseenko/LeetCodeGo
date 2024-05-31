package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumHappinessSum(t *testing.T) {
	testCases := []struct {
		list     []int
		turns    int
		expected int64
	}{
		{
			list:     []int{2, 98, 45},
			turns:    1,
			expected: 98,
		},
		{
			list:     []int{1, 2, 3},
			turns:    2,
			expected: 4,
		},
		{
			list:     []int{1, 1, 1, 1},
			turns:    2,
			expected: 1,
		},
		{
			list:     []int{2, 3, 4, 5},
			turns:    1,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		res := maximumHappinessSum(tc.list, tc.turns)

		require.Equal(t, tc.expected, res)
	}
}
