package leetcode

import (
	"fmt"
	"testing"
)

func createListNode(s []int) *ListNode {
	var result = new(ListNode)
	if len(s) > 0 {
		result.Val = s[0]
		node := createListNode(s[1:])
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

func printListNode(listNode *ListNode) {
	fmt.Printf("%v ", listNode.Val)
	if listNode.Next != nil {
		printListNode(listNode.Next)
	}
}

func TestReverseList(t *testing.T) {
	node := createListNode([]int{1, 2})
	node = reverseList(node)
	printListNode(node)
}
