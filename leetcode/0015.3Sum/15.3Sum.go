package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}

	l := len(nums)
	if l < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left, right := i+1, l-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// time limit exceed
// func threeSum(nums []int) [][]int {
// 	var result [][]int
// 	numMap := make(map[[3]int]bool)

// 	sort.Ints(nums)

// 	for len(nums) >= 3 {
// 		for j := 1; j < len(nums); j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[0]+nums[j]+nums[k] == 0 {
// 					numMap[[3]int{nums[0], nums[j], nums[k]}] = true
// 				}
// 			}
// 		}

// 		nums = nums[1:]
// 	}

// 	for key := range numMap {
// 		result = append(result, key[:])
// 	}

// 	return result
// }
