package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmallestNumber(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{
			n:        5,
			expected: 7,
		},
		{
			n:        10,
			expected: 15,
		},
		{
			n:        3,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		res := smallestNumber(tc.n)

		require.Equal(t, tc.expected, res)
	}
}
