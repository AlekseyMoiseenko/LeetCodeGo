package leetcode

func largestLocal(grid [][]int) [][]int {
	n := len(grid) - 2
	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			local := []int{}
			max := 0

			for k := 0; k < 3; k++ {
				local = append(local, grid[i+k][j:j+3]...)
			}

			for k := 0; k < len(local); k++ {
				if local[k] > max {
					max = local[k]
				}
			}

			res[i][j] = max
		}
	}

	return res
}

// First app
// func largestLocal(grid [][]int) [][]int {
// 	n := len(grid) - 2
// 	res := [][]int{}

// 	for i := 0; i < n; i++ {
// 		res = append(res, []int{})
// 		for j := 0; j < n; j++ {
// 			local := []int{}

// 			for k := 0; k < 3; k++ {
// 				local = append(local, grid[i+k][j:j+3]...)
// 			}
// 			sort.Ints(local)

// 			res[i] = append(res[i], local[len(local)-1])
// 		}
// 	}

// 	return res
// }
