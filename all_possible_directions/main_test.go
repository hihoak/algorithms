package all_possible_directions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunc(t *testing.T) {
	cases := []struct {
		n   int
		m   int
		res int64
	}{
		{1, 1, 1},
		{1, 3, 1},
		{2, 2, 2},
		{3, 2, 3},
		{3, 3, 6},
		{18, 18, 2333606220},
	}

	for _, tc := range cases {
		require.Equal(t, tc.res, allPossibleDirections(tc.n, tc.m))
	}
}
