package hash

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
	"github.com/stretchr/testify/require"
)

// Constraints:
// The number of the nodes in the list is in the range [0, 10^4].
// -10^5 <= Node.val <= 10^5
// pos is -1 or a valid index in the linked-list.
func TestHasCycle(t *testing.T) {
	testCases := []struct {
		list     *util.LinkedList
		pos      int
		expected bool
	}{
		{
			list:     util.CreateLinkedListFromArray([]int{3, 2, 0, -4}),
			pos:      1,
			expected: true,
		},
		{
			list:     util.CreateLinkedListFromArray([]int{1, 2}),
			pos:      0,
			expected: true,
		},
		{
			list:     util.CreateLinkedListFromArray([]int{1}),
			pos:      -1,
			expected: false,
		},
		{
			list:     util.CreateLinkedListFromArray([]int{3, 2, 0, -4, 3, 2, 0, -4, 3, 2, 0, -4}),
			pos:      -1,
			expected: false,
		},
		{
			list:     util.CreateLinkedListFromArray([]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}),
			pos:      -1,
			expected: false,
		},
		{
			list:     util.CreateLinkedListFromArray([]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}),
			pos:      6,
			expected: true,
		},
		{
			list:     util.CreateLinkedListFromArray([]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}),
			pos:      0,
			expected: true,
		},
		{
			list:     util.CreateLinkedListFromArray([]int{}),
			pos:      -1,
			expected: false,
		},
	}

	for _, tc := range testCases {
		if tc.pos >= 0 {
			tc.list.Tail.Next = tc.list.GetNodeByIndex(tc.pos)
		}
		require.Equal(t, tc.expected, hasCycle(tc.list.Head))
	}
}
