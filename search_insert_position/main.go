package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5, 6, 70, 80, 90, 100}, 5))
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for idx := (right + left) / 2; ; idx = (right + left) / 2 {
		if nums[idx] == target {
			return idx
		}
		if idx == len(nums)-1 && target > nums[idx] {
			return len(nums)
		}
		if idx == 0 && target < nums[idx] {
			return 0
		}
		if target > nums[idx] && target < nums[idx+1] {
			return idx + 1
		}

		if target < nums[idx] {
			right = idx - 1
		}
		if target > nums[idx] {
			left = idx + 1
		}
	}
}
