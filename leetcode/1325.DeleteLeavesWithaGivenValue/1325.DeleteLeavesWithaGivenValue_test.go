package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestRemoveLeafNodes(t *testing.T) {
	testCases := []struct {
		nodeList []int
		target   int
		expected []int
	}{
		{
			nodeList: []int{1, 2, util.NULL, 2, util.NULL, 2},
			target:   2,
			expected: []int{1},
		},
		{
			nodeList: []int{1, 2, 3, 2, util.NULL, 2, 4},
			target:   2,
			expected: []int{1, util.NULL, 3, util.NULL, 4},
		},
		{
			nodeList: []int{1, 3, 3, 3, 2},
			target:   3,
			expected: []int{1, 3, util.NULL, util.NULL, 2},
		},
	}

	for _, tc := range testCases {
		node := util.IntsToTreeNode(tc.nodeList)
		res := removeLeafNodes(node, tc.target)

		require.Equal(t, tc.expected, util.TreeToInts(res))
	}
}
