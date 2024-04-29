package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		a        string
		b        string
		expected string
	}{
		{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
	}

	for _, tc := range testCases {
		result := addBinary(tc.a, tc.b)
		require.Equal(t, result, tc.expected)
	}
}
