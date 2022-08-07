// https://leetcode.com/problems/3sum/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{0,0,0,0,0}))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for idx, num := range nums {
		if num > 0 {
			break
		}

		if idx != 0 && num == nums[idx - 1] {
			continue
		}

		leftIdx := idx + 1
		rightIdx := len(nums) - 1

		for leftIdx < rightIdx {
			currentSum := num + nums[leftIdx] + nums[rightIdx]
			switch {
			case currentSum == 0:
				res = append(res, []int{num, nums[leftIdx], nums[rightIdx]})
				leftIdx = skipUntilEqualLeft(nums, leftIdx)
				rightIdx = skipUntilEqualRight(nums, rightIdx)
			case currentSum > 0:
				rightIdx = skipUntilEqualRight(nums, rightIdx)
			case currentSum < 0:
				leftIdx = skipUntilEqualLeft(nums, leftIdx)
			}
		}
	}

	return res
}

func skipUntilEqualLeft(nums []int, index int) int {
	for idx := index + 1; idx < len(nums); idx++ {
		if nums[idx] != nums[index] {
			return idx
		}
	}
	return len(nums)
}

func skipUntilEqualRight(nums []int, index int) int {
	for idx := index - 1; idx >= 0; idx-- {
		if nums[idx] != nums[index] {
			return idx
		}
	}
	return 0
}
