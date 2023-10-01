package can_sum

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunc(t *testing.T) {
	cases := []struct {
		number  int
		numbers []int
		res     bool
	}{
		{1, []int{1, 2, 3, 4}, true},
		{2, []int{1}, true},
		{7, []int{2, 3}, true},
		{7, []int{2, 4}, false},
		{300, []int{7, 14}, false},
		{100000000, []int{2, 4, 6, 8}, true},
		//{100000000, []int{3}, false},
	}

	for idx, tc := range cases {
		require.Equalf(t, tc.res, canSum(tc.number, tc.numbers), fmt.Sprintf("%d", idx))
	}
}
