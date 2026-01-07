package main

import (
	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

const MOD = 1000000007 //const MOD int = 1e9 + 7

// Hint: If we know the sum of a subtree, the answer is max((total_sum - subtree_sum) * subtree_sum)) in each node.
func maxProduct(root *util.TreeNode) int {
	maxProd := 0
	totalSum := getTotalSum(root)

	var nodeDfs func(*util.TreeNode) int
	nodeDfs = func(node *util.TreeNode) int {
		if node == nil {
			return 0
		}

		left := nodeDfs(node.Left)
		right := nodeDfs(node.Right)

		subtreeSum := node.Val + left + right
		product := (totalSum - subtreeSum) * subtreeSum
		if product > maxProd {
			maxProd = product
		}

		return subtreeSum
	}

	nodeDfs(root)

	return maxProd % MOD
}

func getTotalSum(node *util.TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Val + getTotalSum(node.Left) + getTotalSum(node.Right)
}
