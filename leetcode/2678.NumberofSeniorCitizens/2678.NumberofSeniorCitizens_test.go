package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountSeniors(t *testing.T) {
	testCases := []struct {
		list     []string
		expected int
	}{
		{
			list:     []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"},
			expected: 2,
		},
		{
			list:     []string{"1313579440F2036", "2921522980M5644"},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		res := countSeniors(tc.list)
		require.Equal(t, tc.expected, res)
	}
}
