package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		listNodeOrig *util.ListNode
		expected     []int
	}{
		{
			listNodeOrig: util.CreateListNode([]int{1, 2, 3, 4}),
			expected:     []int{2, 1, 4, 3},
		},
		{
			listNodeOrig: util.CreateListNode([]int{}),
			expected:     []int{},
		},
		{
			listNodeOrig: util.CreateListNode([]int{1}),
			expected:     []int{1},
		},
		{
			listNodeOrig: util.CreateListNode([]int{1, 2, 3}),
			expected:     []int{2, 1, 3},
		},
		{
			listNodeOrig: util.CreateListNode([]int{99, 98, 97, 96, 95, 96, 97, 98, 99}),
			expected:     []int{98, 99, 96, 97, 96, 95, 98, 97, 99},
		},
	}

	for _, tc := range testCases {
		result := swapPairs(tc.listNodeOrig)

		require.Equal(t, tc.expected, util.ListToInts(result))
	}
}
