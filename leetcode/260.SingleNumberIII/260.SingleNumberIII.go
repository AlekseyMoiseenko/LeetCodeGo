package leetcode

func singleNumber(nums []int) []int {
	numMap := make(map[int]int, len(nums))

	for _, v := range nums {
		numMap[v]++
	}

	var result []int
	for key := range numMap {
		if numMap[key] == 1 {
			result = append(result, key)
		}
	}

	return result
}
