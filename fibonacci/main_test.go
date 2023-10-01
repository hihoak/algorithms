package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunc(t *testing.T) {
	cases := []struct {
		n   int
		res int64
	}{
		{1, 1},
		{2, 1},
		{60, 1548008755920},
	}

	for _, tc := range cases {
		require.Equal(t, tc.res, fib(tc.n))
	}
}
