package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		s        string
		expected int
	}{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
		{
			s:        "",
			expected: 0,
		},
		{
			s:        "dvdf",
			expected: 3,
		},
		{
			s:        "ckilbkd",
			expected: 5,
		},
		{
			s:        "tmmzuxt",
			expected: 5,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, lengthOfLongestSubstring(tc.s))
	}
}
