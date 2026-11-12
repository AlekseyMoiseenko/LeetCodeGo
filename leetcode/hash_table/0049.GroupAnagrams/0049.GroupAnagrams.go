package main

import "slices"

// Time Complexity: O(N * M log M)
// Space Complexity: O(N * M)
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		slices.Sort(s)
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}

	res := make([][]string, len(m))
	i := 0
	for _, v := range m {
		res[i] = v
		i++
	}

	return res
}
