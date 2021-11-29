package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

const (
	parseStrToIntErrorMsg = "something goes wrong while parsing string to int. Error: %s"
	readFromStdinErrorMsg = "something goes wrong while reading from stdin. Error: %s"
)

func main() {
	lowerBound, upperBound, numbers, err := getData()
	if err != nil {
		color.HiRed("%s", err.Error())
		return
	}

	res := solve(lowerBound, upperBound, numbers)
	color.HiGreen("Result is %v", res)
}

func getNumber(reader *bufio.Reader) (int, error) {
	strNumber, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf(readFromStdinErrorMsg, err.Error())
	}
	strNumber = strings.TrimSpace(strNumber)

	number, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0, fmt.Errorf(parseStrToIntErrorMsg, err.Error())
	}

	return number, nil
}

func getSliceOfInts(reader *bufio.Reader) ([]int, error) {
	strNumbers, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf(readFromStdinErrorMsg, err.Error())
	}

	strNumbers = strings.TrimSpace(strNumbers)
	strNumbersSlice := strings.Split(strNumbers, " ")

	res := make([]int, len(strNumbersSlice))
	for idx, val := range strNumbersSlice {
		if val == " " {
			continue
		}
		res[idx], err = strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf(parseStrToIntErrorMsg, err.Error())
		}
	}

	return res, nil
}

func getData() (int, int, []int, error) {
	var lowerBound, upperBound int
	var numbers []int
	reader := bufio.NewReader(os.Stdin)
	var err error
	color.HiCyan("Input lower bound:")
	lowerBound, err = getNumber(reader)
	color.HiCyan("Input upper bound:")
	upperBound, err = getNumber(reader)
	color.HiCyan("Input array of numbers:")
	numbers, err = getSliceOfInts(reader)

	if err != nil {
		return 0, 0, nil, err
	}

	return lowerBound, upperBound, numbers, nil
}

func solve(lowerBound, upperBound int, numbers []int) [][2]int {
	res := make([][2]int, 0)

	// cases when all digits missed
	if len(numbers) == 0 || numbers[0] > upperBound || numbers[len(numbers) - 1] < lowerBound {
		res = append(res, [2]int{lowerBound, upperBound})
		return res
	}

	idx := 0

	// skipping all numbers that lower than lowerBound
	for ; numbers[idx] < lowerBound; idx++ {}

	// case when in numbers missed lowerBound
	if lowerBound != numbers[idx] && numbers[idx] <= upperBound {
		res = append(res, [2]int{lowerBound, numbers[idx] - 1})
	}

	// main loop
	closed := false
	for ; idx < len(numbers) - 1 && numbers[idx + 1] <= upperBound; idx++ {
		if numbers[idx] - numbers[idx + 1] < -1 {
			res = append(res, [2]int{numbers[idx] + 1, numbers[idx + 1] - 1})
		}
		if numbers[idx] == upperBound {
			closed = true
		}
	}

	// not need to watch cases with missed upperBound in numbers
	if closed {
		return res
	}

	if numbers[idx] < upperBound {
		res = append(res, [2]int{numbers[idx] + 1, upperBound})
	} else if numbers[idx] > upperBound {
		res = append(res, [2]int{numbers[idx - 1] + 1, upperBound})
	}

	return res
}
