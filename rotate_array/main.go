package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumbers() (*int, []int, error) {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}

	count, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return nil, nil, err
	}

	s, err = reader.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}

	numbers := make([]int, 0)
	for _, strNumber := range strings.Split(s, " ") {
		if number, err := strconv.Atoi(strings.TrimSpace(strNumber)); err == nil {
			numbers = append(numbers, number)
		} else {
			return nil, nil, err
		}
	}
	return &count, numbers, nil
}

func swap(arr []int, i int, j int) []int {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
	return arr
}

func solve(arr []int, count int) []int {
	realCount := count % len(arr)
	if len(arr) <= 1 || realCount == 0 {
		return arr
	}

	for i := realCount; i < len(arr); i++ {
		for j := i; j > i - realCount; j-- {
			arr = swap(arr, j, j - 1)
		}
	}
	return arr
}

func main() {
	count, arr, err := getNumbers()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	res := solve(arr, *count)
	if err != nil {
		fmt.Printf("Something goes wrong, %v", err)
		return
	}
	fmt.Printf("All done.\nResult array: %v", res)
}
