package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinExtraChar(t *testing.T) {
	testCases := []struct {
		s          string
		dictionary []string
		expected   int
	}{
		{
			s:          "leetscode",
			dictionary: []string{"leet", "code", "leetcode"},
			expected:   1,
		},
		{
			s:          "sayhelloworld",
			dictionary: []string{"hello", "world"},
			expected:   3,
		},
		{
			s:          "dwmodizxvvbosxxw",
			dictionary: []string{"ox", "lb", "diz", "gu", "v", "ksv", "o", "nuq", "r", "txhe", "e", "wmo", "cehy", "tskz", "ds", "kzbu"},
			expected:   7,
		},
		{
			s:          "kevlplxozaizdhxoimmraiakbak",
			dictionary: []string{"yv", "bmab", "hv", "bnsll", "mra", "jjqf", "g", "aiyzi", "ip", "pfctr", "flr", "ybbcl", "biu", "ke", "lpl", "iak", "pirua", "ilhqd", "zdhx", "fux", "xaw", "pdfvt", "xf", "t", "wq", "r", "cgmud", "aokas", "xv", "jf", "cyys", "wcaz", "rvegf", "ysg", "xo", "uwb", "lw", "okgk", "vbmi", "v", "mvo", "fxyx", "ad", "e"},
			expected:   9,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, minExtraChar(tc.s, tc.dictionary))
	}
}
