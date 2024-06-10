package leetcode

import "sort"

func heightChecker(heights []int) int {
	copy := []int{}
	copy = append(copy, heights...)
	sort.Ints(copy)

	res := 0

	for i := 0; i < len(heights); i++ {
		if heights[i] != copy[i] {
			res++
		}
	}

	return res
}
