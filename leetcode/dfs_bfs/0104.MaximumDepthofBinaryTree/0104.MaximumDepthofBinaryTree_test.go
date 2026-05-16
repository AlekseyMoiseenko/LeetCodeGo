package main

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		root     *util.TreeNode
		expected int
	}{
		{
			root:     util.IntsToTreeNode([]int{3, 9, 20, util.NULL, util.NULL, 15, 7}),
			expected: 3,
		},
		{
			root:     util.IntsToTreeNode([]int{1, util.NULL, 2}),
			expected: 2,
		},
		{
			root:     util.IntsToTreeNode([]int{3, 9, 20, 15, 6}),
			expected: 3,
		},
		{
			root:     nil,
			expected: 0,
		},
		{
			root:     util.IntsToTreeNode([]int{1}),
			expected: 1,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, maxDepthDfs(tc.root))
		require.Equal(t, tc.expected, maxDepthBfs(tc.root))
	}
}
