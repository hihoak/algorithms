package generate_parentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	generate(&res, 0, 0, "", n)
	return res
}

func generate(res *[]string, left, right int, s string, n int) {
	if len(s) == n*2 {
		*res = append(*res, s)
		return
	}
	if left < n {
		generate(res, left+1, right, s+"(", n)
	}
	if right < left {
		generate(res, left, right+1, s+")", n)
	}
}
