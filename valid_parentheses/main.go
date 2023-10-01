// https://leetcode.com/problems/valid-parentheses/

package main

import "log"

func main() {
	testCases := []struct {
		val         string
		expectedRes bool
	}{
		{
			"(((())))",
			true,
		},
		{
			"((((",
			false,
		},
		{
			"[](){}{}{}{}()()()[][][][][[{{}}]]",
			true,
		},
		{
			"[[[]]]]",
			false,
		},
		{
			"",
			true,
		},
		{
			"}",
			false,
		},
		{
			"{{{}}]",
			false,
		},
	}

	for _, tc := range testCases {
		if tc.expectedRes != isValid(tc.val) {
			log.Fatal("test with value is invalid: ", tc.val)
		}
	}
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	openBrackets := make([]rune, 0, len(s)/2)
	for _, bracket := range s {
		switch bracket {
		case '{', '(', '[':
			openBrackets = append(openBrackets, bracket)
		case '}':
			if len(openBrackets) == 0 || openBrackets[len(openBrackets)-1] != '{' {
				return false
			}
			openBrackets = openBrackets[:len(openBrackets)-1]
		case ']':
			if len(openBrackets) == 0 || openBrackets[len(openBrackets)-1] != '[' {
				return false
			}
			openBrackets = openBrackets[:len(openBrackets)-1]
		case ')':
			if len(openBrackets) == 0 || openBrackets[len(openBrackets)-1] != '(' {
				return false
			}
			openBrackets = openBrackets[:len(openBrackets)-1]
		}
	}

	if len(openBrackets) != 0 {
		return false
	}
	return true
}
