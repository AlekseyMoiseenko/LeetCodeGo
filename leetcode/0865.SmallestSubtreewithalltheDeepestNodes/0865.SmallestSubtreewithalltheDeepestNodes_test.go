package main

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestMaxProduct(t *testing.T) {
	// The number of nodes in the tree will be in the range [1, 500].
	// 0 <= Node.val <= 500
	// The values of the nodes in the tree are unique.

	node1 := util.IntsToTreeNode([]int{3, 5, 1, 6, 2, 0, 8, util.NULL, util.NULL, 7, 4})
	node2 := util.IntsToTreeNode([]int{1})
	node3 := util.IntsToTreeNode([]int{0, 1, 3, util.NULL, 111})
	node4 := util.IntsToTreeNode([]int{3, 5, 1, 6, 2, 0, 8, util.NULL, util.NULL, 7, 4, util.NULL, util.NULL, 11, 200})
	node5 := util.IntsToTreeNode([]int{0, 3, 1, 4, util.NULL, 2, util.NULL, util.NULL, 6, util.NULL, 5})

	testCases := []struct {
		Node     *util.TreeNode
		Expected *util.TreeNode
	}{
		{
			Node:     node1,
			Expected: node1.Left.Right,
		},
		{
			Node:     node2,
			Expected: node2,
		},
		{
			Node:     node3,
			Expected: node3.Left.Right,
		},
		{
			Node:     node4,
			Expected: node4,
		},
		{
			Node:     node5,
			Expected: node5,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.Expected, subtreeWithAllDeepest(tc.Node))
	}
}
