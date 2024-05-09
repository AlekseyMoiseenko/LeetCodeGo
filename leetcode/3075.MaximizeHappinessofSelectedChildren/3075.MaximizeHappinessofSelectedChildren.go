package leetcode

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	index := 0
	resultList := []int{}

	for i := len(happiness) - 1; i >= 0; i-- {
		hap := happiness[i] - index
		if hap > 0 {
			resultList = append(resultList, hap)
		}

		index++
		if index >= k {
			break
		}
	}

	var res int64
	for i := range resultList {
		res += int64(resultList[i])
	}

	return res
}
