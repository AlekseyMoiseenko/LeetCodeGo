package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPartition(t *testing.T) {
	testCases := []struct {
		str      string
		expected [][]string
	}{
		{
			str:      "aab",
			expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			str:      "a",
			expected: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		res := partition(tc.str)

		require.Equal(t, tc.expected, res)
	}
}
