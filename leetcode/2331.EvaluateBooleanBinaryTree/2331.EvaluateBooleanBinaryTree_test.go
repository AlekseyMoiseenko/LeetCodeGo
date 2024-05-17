package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestEvaluateTree(t *testing.T) {
	testCases := []struct {
		node     *TreeNode
		expected bool
	}{
		{
			node:     util.IntsToTreeNode([]int{2, 1, 3, util.NULL, util.NULL, 0, 1}),
			expected: true,
		},
		{
			node:     util.IntsToTreeNode([]int{0}),
			expected: false,
		},
	}

	for _, tc := range testCases {
		res := evaluateTree(tc.node)

		require.Equal(t, res, tc.expected)
	}
}
