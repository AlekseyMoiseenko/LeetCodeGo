package main

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

// The number of nodes in the tree is in the range [1, 104].
// -105 <= Node.val <= 105
func TestMaxLevelSum(t *testing.T) {
	testCases := []struct {
		Node     *util.TreeNode
		Expected int
	}{
		{
			Node:     util.IntsToTreeNode([]int{1, 7, 0, 7, -8, util.NULL, util.NULL}),
			Expected: 2,
		},
		{
			Node:     util.IntsToTreeNode([]int{989, util.NULL, 10250, 98693, -89388, util.NULL, util.NULL, util.NULL, -32127}),
			Expected: 2,
		},
		{
			Node:     util.IntsToTreeNode([]int{1}),
			Expected: 1,
		},
		{
			Node:     util.IntsToTreeNode([]int{10, 10}),
			Expected: 1, // Return the smallest level x
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.Expected, maxLevelSum(tc.Node))
	}
}
