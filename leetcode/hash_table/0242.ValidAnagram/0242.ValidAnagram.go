package hash

import "slices"

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := []byte(s)
	s2 := []byte(t)
	slices.Sort(s1)
	slices.Sort(s2)
	return slices.Equal(s1, s2)
}

// first approach
// Time complexity: O(n)
// Space complexity: O(n)
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
