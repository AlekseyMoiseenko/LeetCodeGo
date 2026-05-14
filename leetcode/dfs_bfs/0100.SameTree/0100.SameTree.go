package main

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

func isSameTreeBfs(p *util.TreeNode, q *util.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	queue := [][2]*util.TreeNode{{p, q}}

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

		queue = append(queue, [2]*util.TreeNode{a.Left, b.Left})
		queue = append(queue, [2]*util.TreeNode{a.Right, b.Right})
	}

	return true
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func isSameTreeDfs(p *util.TreeNode, q *util.TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreeDfs(p.Left, q.Left) && isSameTreeDfs(p.Right, q.Right)
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func isSameTreeDfs1(p *util.TreeNode, q *util.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	leftIs := isSameTreeDfs1(p.Left, q.Left)
	if !leftIs {
		return false
	}

	rightIs := isSameTreeDfs1(p.Right, q.Right)
	if !rightIs {
		return false
	}

	return leftIs && rightIs
}
