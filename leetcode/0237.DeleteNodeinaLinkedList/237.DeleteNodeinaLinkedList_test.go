package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

func TestDeleteNode(t *testing.T) {
	node := util.CreateListNode([]int{4, 5, 1, 9})
	deleteNode(node.Next)
	util.PrintListNode(node)
}
