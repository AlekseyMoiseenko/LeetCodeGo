package hash

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

// Constraints:
// The number of nodes of listA is in the m.
// The number of nodes of listB is in the n.
// 1 <= m, n <= 3 * 10^4
// 1 <= Node.val <= 10^5
// 0 <= skipA <= m
// 0 <= skipB <= n
// intersectVal is 0 if listA and listB do not intersect.
// intersectVal == listA[skipA] == listB[skipB] if listA and listB intersect.
func TestGetIntersectionNode(t *testing.T) {
	testCases := []struct {
		mergeList *util.LinkedList
		listA     *util.LinkedList
		listB     *util.LinkedList
	}{
		{
			mergeList: util.CreateLinkedListFromArray([]int{8, 4, 5}),
			listA:     util.CreateLinkedListFromArray([]int{4, 1}),
			listB:     util.CreateLinkedListFromArray([]int{5, 6, 1}),
		},
		{
			mergeList: util.CreateLinkedListFromArray([]int{2, 4}),
			listA:     util.CreateLinkedListFromArray([]int{1, 9, 1}),
			listB:     util.CreateLinkedListFromArray([]int{3}),
		},
		{
			mergeList: nil,
			listA:     util.CreateLinkedListFromArray([]int{2, 6, 4}),
			listB:     util.CreateLinkedListFromArray([]int{1, 5}),
		},
		{
			mergeList: nil,
			listA:     util.CreateLinkedListFromArray([]int{1}),
			listB:     util.CreateLinkedListFromArray([]int{100000}),
		},
		{
			mergeList: util.CreateLinkedListFromArray([]int{2}),
			listA:     util.CreateLinkedListFromArray([]int{1}),
			listB:     util.CreateLinkedListFromArray([]int{100000}),
		},
	}

	for _, tc := range testCases {
		var expected *util.ListNode
		if tc.mergeList != nil {
			tc.listA.Tail.Next = tc.mergeList.Head
			tc.listB.Tail.Next = tc.mergeList.Head
			expected = tc.mergeList.Head
		}

		require.Equal(t, expected, getIntersectionNode(tc.listA.Head, tc.listB.Head))
	}
}
