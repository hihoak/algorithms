package _Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)
	for idx := 0; idx < len(nums)-3; idx++ {
		// break because we can't found answer
		if nums[idx] > target/4 {
			break
		}

		if idx != 0 && nums[idx] == nums[idx-1] {
			// previous result
			continue
		}

		// solving 3Sum problem here
		for jdx := idx + 1; jdx < len(nums)-2; jdx++ {
			if jdx != idx+1 && nums[jdx] == nums[jdx-1] {
				continue
			}

			// solving 2Sum Problem here
			left := jdx + 1
			right := len(nums) - 1
			for left < right {
				tempSum := nums[idx] + nums[jdx] + nums[left] + nums[right]
				switch {
				case tempSum < target:
					left = moveLeftUntilRepeated(nums, left)
				case tempSum > target:
					right = moveRightUntilRepeated(nums, right)
				case tempSum == target:
					res = append(res, []int{nums[idx], nums[jdx], nums[left], nums[right]})
					left = moveLeftUntilRepeated(nums, left)
				}
			}
		}
	}
	return res
}

// -30 -15 -10 0 5 10 15 20 25

func moveLeftUntilRepeated(nums []int, left int) int {
	left++
	for ; left < len(nums); left++ {
		if nums[left-1] != nums[left] {
			break
		}
	}
	return left
}

func moveRightUntilRepeated(nums []int, right int) int {
	right--
	for ; right >= 0; right-- {
		if nums[right] != nums[right+1] {
			break
		}
	}
	return right
}
