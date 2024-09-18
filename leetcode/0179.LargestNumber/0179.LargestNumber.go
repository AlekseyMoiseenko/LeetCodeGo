package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := range strs {
		strs[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(strs, func(i, j int) bool {
		return (strs[i] + strs[j]) > (strs[j] + strs[i])
	})

	result := strings.Join(strs, "")
	if result[0] == '0' {
		result = "0"
	}

	return result
}

// func strcmp(s1, s2 string) int {
// 	lens := len(s1)
// 	if lens > len(s2) {
// 		lens = len(s2)
// 	}
// 	for i := 0; i < lens; i++ {
// 		if s1[i] != s2[i] {
// 			return int(s1[i]) - int(s2[i])
// 		}
// 	}
// 	return len(s1) - len(s2) // 0
// }
