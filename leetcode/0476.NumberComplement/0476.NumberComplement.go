package leetcode

import (
	"strconv"
	"strings"
)

func findComplement(num int) int {
	s := strconv.FormatInt(int64(num), 2)
	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		if s[i] == 49 {
			sb.WriteByte(48)
		} else {
			sb.WriteByte(49)
		}
	}

	res, _ := strconv.ParseInt(sb.String(), 2, 64)

	return int(res)
}
