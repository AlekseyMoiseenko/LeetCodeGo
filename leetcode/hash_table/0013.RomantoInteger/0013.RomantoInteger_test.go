package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= s.length <= 15
// s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
// It is guaranteed that s is a valid roman numeral in the range [1, 3999].
func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		s        string
		expected int
	}{
		{
			s:        "III",
			expected: 3,
		},
		{
			s:        "LVIII",
			expected: 58,
		},
		{
			s:        "MCMXCIV",
			expected: 1994,
		},
		{
			s:        "MMMDCCCLXXXVIII",
			expected: 3888,
		},
		{
			s:        "CDXLIV",
			expected: 444,
		},
		{
			s:        "I",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, romanToInt(tc.s))
	}
}
