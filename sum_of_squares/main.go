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
	number, err := getData()
	if err != nil {
		color.Red("%v", err)
	}
	numbers, err := Map(solve(number), strconv.Itoa)
	color.Green("Min count of numbers needed to calculate number is '%d'.\n%s = %d",
		len(numbers), strings.Join(numbers, "² + ") + "²", number)
}

func getData() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	color.Cyan("Input number to calculate")

	strNumber, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error while reading from stdin. Error %v", err)
	}
	strNumber = strings.TrimSpace(strNumber)

	number, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0, fmt.Errorf("error while parsing int from string. Error %v", err)
	}

	return number, nil
}

func Map(arr []int, f func(int) string) ([]string, error) {
	res := make([]string, len(arr))
	for idx, val := range arr {
		res[idx] = f(val)
	}
	return res, nil
}

func solve(number int) []int {
	res := make([]int, 0)
	for number != 0 {
		val := int(math.Sqrt(float64(number)))
		number -= int(math.Pow(float64(val), 2))
		res = append(res, val)
	}
	return res
}
