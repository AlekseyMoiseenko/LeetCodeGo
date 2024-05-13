package leetcode

func matrixScore(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	res := (1 << (m - 1)) * n

	for i := 1; i < m; i++ {
		curr := 0
		for j := 0; j < n; j++ {
			if grid[j][i] == grid[j][0] {
				curr++
			}
		}
		if curr > n-curr {
			res += (1 << (m - i - 1)) * curr
		} else {
			res += (1 << (m - i - 1)) * (n - curr)
		}
	}
	return res
}
