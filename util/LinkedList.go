package util

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func CreateLinkedListFromArray(s []int) *LinkedList {
	newList := &LinkedList{}
	newList.Head = newList.createNodesFromArray(s)
	return newList
}

func (list *LinkedList) createNodesFromArray(s []int) *ListNode {
	if len(s) <= 0 {
		return nil
	}

	result := &ListNode{Val: s[0]}
	node := list.createNodesFromArray(s[1:])
	if node != nil {
		result.Next = node
	} else {
		list.Tail = result
		result.Next = nil
	}

	return result
}

func (list *LinkedList) GetNodeByIndex(idx int) *ListNode {
	if idx < 0 {
		return nil
	}

	currentNode := list.Head
	for idx > 0 {
		if currentNode.Next == nil {
			return nil
		}
		currentNode = currentNode.Next
		idx--
	}

	return currentNode
}
