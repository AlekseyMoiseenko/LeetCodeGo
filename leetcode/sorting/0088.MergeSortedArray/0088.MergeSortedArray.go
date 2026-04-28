package sorting

// Time Complexity: O(m+n+k)
// Space Complexity: O(k)
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 {
		return
	}

	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}

	min, max := nums1[0], nums1[0]
	for i := 0; i < len(nums1); i++ {
		if nums1[i] < min {
			min = nums1[i]
		}
		if nums1[i] > max {
			max = nums1[i]
		}
	}

	counter := make([]int, max-min+1)
	for i := 0; i < len(nums1); i++ {
		counter[nums1[i]-min]++
	}

	index := 0
	for i, v := range counter {
		for j := 0; j < v; j++ {
			nums1[index] = i + min
			index++
		}
	}
}

// first approach
// time: O(m+n+k)
// space: O(k)
// func merge1(nums1 []int, m int, nums2 []int, n int) {
// 	if len(nums1) == 0 {
// 		return
// 	}

// 	min, max := nums1[0], nums1[0]
// 	for i := 0; i < m; i++ {
// 		if nums1[i] < min {
// 			min = nums1[i]
// 		}
// 		if nums1[i] > max {
// 			max = nums1[i]
// 		}
// 	}

// 	for i := 0; i < n; i++ {
// 		if nums2[i] < min {
// 			min = nums2[i]
// 		}
// 		if nums2[i] > max {
// 			max = nums2[i]
// 		}
// 	}

// 	counter := make([]int, max-min+1)
// 	for i := 0; i < m; i++ {
// 		counter[nums1[i]-min]++
// 	}
// 	for i := 0; i < n; i++ {
// 		counter[nums2[i]-min]++
// 	}

// 	index := 0
// 	for i, v := range counter {
// 		for j := 0; j < v; j++ {
// 			nums1[index] = i + min
// 			index++
// 		}
// 	}
// }
