package main

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestMaxProduct(t *testing.T) {
	// The number of nodes in the tree is in the range [2, 5 * 104].
	// 1 <= Node.val <= 104
	testCases := []struct {
		Node     *util.TreeNode
		Expected int
	}{
		{
			Node:     util.IntsToTreeNode([]int{1, 2, 3, 4, 5, 6}),
			Expected: 110,
		},
		{
			Node:     util.IntsToTreeNode([]int{1, util.NULL, 2, 3, 4, util.NULL, util.NULL, 5, 6}),
			Expected: 90,
		},
		{
			Node:     util.IntsToTreeNode([]int{10, 10}),
			Expected: 100,
		},
		{
			Node:     util.IntsToTreeNode([]int{3, 7, 1}),
			Expected: 28,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.Expected, maxProduct(tc.Node))
	}
}
