// https://leetcode.com/problems/3sum/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for idx, num := range nums {
		// already can't get 0
		if num > 0 {
			break
		}

		// skip equal numbers
		if idx != 0 && num == nums[idx-1] {
			continue
		}

		// solve 2Sum problem inside cycle
		leftIdx := idx + 1
		rightIdx := len(nums) - 1

		for leftIdx < rightIdx {
			currentSum := num + nums[leftIdx] + nums[rightIdx]
			switch {
			case currentSum == 0:
				res = append(res, []int{num, nums[leftIdx], nums[rightIdx]})
				// skip equal digits on the left because of founded answer
				leftIdx = skipUntilEqualLeft(nums, leftIdx)
				// skip equal digits on the right because of founded answer
				rightIdx = skipUntilEqualRight(nums, rightIdx)
			case currentSum > 0:
				// skip digits on the right because we can't get 0 with such big numbers
				rightIdx = skipUntilEqualRight(nums, rightIdx)
			case currentSum < 0:
				// skip digits on the left because we can't get 0 with such small numbers
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
