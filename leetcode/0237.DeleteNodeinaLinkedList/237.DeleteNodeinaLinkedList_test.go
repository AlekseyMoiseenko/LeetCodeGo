package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestDeleteNode(t *testing.T) {
	testCases := []struct {
		n        int
		list     []int
		expected []int
	}{
		{
			n:        5,
			list:     []int{4, 5, 1, 9},
			expected: []int{4, 1, 9},
		},
		{
			n:        1,
			list:     []int{4, 5, 1, 9},
			expected: []int{4, 5, 9},
		},
	}

	for _, tc := range testCases {
		head := util.CreateListNode([]int{4, 5, 1, 9})
		node := util.FindNode(head, tc.n)

		deleteNode(node)

		require.Equal(t, util.ListToInts(head), tc.expected)
	}
}
