package leetcode

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

func swapPairs(head *util.ListNode) *util.ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	first := head.Next
	first.Next, head.Next = head, swapPairs(head.Next.Next)

	return first
}
