package leetcode

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

type TreeNode = util.TreeNode

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}

	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}

	if root.Val == target && root.Left == nil && root.Right == nil {
		root = nil
	}

	return root
}
