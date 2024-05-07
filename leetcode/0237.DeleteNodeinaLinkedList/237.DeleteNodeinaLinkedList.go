package leetcode

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

type ListNode = util.ListNode

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
