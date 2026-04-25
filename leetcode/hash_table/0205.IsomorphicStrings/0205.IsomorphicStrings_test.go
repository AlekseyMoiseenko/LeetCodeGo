package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 1 <= s.length <= 5 * 10^4
// t.length == s.length
// s and t consist of any valid ascii character.
func TestIsIsomorphic(t *testing.T) {
	testCases := []struct {
		initStr   string
		resultStr string
		expected  bool
	}{
		{
			initStr:   "egg",
			resultStr: "add",
			expected:  true,
		},
		{
			initStr:   "f11",
			resultStr: "b23",
			expected:  false,
		},
		{
			initStr:   "paper",
			resultStr: "title",
			expected:  true,
		},
		{
			initStr:   "egg",
			resultStr: "egg",
			expected:  true,
		},
		{
			initStr:   "pppee",
			resultStr: "ggppp",
			expected:  false,
		},
		{
			initStr:   "1573102ab0",
			resultStr: "1573102ab0",
			expected:  true,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, isIsomorphic(tc.initStr, tc.resultStr))
	}
}
