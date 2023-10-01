package can_sum

func canSum(canNumber int, numbers []int) bool {
	if canNumber == 0 {
		return true
	}
	cache := make(map[int]bool, len(numbers))
	return canSumRec(canNumber, numbers, &cache)
}

func canSumRec(canNumber int, numbers []int, cache *map[int]bool) bool {
	if res, ok := (*cache)[canNumber]; ok {
		return res
	}
	if canNumber == 0 {
		return true
	}
	if canNumber < 0 {
		return false
	}

	for _, number := range numbers {
		if canNumber%number == 0 {
			return true
		}
		if canSumRec(canNumber-number, numbers, cache) {
			return true
		}
		(*cache)[canNumber-number] = false
	}
	return false
}
