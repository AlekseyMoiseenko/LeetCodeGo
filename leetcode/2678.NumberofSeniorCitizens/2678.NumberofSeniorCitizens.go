package leetcode

import (
	"strconv"
)

const (
	ageStart = 11
	ageEnd   = 13
)

func countSeniors(details []string) int {
	res := 0

	for _, s := range details {
		if v, ok := strconv.Atoi(s[ageStart:ageEnd]); ok == nil {
			if v > 60 {
				res += 1
			}
		}
	}

	return res
}
