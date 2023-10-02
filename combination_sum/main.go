package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	recSum(candidates, target, nil, &result, 0)
	return result
}

func recSum(candidates []int, target int, subResult []int, result *[][]int, index int) {
	if target == 0 {
		*result = append(*result, append([]int{}, subResult...))
		return
	}

	for idx := index; idx < len(candidates) && target-candidates[idx] >= 0; {
		subResult = append(subResult, candidates[idx])
		recSum(candidates, target-candidates[idx], subResult, result, idx)
		idx++
		subResult = subResult[:len(subResult)-1]
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
