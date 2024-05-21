package leetcode

func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for i := 0; i < len(nums); i++ {
		l := len(res)

		for j := 0; j < l; j++ {
			item := res[j]

			tmp := []int{}
			tmp = append(tmp, item...)
			tmp = append(tmp, nums[i])

			res = append(res, tmp)
		}
	}

	return res
}
