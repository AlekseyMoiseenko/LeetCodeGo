package main

// Time Complexity: O(log n)
// Space Complexity: O(1)
func mySqrt(x int) int {
	left, right := 0, x+1

	for left < right {
		mid := left + (right-left)/2

		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left - 1
}

// // Sqrt вычисляет квадратный корень методом Ньютона
// func Sqrt(x float64) float64 {
// 	if x < 0 {
// 		return math.NaN()
// 	}
// 	if x == 0 {
// 		return 0
// 	}

// 	z := x / 2.0
// 	// Задаем точность (эпсилон)
// 	const epsilon = 1e-15

// 	for {
// 		nextZ := (z + x/z) / 2.0
// 		// Если разница между текущим и следующим шагом меньше epsilon, мы нашли корень
// 		if math.Abs(z-nextZ) < epsilon {
// 			break
// 		}
// 		z = nextZ
// 	}

// 	return z
// }
