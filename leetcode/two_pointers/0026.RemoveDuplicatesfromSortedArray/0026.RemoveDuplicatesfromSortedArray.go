package main

func removeDuplicates(nums []int) int {
	nextUniqueIndex := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nextUniqueIndex] = nums[i]
			nextUniqueIndex++
		}
	}

	return nextUniqueIndex
}
