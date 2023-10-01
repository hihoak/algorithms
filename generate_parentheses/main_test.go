package generate_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunc(t *testing.T) {
	cases := []struct {
		depth  int
		result []string
	}{
		{
			1,
			[]string{"()"},
		},
		{
			2,
			[]string{"(())", "()()"},
		},
		{
			3,
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tc := range cases {
		require.Equal(t, tc.result, generateParenthesis(tc.depth))
	}
}
