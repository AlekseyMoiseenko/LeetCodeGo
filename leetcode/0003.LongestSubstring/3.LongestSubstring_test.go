package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	v := lengthOfLongestSubstring("abcabcbb")
	require.Equal(t, v, 3)
	v = lengthOfLongestSubstring("bbbbb")
	require.Equal(t, v, 1)
	v = lengthOfLongestSubstring("pwwkew")
	require.Equal(t, v, 3)
}
