package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLexicalOrder(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{
			n:        13,
			expected: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			n:        2,
			expected: []int{1, 2},
		},
		{
			n:        1,
			expected: []int{1},
		},
		{
			n:        100,
			expected: []int{1, 10, 100, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 3, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 4, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 5, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 6, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 7, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 8, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 9, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, lexicalOrder(tc.n))
	}
}
