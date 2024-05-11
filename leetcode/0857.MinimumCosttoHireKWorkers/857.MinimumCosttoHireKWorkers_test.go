package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMincostToHireWorkers(t *testing.T) {
	testCases := []struct {
		quality  []int
		wage     []int
		k        int
		expected float64
	}{
		{
			quality:  []int{10, 20, 5},
			wage:     []int{70, 50, 30},
			k:        2,
			expected: float64(105),
		},
		{
			quality:  []int{3, 1, 10, 10, 1},
			wage:     []int{4, 8, 2, 2, 7},
			k:        3,
			expected: float64(30.666666666666668),
		},
	}

	for _, tc := range testCases {
		res := mincostToHireWorkers(tc.quality, tc.wage, tc.k)

		require.Equal(t, res, tc.expected)
	}
}
