package main

var (
	phoneMap    = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	offsetDigit = '2'
)

// Time complexity: ( O(4^n) ), where ( n ) is the length of the input string. In the worst case, each digit can represent 4 letters, so there will be 4 recursive calls for each digit.
// Space complexity: ( O(n) ), where ( n ) is the length of the input string. This accounts for the recursion stack space.
func letterCombinations(digits string) []string {
	var output []string

	var backtrack func(combination string, nextDigits string)
	backtrack = func(combination string, nextDigits string) {
		if nextDigits == "" {
			output = append(output, combination)
		} else {
			letters := phoneMap[nextDigits[0]-byte(offsetDigit)]
			for _, letter := range letters {
				backtrack(combination+string(letter), nextDigits[1:])
			}
		}
	}

	backtrack("", digits)
	return output
}
