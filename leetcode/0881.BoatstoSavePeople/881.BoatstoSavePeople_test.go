package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumRescueBoats(t *testing.T) {
	testCases := []struct {
		peopleList []int // You are given an array people where people[i] is the weight of the i'th person
		boatLimit  int
		expected   int
	}{
		{
			peopleList: []int{3, 2, 2, 1},
			boatLimit:  3,
			expected:   3,
		},
		{
			peopleList: []int{3, 5, 3, 4},
			boatLimit:  5,
			expected:   4,
		},
		{
			peopleList: []int{1, 2},
			boatLimit:  3,
			expected:   1,
		},
	}

	for _, tc := range testCases {
		res := numRescueBoats(tc.peopleList, tc.boatLimit)
		require.Equal(t, tc.expected, res)
	}
}
