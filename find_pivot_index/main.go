// https://leetcode.com/problems/find-pivot-index
package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{2}))
}

func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := 0
	for _, val := range nums {
		rightSum += val
	}

	idx := 0
	for idx < len(nums) {
		rightSum -= nums[idx]
		if leftSum == rightSum {
			return idx
		}
		leftSum += nums[idx]
		idx++
	}

	return -1
}
