package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData() (*string, []string, error) {
	reader := bufio.NewReader(os.Stdin)

	enteredNumber, err := reader.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}
	enteredNumber = strings.TrimSpace(enteredNumber)

	s, err := reader.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}

	validWords := make([]string, 0)
	for _, word := range strings.Split(s, " ") {
		validWords = append(validWords, strings.TrimSpace(word))
	}
	return &enteredNumber, validWords, nil
}

func realNumberLength(number string) int {
	length := 0
	for idx, _ := range number {
		if number[idx] == '0' || number[idx] == '1' {
			continue
		}
		length++
	}
	return length
}

func filterByLength(words []string, length int) []string {
	newWords := make([]string, 0)
	for _, word := range words {
		if len(word) == length {
			newWords = append(newWords, word)
		}
	}
	return newWords
}

func filterByEnteredNumber(enteredNumber string, keyboard map[int][]string, word string) (bool, error) {
	i := 0
	for idx := range enteredNumber {
		if enteredNumber[idx] == '0' || enteredNumber[idx] == '1' {
			continue
		}

		key, err := strconv.Atoi(string(enteredNumber[idx]))
		if err != nil {
			return false, err
		}
		correct := false
		for _, letter := range keyboard[key] {
			if letter == string(word[i]) {
				correct = true
				break
			}
		}

		if !correct {
			return false, nil
		}

		i++
	}
	return true, nil
}

func solve(enteredNumber string, words []string, keyboard map[int][]string) ([]string, error) {
	res := make([]string, 0)

	words = filterByLength(words, realNumberLength(enteredNumber))
	for _, word := range words {
		ok, err := filterByEnteredNumber(enteredNumber, keyboard, word)
		if err != nil {
			return nil, err
		}
		if ok {
			res = append(res, word)
		}
	}
	return res, nil
}

func main() {
	keyboard := map[int][]string{
		1: {},
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
		0: {},
	}

	enteredNumber, validWords, err := getData()
	if err != nil {
		fmt.Printf("Something goes wrong %v", err)
		return
	}

	res, err := solve(*enteredNumber, validWords, keyboard)
	if err != nil {
		fmt.Printf("Something goes wrong %v", err)
		return
	}

	fmt.Printf("You can input this words: %v", res)
}
