package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(s []int) *ListNode {
	var result = new(ListNode)
	if len(s) > 0 {
		result.Val = s[0]
		node := CreateListNode(s[1:])
		if node != nil {
			result.Next = node
		} else {
			result.Next = nil
		}
	} else {
		return nil
	}

	return result
}

func PrintListNode(listNode *ListNode) {
	fmt.Printf("%v ", listNode.Val)
	if listNode.Next != nil {
		PrintListNode(listNode.Next)
	}
}
