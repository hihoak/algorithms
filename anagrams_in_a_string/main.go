package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	line, pattern, err := getData()
	if err != nil {
		fmt.Printf("Something goes wrong check error: %v", err)
		return
	}

	res := solve(line, pattern)
	fmt.Println("Indexes where I find anagram: ", res)
}

func getData() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	var sourceString, pattern string
	var err error
	fmt.Printf("Input string: ")
	if sourceString, err = reader.ReadString('\n'); err != nil {
		return "", "", err
	}
	sourceString = strings.TrimRight(sourceString, "\n")
	fmt.Printf("Input pattetn: ")
	if pattern, err = reader.ReadString('\n'); err != nil {
		return "", "", err
	}
	pattern = strings.TrimRight(pattern, "\n")
	return sourceString, pattern, nil
}

func patternToMap(pattern string) map[string]int {
	res := make(map[string]int)
	for _, val := range pattern {
		if _, ok := res[string(val)]; ok {
			res[string(val)]++
		} else {
			res[string(val)] = 1
		}
	}
	return res
}

func solve(data, pattern string) []int {
	res := make([]int, 0)

	patternMap := patternToMap(pattern)
	compareMap := make(map[string]int)
	compareSize := 0
	for i := 0; i < len(data); i++ {
		if _, ok := patternMap[string(data[i])]; !ok {
			compareMap = map[string]int{}
			continue
		}
		if _, ok := compareMap[string(data[i])]; ok {
			compareMap[string(data[i])]++
		}
		compareMap[string(data[i])] = 1
		if compareSize > len(pattern) {
			compareMap[string(data[i - len(pattern)])]--
		}

		if reflect.DeepEqual(patternMap, compareMap) {
			res = append(res, i - len(pattern) + 1)
		}
	}
	return res
}