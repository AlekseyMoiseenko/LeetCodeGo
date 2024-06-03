package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendCharacters(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected int
	}{
		{
			s:        "coaching",
			t:        "coding",
			expected: 4,
		},
		{
			s:        "z",
			t:        "abcde",
			expected: 5,
		},
		{
			s:        "abcde",
			t:        "a",
			expected: 0,
		},
		{
			s:        "lbg",
			t:        "g",
			expected: 0,
		},
		{
			s:        "grcedr",
			t:        "cedrgr",
			expected: 2,
		},
	}

	for _, tc := range testCases {
		res := appendCharacters(tc.s, tc.t)

		require.Equal(t, tc.expected, res)
	}
}
