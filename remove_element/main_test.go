package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunc(t *testing.T) {
	cases := []struct {
		nums         []int
		valToRemove  int
		expectedNums []int
		expectedRes  int
	}{
		{
			nums:         []int{1, 2, 3, 4, 5},
			valToRemove:  2,
			expectedNums: []int{1, 3, 4, 5},
			expectedRes:  4,
		},
		{
			nums:         []int{},
			valToRemove:  2,
			expectedNums: []int{},
			expectedRes:  0,
		},
		{
			nums:         []int{1, 1, 1, 1, 1},
			valToRemove:  1,
			expectedNums: []int{},
			expectedRes:  0,
		},
		{
			nums:         []int{1, 2, 2, 2, 2, 1},
			valToRemove:  2,
			expectedNums: []int{1, 1},
			expectedRes:  2,
		},
		{
			nums:         []int{1, 2, 2, 2, 2, 1},
			valToRemove:  1,
			expectedNums: []int{2, 2, 2, 2},
			expectedRes:  4,
		},
		{
			nums:         []int{1, 2, 1, 2, 2, 1},
			valToRemove:  1,
			expectedNums: []int{2, 2, 2},
			expectedRes:  3,
		},
	}

	for _, tc := range cases {
		require.Equal(t, tc.expectedRes, removeElement(tc.nums, tc.valToRemove))
		require.Equal(t, tc.expectedNums, tc.nums[:tc.expectedRes])
	}
}
