package happy_number

func isHappy(n int) bool {
	res := map[int]interface{}{}
	for n != 1 {
		n = calculateSquares(n)
		if _, ok := res[n]; ok {
			return false
		}
		res[n] = struct{}{}
	}
	return true
}

func calculateSquares(n int) int {
	res := 0
	for n != 0 {
		i := n % 10
		res += i * i
		n /= 10
	}
	return res
}
