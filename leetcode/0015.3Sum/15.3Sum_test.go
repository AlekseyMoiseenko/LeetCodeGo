package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		nums    []int
		equalTo [][]int
	}{
		{
			nums:    []int{-1, 0, 1, 2, -1, -4},
			equalTo: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:    []int{0, 0, 0, 0},
			equalTo: [][]int{{0, 0, 0}},
		},
		{
			nums:    []int{0, 1, 1},
			equalTo: [][]int{},
		},
		{
			nums:    []int{0, 0, 0},
			equalTo: [][]int{{0, 0, 0}},
		},
	}

	for _, tc := range testCases {
		res := threeSum(tc.nums)
		require.Equal(t, res, tc.equalTo)
	}
}
