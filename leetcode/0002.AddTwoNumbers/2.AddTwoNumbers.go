package leetcode

import (
	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

type ListNode = util.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = new(ListNode)
	temp := result
	var increase int
	for l1 != nil || l2 != nil || increase != 0 {
		temp.Next = new(ListNode)
		temp = temp.Next
		if l1 != nil && l2 != nil {
			temp.Val = (l1.Val + l2.Val + increase) % 10
			increase = (l1.Val + l2.Val + increase) / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			temp.Val = (l1.Val + increase) % 10
			increase = (l1.Val + increase) / 10
			l1 = l1.Next
			l2 = nil
		} else if l2 != nil {
			temp.Val = (l2.Val + increase) % 10
			increase = (l2.Val + increase) / 10
			l1 = nil
			l2 = l2.Next
		} else if increase > 0 {
			temp.Val = increase
			break
		} else {
			break
		}
	}
	return result.Next
}

// first approach
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	return addTwoNumbersN(0, l1, l2)
// }

// func addTwoNumbersN(increase int, l1 *ListNode, l2 *ListNode) *ListNode {
// 	var result *ListNode
// 	if l1 != nil && l2 != nil {
// 		result = calculate(l1.Val, l2.Val, increase, l1.Next, l2.Next)
// 	} else {
// 		if l1 != nil {
// 			result = calculate(l1.Val, 0, increase, l1.Next, nil)
// 		} else if l2 != nil {
// 			result = calculate(0, l2.Val, increase, nil, l2.Next)
// 		} else if increase > 0 {
// 			result = new(ListNode)
// 			result.Val = increase
// 		} else {
// 			return nil
// 		}
// 	}

// 	return result
// }

// func calculate(v1, v2, increase int, next1, next2 *ListNode) *ListNode {
// 	var result = new(ListNode)

// 	result.Val = (v1 + v2 + increase) % 10
// 	increase = (v1 + v2 + increase) / 10
// 	node := addTwoNumbersN(increase, next1, next2)
// 	if node != nil {
// 		result.Next = node
// 	}

// 	return result
// }
