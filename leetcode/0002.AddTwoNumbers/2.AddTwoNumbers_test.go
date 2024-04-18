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

func TestAddTwoNumbers(t *testing.T) {
	//TODO: TC; isEqual
	l1 := createListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := createListNode([]int{9, 9, 9, 9})
	printListNode(l1)
	fmt.Println()
	printListNode(l2)
	fmt.Println()
	test := addTwoNumbers(l1, l2)
	printListNode(test)
}
