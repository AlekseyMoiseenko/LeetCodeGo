package leetcode

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

type ListNode = util.ListNode

func doubleIt(head *ListNode) *ListNode {
	head = &ListNode{Val: 0, Next: head}
	temp := head

	for head.Next != nil {
		head.Val += (head.Next.Val * 2) / 10
		head.Next.Val = (head.Next.Val * 2) % 10
		head = head.Next
	}

	if temp.Val > 0 {
		return temp
	}
	return temp.Next
}

// Faster way from Leetcode
// func doubleIt(head *ListNode) *ListNode {

//     head = &ListNode{Next: head}

//     for curr, next := head, head.Next; next != nil; curr, next = next, next.Next {
//         next.Val *= 2
//         curr.Val += next.Val / 10
//         next.Val %= 10
//     }

//     if head.Val == 0 {
//         head = head.Next
//     }

//     return head
// }
