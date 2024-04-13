package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	indices := twoSum([]int{2, 7, 11, 15}, 9)
	require.Equal(t, indices, []int{0, 1})

	indices = twoSum([]int{3, 2, 4}, 6)
	require.Equal(t, indices, []int{1, 2})

	indices = twoSum([]int{3, 3}, 6)
	require.Equal(t, indices, []int{0, 1})

	indices = twoSum([]int{1, 2, 3, 4}, 15)
	require.Nil(t, indices)
}
