package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		testStrs    []string
		expectedStr string
	}{
		{
			testStrs:    []string{"flower", "flow", "flight"},
			expectedStr: "fl",
		},
		{
			testStrs:    []string{"dog", "racecar", "car"},
			expectedStr: "",
		},
	}
	for _, tc := range testCases {
		res := longestCommonPrefix(tc.testStrs)
		require.Equal(t, tc.expectedStr, res)
	}
}
