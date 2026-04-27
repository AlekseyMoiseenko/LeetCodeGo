package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= s.length, t.length <= 5 * 10^4
// s and t consist of lowercase English letters.
func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			s:        "r",
			t:        "c",
			expected: false,
		},
		{
			s:        "c",
			t:        "c",
			expected: true,
		},
		{
			s:        "ac",
			t:        "ca",
			expected: true,
		},
		{
			s:        "a",
			t:        "ab",
			expected: false,
		},
		{
			s:        "ba",
			t:        "a",
			expected: false,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, isAnagram(tc.s, tc.t))
		require.Equal(t, tc.expected, isAnagram1(tc.s, tc.t))
	}
}
