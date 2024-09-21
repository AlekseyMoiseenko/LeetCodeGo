package leetcode

func lexicalOrder(n int) []int {
	res := []int{}
	current := 1

	for len(res) < n {
		res = append(res, current)

		if current*10 <= n {
			current *= 10
		} else {
			if current >= n {
				current /= 10
			}

			current++

			for current%10 == 0 {
				current /= 10
			}
		}
	}

	return res
}

// First app
// func lexicalOrder(n int) []int {
// 	res := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		res[i] = i + 1
// 	}

// 	sort.Slice(res, func(i, j int) bool {
// 		if strings.Compare(strconv.Itoa(res[i]), strconv.Itoa(res[j])) > 0 {
// 			return false
// 		} else {
// 			return true
// 		}
// 	})

// 	return res
// }
