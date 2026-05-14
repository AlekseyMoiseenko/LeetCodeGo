package main

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

func isSymmetricBfs(root *util.TreeNode) bool {
	queue := [][2]*util.TreeNode{{root.Left, root.Right}}

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		a, b := pair[0], pair[1]

		if a == nil && b == nil {
			continue
		}

		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}

		queue = append(queue, [2]*util.TreeNode{a.Right, b.Left})
		queue = append(queue, [2]*util.TreeNode{a.Left, b.Right})
	}

	return true
}

func isSymmetricDfs(root *util.TreeNode) bool {
	return isSymmetricDfs1(root.Left, root.Right)
}

func isSymmetricDfs1(left, right *util.TreeNode) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricDfs1(left.Right, right.Left) && isSymmetricDfs1(left.Left, right.Right)
}
