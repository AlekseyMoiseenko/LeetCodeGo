package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScoreOfString(t *testing.T) {
	testCases := []struct {
		str      string
		expected int
	}{
		{
			str:      "hello",
			expected: 13,
		},
		{
			str:      "zaz",
			expected: 50,
		},
	}

	for _, tc := range testCases {
		res := scoreOfString(tc.str)

		require.Equal(t, tc.expected, res)
	}
}
