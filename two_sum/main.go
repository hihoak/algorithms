// https://leetcode.com/problems/two-sum/

package main

import (
	"bufio"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers, sum, err := getData()
	if err != nil {
		color.HiRed(err.Error())
		return
	}

	res := twoSum(numbers, sum)
	color.HiGreen("Result is: %v", res)
}

func getData() ([]int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	color.HiCyan("Input array of numbers:")
	strNumbers, err := reader.ReadString('\n')
	if err != nil {
		return nil, 0, err
	}
	numbersSliceStr := strings.Split(strings.TrimSpace(strNumbers), " ")

	numbers := make([]int, len(numbersSliceStr))
	for idx, numberStr := range numbersSliceStr {
		numbers[idx], err = strconv.Atoi(strings.TrimSpace(numberStr))
		if err != nil {
			return nil, 0, err
		}
	}

	color.HiCyan("Input sum to find:")
	sumStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, 0, err
	}
	sumStr = strings.TrimSpace(sumStr)
	sum, err := strconv.Atoi(sumStr)
	if err != nil {
		return nil, 0, err
	}

	return numbers, sum, nil
}

func twoSum(nums []int, target int) []int {
	mapCompletions := make(map[int]int)
	for idx, val := range nums {
		if jdx, ok := mapCompletions[target - val]; ok {
			return []int{idx, jdx}
		} else {
			mapCompletions[val] = idx
		}
	}
	return nil
}
