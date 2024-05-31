package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKthSmallestPrimeFraction(t *testing.T) {
	testCases := []struct {
		list     []int
		k        int
		expected []int
	}{
		{
			list:     []int{1, 2, 3, 5},
			k:        3,
			expected: []int{2, 5},
		},
		{
			list:     []int{1, 7},
			k:        1,
			expected: []int{1, 7},
		},
	}

	for _, tc := range testCases {
		res := kthSmallestPrimeFraction(tc.list, tc.k)
		require.Equal(t, tc.expected, res)
	}
}
