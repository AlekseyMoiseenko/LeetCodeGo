package leetcode

func sortColors(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i+1], nums[i] = nums[i], nums[i+1]
			i = -1
		}
	}
}
