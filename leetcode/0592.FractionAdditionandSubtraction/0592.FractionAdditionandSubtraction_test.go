package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFractionAddition(t *testing.T) {
	testCases := []struct {
		expression string
		expected   string
	}{
		{
			expression: "-1/2+1/2+1/3",
			expected:   "1/3",
		},
		{
			expression: "-1/2+1/2",
			expected:   "0/1",
		},
		{
			expression: "1/3-1/2",
			expected:   "-1/6",
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, fractionAddition(tc.expression))
	}
}
