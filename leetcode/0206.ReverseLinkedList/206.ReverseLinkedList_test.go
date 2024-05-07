package leetcode

import (
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

func TestReverseList(t *testing.T) {
	node := util.CreateListNode([]int{1, 2})
	node = reverseList(node)
	util.PrintListNode(node)
}
