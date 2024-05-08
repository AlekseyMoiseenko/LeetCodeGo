package leetcode

import (
	"fmt"
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	res := findRelativeRanks([]int{10, 3, 8, 9, 4})
	fmt.Printf("res: %v\n", res)
}
