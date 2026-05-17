package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -10^9 <= nums1[i], nums2[j] <= 10^9
func TestMerge(t *testing.T) {
	testCases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{1, 1, 1, 2, 3, 0, 0, 0},
			m:        5,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 1, 1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1, 1, 1, 2, 3, 0, 0, 0, 0, 0},
			m:        5,
			nums2:    []int{2, 5, 6, 99, 99},
			n:        5,
			expected: []int{1, 1, 1, 2, 2, 3, 5, 6, 99, 99},
		},
		{
			nums1:    []int{199, 199, 199, 199, 199, 199, 199, 199, 0, 0, 0, 0, 0},
			m:        8,
			nums2:    []int{2, 5, 6, 9, 9},
			n:        5,
			expected: []int{2, 5, 6, 9, 9, 199, 199, 199, 199, 199, 199, 199, 199},
		},
	}

	for _, tc := range testCases {
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		require.Equal(t, tc.expected, tc.nums1)
	}
}
