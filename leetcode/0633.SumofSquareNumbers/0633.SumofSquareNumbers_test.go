package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJudgeSquareSum(t *testing.T) {
	testCases := []struct {
		val      int
		expected bool
	}{
		{
			val:      2,
			expected: true,
		},
		{
			val:      1000000000,
			expected: true,
		},
		{
			val:      3,
			expected: false,
		},
		{
			val:      5,
			expected: true,
		},
	}

	for _, tc := range testCases {
		res := judgeSquareSum(tc.val)

		require.Equal(t, tc.expected, res)
	}
}
