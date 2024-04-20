package leetcode

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	max := 0
	runeMap := make(map[rune]struct{})

	for len(runes) > 0 {
		for _, v := range runes {
			if _, contains := runeMap[v]; contains {
				break
			}

			runeMap[v] = struct{}{}
		}
		if len(runeMap) > max {
			max = len(runeMap)
		}
		runes = runes[1:]

		for k := range runeMap {
			delete(runeMap, k)
		}
	}

	return max
}

// first approach
// func lengthOfLongestSubstring(s string) int {
// 	runes := []rune(s)
// 	max := 0

// 	for len(runes) > 0 {
// 		runeMap := make(map[rune]struct{})
// 		for _, v := range runes {
// 			if _, contains := runeMap[v]; contains {
// 				break
// 			}

// 			runeMap[v] = struct{}{}
// 		}
// 		if len(runeMap) > max {
// 			max = len(runeMap)
// 		}
// 		runes = runes[1:]
// 	}

// 	return max
// }
