package hash

var (
	romanMap = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

// Time complexity: O(n)
// Space complexity: O(1)
func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanMap[s[i]] < romanMap[s[i+1]] {
			sum -= romanMap[s[i]]
		} else {
			sum += romanMap[s[i]]
		}
	}
	return sum
}
