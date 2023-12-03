package main

import "fmt"

func search(nums []int, target int) int {
	splitPoint := findSplitPoint(nums, 0, len(nums)-1)
	if splitPoint == -1 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}
	res := binarySearch(nums, 0, splitPoint-1, target)
	if res != -1 {
		return res
	}
	return binarySearch(nums, splitPoint, len(nums)-1, target)
}

func findSplitPoint(nums []int, left, right int) int {
	if right-left < 1 {
		return -1
	}
	if right-left < 2 {
		if nums[left] > nums[right] {
			return right
		}
		return -1
	}
	idx := (right + left) / 2
	if nums[idx] > nums[idx+1] {
		return idx + 1
	}
	res := findSplitPoint(nums, left, idx)
	if res != -1 {
		return res
	}
	return findSplitPoint(nums, idx+1, right)
}

func binarySearch(nums []int, left, right, target int) int {
	for pos := (left + right) / 2; left <= right; pos = (left + right) / 2 {
		if target == nums[pos] {
			return pos
		}
		if target < nums[pos] {
			right = pos - 1
			continue
		}
		if target > nums[pos] {
			left = pos + 1
			continue
		}
	}

	return -1
}

func main() {
	//fmt.Println(search([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 0, 1, 2, 3}, 0))
	//fmt.Println(search([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 0, 1, 2, 3}, 4))
	//fmt.Println(search([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 0, 1, 2, 3}, 100))
	//fmt.Println(search([]int{4}, 4))
	fmt.Println(search([]int{1, 3}, 1))
}
