package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortestPalindrome(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{
			str:      "abcd",
			expected: "dcbabcd",
		},
		{
			str:      "aacecaaa",
			expected: "aaacecaaa",
		},
		{
			str:      "abcdabcd",
			expected: "dcbadcbabcdabcd",
		},
		{
			str:      "abcbbbcd",
			expected: "dcbbbcbabcbbbcd",
		},
		{
			str:      "abcba",
			expected: "abcba",
		},
		{
			str:      "abbacd",
			expected: "dcabbacd",
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, shortestPalindrome(tc.str))
	}
}
