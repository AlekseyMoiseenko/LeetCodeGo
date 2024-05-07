package util

import "fmt"

//Default leetcode ListNode
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

func FindNode(head *ListNode, n int) *ListNode {
	for head != nil {
		if head.Val == n {
			return head
		}

		head = head.Next
	}

	return nil
}

func ListToInts(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func IsEqualIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
