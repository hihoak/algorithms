package all_possible_directions

import "fmt"

func allPossibleDirections(n, m int) int64 {
	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	cache := make(map[string]int64, n+m)
	return allPossibleDirectionsRec(n-1, m, &cache) + allPossibleDirectionsRec(n, m-1, &cache)
}

func allPossibleDirectionsRec(n, m int, cache *map[string]int64) int64 {
	key1, key2 := fmt.Sprintf("%d,%d", n, m), fmt.Sprintf("%d,%d", m, n)
	if res, ok := (*cache)[key1]; ok {
		return res
	}
	if res, ok := (*cache)[key2]; ok {
		return res
	}

	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}

	res := allPossibleDirectionsRec(n-1, m, cache) + allPossibleDirectionsRec(n, m-1, cache)
	(*cache)[key1], (*cache)[key2] = res, res
	return res
}
