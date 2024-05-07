package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{
			list1:    []int{9, 9, 9, 9, 9, 9, 9},
			list2:    []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			list1:    []int{2, 4, 3},
			list2:    []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			list1:    []int{0},
			list2:    []int{0},
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		l1 := util.CreateListNode(tc.list1)
		l2 := util.CreateListNode(tc.list2)

		res := addTwoNumbers(l1, l2)

		require.Equal(t, true, util.IsEqualIntSlices(util.ListToInts(res), tc.expected))
	}
}
