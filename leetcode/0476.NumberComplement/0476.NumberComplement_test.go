package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindComplement(t *testing.T) {
	testCases := []struct {
		val      int
		expected int
	}{
		{
			val:      5,
			expected: 2,
		},
		{
			val:      1,
			expected: 0,
		},
		{
			val:      144,
			expected: 111,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, findComplement(tc.val))
	}
}
