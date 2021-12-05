// https://leetcode.com/problems/palindrome-number/

package main

import (
	"bufio"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

func main() {
	number, err := getData()
	if err != nil {
		color.HiRed(err.Error())
		return
	}

	res := solve(number)
	color.HiGreen("Result is %t", res)
}

func getData() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	color.HiCyan("Input number:")
	numStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(strings.TrimSpace(numStr))
	if err != nil {
		return 0, err
	}

	return num, nil
}

func solve(num int) bool {
	if num < 0 || (num % 10 == 0 && num != 0) {
		return false
	}

	// sliceOfNum := make([]int, 0)
	count := 1
	revertedNumber := 0
	tempNum := num
	for {
		shortNum := tempNum / 10
		revertedNumber += (tempNum - shortNum * 10) * count
		tempNum = shortNum

		if shortNum == 0 {
			break
		}
		revertedNumber *= 10
	}

/*	start, finish := 0, len(sliceOfNum)
	for start < finish {
		if sliceOfNum[start] != sliceOfNum[finish - 1] {
			return false
		}
		start++
		finish--
	}*/
	return num == revertedNumber
}