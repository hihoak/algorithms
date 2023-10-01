// https://leetcode.com/problems/remove-element/

package main

func main() {

}

func removeElement(nums []int, val int) int {

	for idx := 0; idx < len(nums); {
		if nums[idx] == val {
			nums = append(nums[:idx], nums[idx+1:]...)
			continue
		}
		idx++
	}
	return len(nums)
}
