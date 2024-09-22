package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindKthNumber(t *testing.T) {
	testCases := []struct {
		n        int
		k        int
		expected int
	}{
		{
			n:        13,
			k:        2,
			expected: 10,
		},
		{
			n:        13,
			k:        1,
			expected: 1,
		},
		{
			n:        1,
			k:        1,
			expected: 1,
		},
		{
			n:        13,
			k:        3,
			expected: 11,
		},
		{
			n:        100000,
			k:        3,
			expected: 100,
		},
		{
			n:        599765895,
			k:        511111111,
			expected: 559999999,
		},
		{
			n:        599765895,
			k:        511111112,
			expected: 56,
		},
		{
			n:        599765895,
			k:        511111113,
			expected: 560,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, findKthNumber(tc.n, tc.k))
	}
}
