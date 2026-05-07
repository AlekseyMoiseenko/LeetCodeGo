package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Constraints:
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10^6 <= nums1[i], nums2[i] <= 10^6
func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			nums1:    []int{1, 1, 3},
			nums2:    []int{1, 2},
			expected: 1.0,
		},
		{
			nums1:    []int{1, 1, 3},
			nums2:    []int{2},
			expected: 1.5,
		},
		{
			nums1:    []int{1, 1, 1},
			nums2:    []int{},
			expected: 1.0,
		},
		{
			nums1:    []int{},
			nums2:    []int{1, 1, 1},
			expected: 1.0,
		},
		{
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.0,
		},
		{
			nums1:    []int{1, 2, 3, 4},
			nums2:    []int{},
			expected: 2.5,
		},
		{
			nums1:    []int{-1000000},
			nums2:    []int{1000000},
			expected: 0.0,
		},
		{
			nums1:    []int{-10, -1},
			nums2:    []int{1, 10},
			expected: 0.0,
		},
		{
			nums1:    []int{1, 2, 7, 10, 15, 20},
			nums2:    []int{1, 1, 3, 4, 5, 6},
			expected: 4.5,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected, findMedianSortedArrays(tc.nums1, tc.nums2))
	}
}
