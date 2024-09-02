package leetcode

import (
	"strconv"
	"strings"
)

func fractionAddition(expression string) string {
	nums := []string{}
	last := 0
	for i, r := range expression {
		if (r == '-' || r == '+') && last != i {
			nums = append(nums, expression[last:i])
			last = i
			continue
		}
		if i+1 > len(expression)-1 {
			nums = append(nums, expression[last:i+1])
		}
	}

	result := calc(nums)
	return strconv.Itoa(result[0]) + "/" + strconv.Itoa(result[1])
}

func calc(input []string) []int {
	pairs := make([][2]int, len(input))
	for i, c := range input {
		nums := strings.Split(c, "/")
		pairs[i][0], _ = strconv.Atoi(nums[0])
		pairs[i][1], _ = strconv.Atoi(nums[1])
	}

	cd := 1
	for _, pair := range pairs {
		cd *= pair[1]
	}

	sum := 0
	for _, pair := range pairs {
		sum += pair[0] * cd / pair[1]
	}

	_gcd := gcd(sum, cd)
	return []int{sum / _gcd, cd / _gcd}
}

func gcd(a, b int) int {
	if a < 0 {
		a *= -1
	}

	for a > 0 {
		a, b = b%a, a
	}
	return b
}
