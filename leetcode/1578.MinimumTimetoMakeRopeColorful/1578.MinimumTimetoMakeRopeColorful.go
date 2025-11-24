package leetcode

func minCost(colors string, neededTime []int) int {
	if len(colors) <= 1 {
		return 0
	}

	total := 0
	for i := 0; i < len(colors)-1; i++ {
		if colors[i] == colors[i+1] {
			if neededTime[i] < neededTime[i+1] {
				total += neededTime[i]
			} else {
				total += neededTime[i+1]
				neededTime[i+1] = neededTime[i]
			}
		}
	}

	return total
}
