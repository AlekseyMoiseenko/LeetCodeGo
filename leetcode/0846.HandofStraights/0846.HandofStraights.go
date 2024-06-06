package leetcode

import (
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	l := len(hand)
	if l%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	for i := 0; i < l/groupSize; i++ {
		tmp := []int{}

		for j := 0; j < l; j++ {
			if hand[j] == -1 {
				continue
			}

			if len(tmp) > 0 {
				if hand[j] == tmp[len(tmp)-1] {
					continue
				}
				if hand[j]-1 != tmp[len(tmp)-1] {
					return false
				}
			}
			tmp = append(tmp, hand[j])
			hand[j] = -1

			if len(tmp) == groupSize {
				break
			}
		}

		if len(tmp) < groupSize {
			return false
		}
	}

	return true
}
