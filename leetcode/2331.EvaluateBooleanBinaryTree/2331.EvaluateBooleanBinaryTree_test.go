package leetcode

import (
	"fmt"
	"testing"
)

// TODO: tests
func TestEvaluateTree(t *testing.T) {
	testTree := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 1},
		},
	}
	res := evaluateTree(testTree)
	fmt.Printf("res: %v\n", res)
}
