package main

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

// Constraints:
// The number of nodes in the linked list will be in the range [1, 10^4].
func Test0382(t *testing.T) {
	testCases := []struct {
		head *util.ListNode
	}{
		{
			head: util.CreateLinkedListFromArray([]int{1, 2, 3}).Head,
		},
	}

	for _, tc := range testCases {
		s := Constructor(tc.head)
		_ = s.GetRandom()
	}
}
