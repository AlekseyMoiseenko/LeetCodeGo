package main

func generate(n int) []string {
	res := []string{}

	var backtrack func(s string, open, close int)
	backtrack = func(s string, open, close int) {
		if len(s) == n*2 {
			res = append(res, s)
			return
		}

		if open < n {
			backtrack(s+"(", open+1, close)
		}
		if close < open {
			backtrack(s+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return res
}
