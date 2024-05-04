package leetcode

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	boatsNum := 0

	startIndex := 0
	endIndex := len(people) - 1

	for startIndex < endIndex {
		if people[startIndex]+people[endIndex] <= limit {
			boatsNum++
			startIndex++
		} else {
			boatsNum++
		}
		endIndex--
	}

	if startIndex == endIndex {
		boatsNum++
	}

	return boatsNum
}
