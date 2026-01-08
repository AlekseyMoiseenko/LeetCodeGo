package main

import (
	"math"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

func maxLevelSum(root *util.TreeNode) int {
	maxSum := math.MinInt64
	result, currentLevel := 1, 1

	queueSlice := []*util.TreeNode{root}

	for len(queueSlice) > 0 {
		size := len(queueSlice)
		sumAtCurrentLevel := 0

		for i := 0; i < size; i++ {
			node := queueSlice[0]
			queueSlice = queueSlice[1:]
			sumAtCurrentLevel += node.Val
			if node.Left != nil {
				queueSlice = append(queueSlice, node.Left)
			}
			if node.Right != nil {
				queueSlice = append(queueSlice, node.Right)
			}
		}

		if sumAtCurrentLevel > maxSum {
			maxSum = sumAtCurrentLevel
			result = currentLevel
		}
		currentLevel++
	}

	return result
}
