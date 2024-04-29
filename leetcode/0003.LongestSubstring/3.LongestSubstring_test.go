package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		str      string
		expected int
	}{
		{
			str:      "abcabcbb",
			expected: 3,
		},
		{
			str:      "bbbbb",
			expected: 1,
		},
		{
			str:      "pwwkew",
			expected: 3,
		},
	}

	for _, tc := range testCases {
		v := lengthOfLongestSubstring(tc.str)
		require.Equal(t, v, tc.expected)
	}
}
