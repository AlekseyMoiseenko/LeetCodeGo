package leetcode

import "github.com/AlekseyMoiseenko/LeetCodeGo/util"

type TreeNode = util.TreeNode

func distributeCoins(root *TreeNode) int {
	moves := 0
	distribute(root, &moves)
	return moves
}

func distribute(node *TreeNode, moves *int) int {
	if node == nil {
		return 0
	}
	left := distribute(node.Left, moves)
	right := distribute(node.Right, moves)
	*moves += abs(left) + abs(right)
	return (left + right + node.Val - 1)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
