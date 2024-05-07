package leetcode

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

type ListNode = util.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}
