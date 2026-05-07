package main

const (
	maxVal = 1000000
	minVal = -1000000
)

// Time Complexity: O(log(min(l1,l2)))
// Space Complexity: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) { // Гарантируем, что nums1 – меньший массив
		nums1, nums2 = nums2, nums1
	}

	l1, l2 := len(nums1), len(nums2)
	total := l1 + l2
	half := (total + 1) / 2 // количество элементов в левой части

	result := 0.0
	low, high := 0, l1
	for low <= high {
		// i – сколько элементов взять из nums1 в левую часть
		i := (low + high) / 2
		// j – сколько элементов взять из nums2 в левую часть
		j := half - i

		// Определяем четыре ключевых значения (с учётом границ)
		var maxLeftA, minRightA int
		if i == 0 {
			maxLeftA = minVal
		} else {
			maxLeftA = nums1[i-1]
		}
		if i == l1 {
			minRightA = maxVal
		} else {
			minRightA = nums1[i]
		}

		var maxLeftB, minRightB int
		if j == 0 {
			maxLeftB = minVal
		} else {
			maxLeftB = nums2[j-1]
		}
		if j == l2 {
			minRightB = maxVal
		} else {
			minRightB = nums2[j]
		}

		// Проверяем условие корректного разделения
		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			// Нашли правильное разделение
			if total%2 == 1 {
				// Нечётная длина – медиана это максимум левой части
				result = float64(max(maxLeftA, maxLeftB))
				break
			} else {
				// Чётная длина – среднее двух центральных элементов
				leftMax := max(maxLeftA, maxLeftB)
				rightMin := min(minRightA, minRightB)
				result = float64(leftMax+rightMin) / 2.0
				break
			}
		} else if maxLeftA > minRightB {
			// i слишком большое, нужно взять меньше элементов из nums1
			high = i - 1
		} else {
			// i слишком маленькое, нужно взять больше элементов из nums1
			low = i + 1
		}
	}

	return result
}
