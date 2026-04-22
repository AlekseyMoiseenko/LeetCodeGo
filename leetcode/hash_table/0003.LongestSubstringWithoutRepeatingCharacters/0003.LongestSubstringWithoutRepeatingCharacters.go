package hash

// Time complexity: O(n)
// Space complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	runeMap := make(map[rune]int)
	offset, maxVal := 0, 0

	for i, currentChar := range s {
		if lastIdx, exist := runeMap[currentChar]; exist {
			offset = max(offset, lastIdx+1)
		}

		runeMap[currentChar] = i
		maxVal = max(maxVal, i-offset+1)
	}

	return maxVal
}
