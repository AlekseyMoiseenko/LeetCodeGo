package main

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

// Constraints:
// The number of nodes in both trees is in the range [0, 100].
// -10^4 <= Node.val <= 10^4
func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		p        *util.TreeNode
		q        *util.TreeNode
		expected bool
	}{
		{
			p:        util.IntsToTreeNode([]int{1, 2, 3}),
			q:        util.IntsToTreeNode([]int{1, 2, 3}),
			expected: true,
		},
		{
			p:        util.IntsToTreeNode([]int{1, 2}),
			q:        util.IntsToTreeNode([]int{1, util.NULL, 2}),
			expected: false,
		},
		{
			p:        util.IntsToTreeNode([]int{1, 2, 1}),
			q:        util.IntsToTreeNode([]int{1, 1, 2}),
			expected: false,
		},
		{
			p:        util.IntsToTreeNode([]int{1, 2, util.NULL, 3}),
			q:        util.IntsToTreeNode([]int{1, util.NULL, 2, util.NULL, 3, util.NULL, util.NULL}),
			expected: false,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, isSameTreeBfs(tc.p, tc.q))
		require.Equal(t, tc.expected, isSameTreeDfs(tc.p, tc.q))
		require.Equal(t, tc.expected, isSameTreeDfs1(tc.p, tc.q))
	}
}
