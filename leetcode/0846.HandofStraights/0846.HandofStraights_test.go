package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNStraightHand(t *testing.T) {
	testCases := []struct {
		list      []int
		groupSize int
		expected  bool
	}{
		{
			list:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize: 3,
			expected:  true,
		},
		{
			list:      []int{1, 1, 2, 2, 3, 3},
			groupSize: 2,
			expected:  false,
		},
		{
			list:      []int{8, 10, 12},
			groupSize: 3,
			expected:  false,
		},
		{
			list:      []int{2, 1, 2, 4, 1, 3, 3, 3},
			groupSize: 2,
			expected:  false,
		},
	}

	for _, tc := range testCases {
		res := isNStraightHand(tc.list, tc.groupSize)

		require.Equal(t, tc.expected, res)
	}
}
