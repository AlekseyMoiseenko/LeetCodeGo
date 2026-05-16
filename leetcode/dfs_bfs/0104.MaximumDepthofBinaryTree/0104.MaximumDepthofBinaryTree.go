package main

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

// Time complexity: O(n)
// Space complexity: O(n)
func maxDepthDfs(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepthDfs(root.Left)
	right := maxDepthDfs(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// Time complexity: O(n)
// Space complexity: O(n)
func maxDepthBfs(root *util.TreeNode) int {
	if root == nil {
		return 0
	}

	type nodeWithDepth struct {
		node  *util.TreeNode
		level int
	}

	maxLevel := 1
	queue := []*nodeWithDepth{{node: root, level: 1}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node.Left != nil {
			queue = append(queue, &nodeWithDepth{node: current.node.Left, level: current.level + 1})
		}
		if current.node.Right != nil {
			queue = append(queue, &nodeWithDepth{node: current.node.Right, level: current.level + 1})
		}

		if current.level > maxLevel {
			maxLevel = current.level
		}
	}

	return maxLevel
}
