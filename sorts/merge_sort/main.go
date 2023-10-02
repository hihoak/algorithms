package main

import "fmt"

func merge(leftArr, rightArr []int) []int {
	res := make([]int, 0, len(leftArr)+len(rightArr))
	var left, right int
	for left < len(leftArr) || right < len(rightArr) {
		if right >= len(rightArr) || (left < len(leftArr) && leftArr[left] < rightArr[right]) {
			res = append(res, leftArr[left])
			left++
		} else {
			res = append(res, rightArr[right])
			right++
		}
	}
	return res
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	leftArr := mergeSort(arr[:len(arr)/2])
	rightArr := mergeSort(arr[len(arr)/2:])

	return merge(leftArr, rightArr)
}

func main() {
	fmt.Println(merge([]int{1}, []int{}))
	fmt.Println(merge([]int{-10, 1, 2, 3, 5, 7, 8}, []int{-1, 0, 2, 3, 4, 5, 6}))
	fmt.Println(1 / 2)
	fmt.Println(mergeSort([]int{1, 2, -1, 0, 2, 5, 10, -9, 2, 5, -1, -3}))
}
