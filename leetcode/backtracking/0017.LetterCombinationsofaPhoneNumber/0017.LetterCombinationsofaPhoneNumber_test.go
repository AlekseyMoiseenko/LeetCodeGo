package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].
func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		digits   string
		expected []string
	}{
		{
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, letterCombinations(tc.digits))
	}
}
