package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareVersion(t *testing.T) {
	testCases := []struct {
		v1       string
		v2       string
		expected int
	}{
		{
			v1:       "1.01",
			v2:       "1.001",
			expected: 0,
		},
		{
			v1:       "1.0",
			v2:       "1.0.0",
			expected: 0,
		},
		{
			v1:       "0.1",
			v2:       "1.1",
			expected: -1,
		},
		{
			v1:       "1.0.1",
			v2:       "1",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		res := compareVersion(tc.v1, tc.v2)
		require.Equal(t, tc.expected, res)
	}
}
