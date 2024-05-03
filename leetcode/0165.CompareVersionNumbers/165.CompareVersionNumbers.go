package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	l1 := len(v1)
	l2 := len(v2)

	for i := 0; i < max(l1, l2); i++ {
		val1, val2 := 0, 0

		if i < l1 {
			val1, _ = strconv.Atoi(v1[i])
		}
		if i < l2 {
			val2, _ = strconv.Atoi(v2[i])
		}

		if val1 > val2 {
			return 1
		} else if val2 > val1 {
			return -1
		}
	}

	return 0
}

// first app
// func compareVersion(version1 string, version2 string) int {
// 	v1 := strings.Split(version1, ".")
// 	v2 := strings.Split(version2, ".")

// 	res := 0
// 	l := 0

// 	if len(v1) > len(v2) {
// 		l = len(v1)
// 	} else {
// 		l = len(v2)
// 	}

// 	for i := 0; i < l; i++ {
// 		if i >= len(v1) {
// 			for j := i; j < len(v2); j++ {
// 				val2, _ := strconv.Atoi(v2[j])
// 				if val2 > 0 {
// 					res = -1
// 					break
// 				}
// 			}
// 			break
// 		} else if i >= len(v2) {
// 			for j := i; j < len(v1); j++ {
// 				val1, _ := strconv.Atoi(v1[j])
// 				if val1 > 0 {
// 					res = 1
// 					break
// 				}
// 			}
// 			break
// 		}

// 		val1, _ := strconv.Atoi(v1[i])
// 		val2, _ := strconv.Atoi(v2[i])

// 		if val1 > val2 {
// 			res = 1
// 			break
// 		} else if val2 > val1 {
// 			res = -1
// 			break
// 		}
// 	}

// 	return res
// }
