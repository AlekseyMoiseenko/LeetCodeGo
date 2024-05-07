package leetcode

import (
	"fmt"
	"testing"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

func TestAddTwoNumbers(t *testing.T) {
	//TODO: TC; isEqual
	l1 := util.CreateListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := util.CreateListNode([]int{9, 9, 9, 9})
	util.PrintListNode(l1)
	fmt.Println()
	util.PrintListNode(l2)
	fmt.Println()
	test := addTwoNumbers(l1, l2)
	util.PrintListNode(test)
}
