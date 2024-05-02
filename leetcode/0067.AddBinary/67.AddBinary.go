package leetcode

import "math/big"

func addBinary(a string, b string) string {
	num1, _ := new(big.Int).SetString(a, 2)
	num2, _ := new(big.Int).SetString(b, 2)
	num1.Add(num1, num2)
	return num1.Text(2)
}

// doesn't work with large values
// func addBinary(a string, b string) string {
// num1, _ := strconv.ParseInt(a, 2, 64)
// num2, _ := strconv.ParseInt(b, 2, 64)
// return strconv.FormatInt((num1 + num2), 2)
// }
