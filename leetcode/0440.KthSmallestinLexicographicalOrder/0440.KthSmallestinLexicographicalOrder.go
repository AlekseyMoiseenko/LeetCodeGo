package leetcode

func findKthNumber(n int, k int) int {
	k--
	current := 1

	for k > 0 {
		step := countSteps(n, current, current+1)
		if step <= k {
			current++
			k -= step
		} else {
			current *= 10
			k--
		}
	}

	return current
}

func countSteps(n, prefix1, prefix2 int) int {
	steps := 0

	for prefix1 <= n {
		steps += min(n+1, prefix2) - prefix1
		prefix1 *= 10
		prefix2 *= 10
	}

	return steps
}

// Second attemp. Time Limit Exceeded. 43/69 testcases passed
// func findKthNumber(n int, k int) int {
// 	var res int
// 	k -= 1
// 	current := 1

// 	for i := 0; i < n; i++ {
// 		if i == k {
// 			res = current
// 			break
// 		}

// 		if current*10 <= n {
// 			current *= 10
// 		} else {
// 			if current >= n {
// 				current /= 10
// 			}

// 			current++

// 			for current%10 == 0 {
// 				current /= 10
// 			}
// 		}
// 	}

// 	return res
// }

// First attemp. Memory Limit Exceeded. 39/69 testcases passed
// func findKthNumber(n int, k int) int {
// 	res := []int{}
// 	current := 1

// 	for len(res) < n {
// 		res = append(res, current)

// 		if current*10 <= n {
// 			current *= 10
// 		} else {
// 			if current >= n {
// 				current /= 10
// 			}

// 			current++

// 			for current%10 == 0 {
// 				current /= 10
// 			}
// 		}
// 	}

// 	return res[k-1]
// }
