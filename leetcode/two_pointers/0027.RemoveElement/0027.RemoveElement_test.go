package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// 0 <= nums.length <= 100
// 0 <= nums[i] <= 50
// 0 <= val <= 100
func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums         []int
		val          int
		expectedNums []int
	}{
		{
			nums:         []int{3, 2, 2, 3},
			val:          3,
			expectedNums: []int{2, 2},
		},
		{
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			expectedNums: []int{0, 1, 3, 0, 4},
		},
	}

	for _, tc := range testCases {
		newLen := removeElement(tc.nums, tc.val)
		require.Equal(t, len(tc.expectedNums), newLen)
		require.Equal(t, tc.expectedNums, tc.nums[:newLen])
	}
}
