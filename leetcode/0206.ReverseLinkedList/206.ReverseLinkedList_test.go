package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestReverseList(t *testing.T) {
	testCases := []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			list:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			list:     []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		node := util.CreateListNode(tc.list)
		node = reverseList(node)

		require.Equal(t, true, util.IsEqualIntSlices(util.ListToInts(node), tc.expected))
	}
}
