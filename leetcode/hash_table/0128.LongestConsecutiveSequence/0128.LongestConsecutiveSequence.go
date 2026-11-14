package main

// Time Complexity: O(n)
// Space Complexity: O(n)
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	max := 0
	for k := range m {
		if _, exist := m[k-1]; exist {
			continue
		}

		curr := k
		length := 0
		for {
			if _, exist := m[curr]; exist {
				curr++
				length++
			} else {
				if max < length {
					max = length
				}
				break
			}
		}
	}

	return max
}
