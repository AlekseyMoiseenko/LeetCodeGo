package hash

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

// Time Complexity: O(n+m)
// Space Complexity: O(1)
func getIntersectionNode(headA, headB *util.ListNode) *util.ListNode {
	m := make(map[*util.ListNode]struct{})

	currentNode := headA
	for currentNode != nil {
		m[currentNode] = struct{}{}
		currentNode = currentNode.Next
	}

	currentNode = headB
	for currentNode != nil {
		if _, exist := m[currentNode]; exist {
			return currentNode
		}

		currentNode = currentNode.Next
	}

	return nil
}
