package hash

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

// Time Complexity: O(n)
// Space Complexity: O(n)
func hasCycle(head *util.ListNode) bool {
	m := make(map[*util.ListNode]struct{})

	currentNode := head
	for currentNode != nil {
		if _, exist := m[currentNode]; exist {
			return true
		}
		m[currentNode] = struct{}{}
		currentNode = currentNode.Next
	}

	return false
}
