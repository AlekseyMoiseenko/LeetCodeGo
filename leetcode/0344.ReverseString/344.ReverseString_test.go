package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		s        []byte
		expected []byte
	}{
		{
			s:        []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:        []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, tc := range testCases {
		reverseString(tc.s)

		require.Equal(t, tc.expected, tc.s)
	}
}
