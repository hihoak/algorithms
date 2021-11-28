package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	color.Cyan("Input how many digits contains number. It must be between 2 ands 9 inclusive")
	size, err := getData()
	if err != nil {
		color.Red("Failed to get data. %v", err)
	}

	color.Cyan("Start finding max '%d' digit pan-number", size)
	res := solve(size)
	color.Green("Result is %d", res)
}

func getData() (int, error) {
	/// getting initial data
	reader := bufio.NewReader(os.Stdin)
	strSize, err := reader.ReadString('\n')
	if err != nil {
		err = fmt.Errorf("something goes wrong. Error: %v", err)
		return 0, err
	}
	strSize = strings.TrimSpace(strSize)

	size, err := strconv.Atoi(strSize)
	if err != nil {
		err = fmt.Errorf("got error while converting string to number. Error: %v", err)
		return 0, err
	}

	if size < 2 && size > 9 {
		return 0, fmt.Errorf("incorrect size. Please input number between 2 and 9 inclusive")
	}

	return size, nil
}

func isPrime(number int) bool {
	/// detecting that number is prime
	for i := 2; i < number / 2; i++ {
		if number % i == 0 {
			return true
		}
	}
	return false
}

func containsNotRepeatedDigits(number int) bool {
	/// detecting that number contains not repeated digits (for example 1234 - true, 1224 - false)
	digits := make([]int, 0)
	for number != 0 {
		digit := number % 10
		for _, val := range digits {
			if val == digit {
				return false
			}
		}
		digits = append(digits, digit)
		number = number / 10
	}
	return true
}

func solve(size int) int {
	/// finding max N digit number 41 Eiler exercise
	minNumber := int(math.Pow10(size - 1))
	maxNumber := int(math.Pow10(size)) - 1

	maxRes := 0
	for i := maxNumber; i >= minNumber; i-- {
		if containsNotRepeatedDigits(i) && ! isPrime(i) {
			maxRes = i
			break
		}
	}
	return maxRes
}
