package happy_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunc(t *testing.T) {
	cases := []struct {
		number int
		res    bool
	}{
		{
			number: 7,
			res:    true,
		},
		{
			number: 1,
			res:    true,
		},
		{
			number: 1111111,
			res:    true,
		},
		{
			number: 19,
			res:    true,
		},
		{
			number: 2,
			res:    false,
		},
	}

	for _, tc := range cases {
		require.Equalf(t, tc.res, isHappy(tc.number), "failed %d", tc.number)
	}
}
