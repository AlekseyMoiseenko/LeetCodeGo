package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkobochki(t *testing.T) {
	testCases := []struct {
		n        int
		expected []string
	}{
		{
			n:        1,
			expected: []string{"()"},
		},
		{
			n:        2,
			expected: []string{"()()", "(())"},
		},
		{
			n:        3,
			expected: []string{"()()()", "()(())", "(())()", "(()())", "((()))"},
		},
	}

	for _, tc := range testCases {
		res := generate(tc.n)
		sort.Slice(res, func(i, j int) bool {
			return i > j
		})
		// sort.Strings(res)
		// slices.Reverse(res)
		require.Equal(t, tc.expected, res)
	}
}
