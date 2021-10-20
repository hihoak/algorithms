package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// maximum continuous sum in array
	numbers := getNumbers()

	if len(numbers) == 0 {
		numbers = generaRandomArr(100, 10)
		fmt.Println("Using random generated arr", numbers)
	}

	fmt.Printf("Result is %d", run(numbers))
}

func getNumbers() []int {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		panic("Something goes wrong")
	}

	numbers := make([]int, 0)
	for _, strNumber := range strings.Split(s, " ") {
		if number, err := strconv.Atoi(strings.TrimSpace(strNumber)); err == nil {
			numbers = append(numbers, number)
		} else {
			fmt.Printf("Bad integer %s\n", strNumber)
		}
	}
	return numbers
}

func generaRandomArr(maxNumber, size int) []int {
	numbers := make([]int, 0, size)
	rand.Seed(time.Now().UnixNano())
	count := 0
	for count < size {
		numbers = append(numbers, rand.Intn(maxNumber * 2) - maxNumber)
		count += 1
	}
	return numbers
}

func run(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	currentSum := arr[0]
	maxSum := arr[0]
	if maxSum < 0 {
		currentSum = 0
	}
	for _, val := range arr[1:] {
		currentSum += val
		if maxSum < currentSum {
			maxSum = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}