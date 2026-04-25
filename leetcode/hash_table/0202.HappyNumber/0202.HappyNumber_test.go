package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= n <= 2^31 - 1
func TestIsHappy(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{
			n:        19,
			expected: true,
		},
		{
			n:        2,
			expected: false,
		},
		{
			n:        2147483647,
			expected: false,
		},
		{
			n:        1,
			expected: true,
		},
		{
			n:        10,
			expected: true,
		},
		{
			n:        101,
			expected: false,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, isHappy(tc.n))
	}
}
