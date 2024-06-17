package leetcode

import "math"

func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		sum := l*l + r*r
		if sum == c {
			return true
		} else if sum > c {
			r--
		} else {
			l++
		}
	}
	return false
}

// first app time limit
// func judgeSquareSum(c int) bool {
// 	for i := 0; i <= c; i++ {
// 		for j := i; j <= c; j++ {
// 			if i*i+j*j == c {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }
