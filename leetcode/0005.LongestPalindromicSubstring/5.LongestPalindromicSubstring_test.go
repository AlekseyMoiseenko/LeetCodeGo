package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{
			str:      "babaddaba",
			expected: "abaddaba",
		},
		{
			str:      "babad",
			expected: "bab",
		},
		{
			str:      "cbbd",
			expected: "bb",
		},
	}

	for _, tc := range testCases {
		result := longestPalindrome(tc.str)
		require.Equal(t, result, tc.expected)
	}
}
