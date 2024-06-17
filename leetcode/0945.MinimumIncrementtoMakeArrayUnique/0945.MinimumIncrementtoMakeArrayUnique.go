package leetcode

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	moves := 0

	for i := 1; i < len(nums); i++ {
		prev := nums[i-1] + 1
		if nums[i] < prev {
			moves += prev - nums[i]
			nums[i] = prev
		}
	}

	return moves
}

// sec app time limit
// func minIncrementForUnique(nums []int) int {
// 	numMap := make(map[int]int, len(nums))

// 	for _, key := range nums {
// 		numMap[key]++
// 	}

// 	var moves int

// 	for i := 0; i < 1; i++ {
// 		for k := range numMap {
// 			if numMap[k] > 1 {
// 				for inc := 1; numMap[k] > 1; inc++ {
// 					numMap[k+inc]++
// 					numMap[k]--
// 					moves += inc
// 				}

// 				i = -1
// 			}
// 		}
// 	}

// 	return moves
// }

// first app time limit
// func minIncrementForUnique(nums []int) int {
// 	moves := 0
// 	sort.Ints(nums)

// 	for i := 0; i < len(nums); i++ {
// 		isUnique := true
// 		for j := 0; j < i; j++ {
// 			if nums[i] == nums[j] {
// 				isUnique = false
// 				break
// 			}
// 		}

// 		if !isUnique {
// 			for !isUnique {
// 				nums[i]++
// 				moves++
// 				for j := 0; j < i; j++ {
// 					if nums[i] == nums[j] {
// 						isUnique = false
// 						break
// 					}

// 					if j+1 >= i {
// 						isUnique = true
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return moves
// }
