package main

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

// Constraints:
// The number of nodes in the tree is in the range [1, 1000].
// -100 <= Node.val <= 100
func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		root     *util.TreeNode
		expected bool
	}{
		{
			root:     util.IntsToTreeNode([]int{1, 2, 2, 3, 4, 4, 3}),
			expected: true,
		},
		{
			root:     util.IntsToTreeNode([]int{1, 2, 2, util.NULL, 3, util.NULL, 3}),
			expected: false,
		},
		{
			root:     util.IntsToTreeNode([]int{1, 2, 2, 3, 5, 6, 3}),
			expected: false,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, isSymmetricBfs(tc.root))
		require.Equal(t, tc.expected, isSymmetricDfs(tc.root))
	}
}
