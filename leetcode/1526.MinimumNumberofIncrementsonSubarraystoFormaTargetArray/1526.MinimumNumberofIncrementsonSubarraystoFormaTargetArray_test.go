package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinNumberOperations(t *testing.T) {
	testCases := []struct {
		target   []int
		expected int
	}{
		{
			target:   []int{},
			expected: 0,
		},
		{
			target:   []int{1, 2, 3, 2, 1},
			expected: 3,
		},
		{
			target:   []int{3, 1, 1, 2},
			expected: 4,
		},
		{
			target:   []int{3, 1, 5, 4, 2},
			expected: 7,
		},
		{
			target:   []int{99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 1001},
			expected: 1010,
		},
	}

	for _, tc := range testCases {
		res := minNumberOperations(tc.target)
		fmt.Println(res)
		require.Equal(t, tc.expected, res)
	}
}
