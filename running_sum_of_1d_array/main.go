// https://leetcode.com/problems/running-sum-of-1d-array
package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
}

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0
	for idx, val := range nums {
		sum += val
		result[idx] = sum
	}

	return result
}
