package leetcode

import (
	"sort"
	"strconv"
)

// All the scores are guaranteed to be unique.
func findRelativeRanks(score []int) []string {
	mapNum := make(map[int]int)
	var res []string
	var sorted []int

	sorted = append(sorted, score...)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	for i := range sorted {
		mapNum[sorted[i]] = i + 1
	}

	for i := range score {
		if mapNum[score[i]] == 1 {
			res = append(res, "Gold Medal")
		} else if mapNum[score[i]] == 2 {
			res = append(res, "Silver Medal")
		} else if mapNum[score[i]] == 3 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(mapNum[score[i]]))
		}
	}

	return res
}

// first app
// func findRelativeRanks(score []int) []string {
// 	res := []string{}

// 	sorted := []int{}
// 	sorted = append(sorted, score...)
// 	sort.Slice(sorted, func(i, j int) bool {
// 		return sorted[i] > sorted[j]
// 	})

// 	for i := 0; i < len(score); i++ {
// 		for j := 0; j < len(sorted); j++ {
// 			if score[i] == sorted[j] {
// 				if j >= 3 {
// 					res = append(res, fmt.Sprint(j+1))
// 				} else if j == 2 {
// 					res = append(res, "Bronze Medal")
// 				} else if j == 1 {
// 					res = append(res, "Silver Medal")
// 				} else {
// 					res = append(res, "Gold Medal")
// 				}

// 				break
// 			}
// 		}
// 	}

// 	return res
// }
