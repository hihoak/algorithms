package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("234"))
}

var mapNumberLetters = map[int][]byte {
	2: {'a', 'b', 'c'},
	3: {'d', 'e', 'f'},
	4: {'g', 'h', 'i'},
	5: {'j', 'k', 'l'},
	6: {},
}

func letterCombinations(digits string) []string {
	var allLettersComb [][]byte
	length := 1
	for idx, ch := range digits {
		allLettersComb = append(allLettersComb, mapNumberLetters[int(ch-'0')])
		length *= len(allLettersComb[idx])
	}

	res := make([]string, length)
	for idx := 0; idx < len(allLettersComb); idx++ {
		for jdx := 0; jdx < len(allLettersComb[idx]); jdx++ {
			if idx == 0 {
				start := jdx * length / len(allLettersComb[idx])
				ends := start + length / len(allLettersComb[idx])
				for kdx := start; kdx < ends; kdx ++ {
					res[kdx] += string(allLettersComb[idx][jdx])
				}
			} else {
				for kdx := jdx; kdx < length; kdx += length / len(allLettersComb[idx]) {
					res[kdx] += string(allLettersComb[idx][jdx])
				}
			}
		}
	}
	return res
}

