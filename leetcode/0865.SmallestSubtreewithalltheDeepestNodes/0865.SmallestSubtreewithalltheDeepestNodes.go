package main

import (
	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

func subtreeWithAllDeepest(root *util.TreeNode) *util.TreeNode {
	var nodeDfs func(*util.TreeNode, int) (int, int, *util.TreeNode)
	nodeDfs = func(node *util.TreeNode, level int) (int, int, *util.TreeNode) {
		if node == nil {
			return 0, 0, nil
		}

		leftNodes, leftLevel, leftDeepNode := nodeDfs(node.Left, level+1)
		rightNodes, rightLevel, rightDeepNode := nodeDfs(node.Right, level+1)

		if leftNodes > 0 || rightNodes > 0 {
			if rightLevel == leftLevel {
				return leftNodes + rightNodes, rightLevel, node
			} else if rightLevel > leftLevel {
				return rightNodes, rightLevel, rightDeepNode
			} else {
				return leftNodes, leftLevel, leftDeepNode
			}
		}

		return 1, level, node
	}

	_, _, node := nodeDfs(root, 0)

	return node
}
