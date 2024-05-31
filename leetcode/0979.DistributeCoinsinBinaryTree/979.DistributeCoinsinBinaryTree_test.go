package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestDistributeCoins(t *testing.T) {
	testCases := []struct {
		intsTree []int
		expected int
	}{
		{
			intsTree: []int{0, 3, 0},
			expected: 3,
		},
		{
			intsTree: []int{3, 0, 0},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		node := util.IntsToTreeNode(tc.intsTree)
		res := distributeCoins(node)

		require.Equal(t, tc.expected, res)
	}
}
