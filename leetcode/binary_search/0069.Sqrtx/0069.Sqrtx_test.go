package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 0 <= x <= 2^31 - 1
func TestMySqrt(t *testing.T) {
	testCases := []struct {
		x        int
		expected int
	}{
		{
			x:        8,
			expected: 2,
		},
		{
			x:        9,
			expected: 3,
		},
		{
			x:        4,
			expected: 2,
		},
		{
			x:        3,
			expected: 1,
		},
		{
			x:        0,
			expected: 0,
		},
		{
			x:        2147483648,
			expected: 46340,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, mySqrt(tc.x))
	}
}
