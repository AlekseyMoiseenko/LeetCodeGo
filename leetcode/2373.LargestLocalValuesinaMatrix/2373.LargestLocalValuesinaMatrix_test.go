package leetcode

import (
	"fmt"
	"testing"
)

func TestLargestLocal(t *testing.T) {
	res := largestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}})
	fmt.Printf("res: %v\n", res)
}
