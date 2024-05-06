package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	head = reverse(head)

	currMax := 0

	var prev, curr *ListNode = nil, head
	for curr != nil {
		if curr.Val < currMax {
			if prev != nil {
				prev.Next = curr.Next
			}
		} else {
			prev = curr
			currMax = curr.Val
		}
		curr = curr.Next
	}

	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

// First submission; misunderstood the task
// func removeNodes(head *ListNode) *ListNode {
// 	var result = new(ListNode)
// 	temp := result
// 	mas := []int{}

// 	for head != nil {
// 		mas = append(mas, head.Val)

// 		if head.Next == nil {
// 			for i := 0; i < len(mas); i++ {
// 				temp.Next = new(ListNode)
// 				temp = temp.Next
// 				temp.Val = mas[i]
// 			}
// 			break
// 		}

// 		if head.Next.Val > mas[0] {
// 			temp.Next = new(ListNode)
// 			temp = temp.Next
// 			temp.Val = head.Next.Val
// 			head = head.Next.Next
// 			mas = nil
// 		} else if head.Next.Val <= mas[0] {
// 			head = head.Next
// 		}
// 	}

// 	return result.Next
// }
