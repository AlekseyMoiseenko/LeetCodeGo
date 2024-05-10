package leetcode

import (
	"sort"
)

type fractionRes struct {
	val1   int
	val2   int
	result float64
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	fractionList := []fractionRes{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			fractionList = append(fractionList, fractionRes{arr[i], arr[j], float64(float64(arr[i]) / float64(arr[j]))})
		}
	}

	sort.Slice(fractionList, func(i, j int) bool {
		return fractionList[i].result < fractionList[j].result
	})

	return []int{fractionList[k-1].val1, fractionList[k-1].val2}
}
