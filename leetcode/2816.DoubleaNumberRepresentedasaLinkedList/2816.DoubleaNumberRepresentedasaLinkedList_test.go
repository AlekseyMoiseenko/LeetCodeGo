package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestDoubleIt(t *testing.T) {
	testCases := []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{1, 8, 9},
			expected: []int{3, 7, 8},
		},
		{
			list:     []int{9, 9, 9},
			expected: []int{1, 9, 9, 8},
		},
	}

	for _, tc := range testCases {
		node := util.CreateListNode(tc.list)
		node = doubleIt(node)

		require.Equal(t, util.ListToInts(node), tc.expected)
	}
}
