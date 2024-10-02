package leetcode

import "sort"

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	copiedArr := append([]int{}, arr...)
	sort.Ints(copiedArr)

	rankMap := make(map[int]int)
	rank := 1

	for _, num := range copiedArr {
		if _, exist := rankMap[num]; !exist {
			rankMap[num] = rank
			rank++
		}
	}

	res := make([]int, len(arr))
	for i, num := range arr {
		res[i] = rankMap[num]
	}

	return res
}
