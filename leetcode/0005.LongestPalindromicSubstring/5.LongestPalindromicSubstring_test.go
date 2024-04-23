package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindrome(t *testing.T) {
	result := longestPalindrome("babaddaba")
	require.Equal(t, result, "abaddaba")

	result = longestPalindrome("babad")
	require.Equal(t, result, "bab")

	result = longestPalindrome("cbbd")
	require.Equal(t, result, "bb")
}
