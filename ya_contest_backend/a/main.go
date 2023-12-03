package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateHash(secondName, name, surname string, day, month, year int) string {
	uniqueWords := make(map[rune]interface{})
	// first
	uniqueWordsCount := 0
	for _, word := range secondName {
		if _, ok := uniqueWords[word]; !ok {
			uniqueWords[word] = word
			uniqueWordsCount++
		}
	}
	for _, word := range name {
		if _, ok := uniqueWords[word]; !ok {
			uniqueWords[word] = word
			uniqueWordsCount++
		}
	}
	for _, word := range surname {
		if _, ok := uniqueWords[word]; !ok {
			uniqueWords[word] = word
			uniqueWordsCount++
		}
	}

	// second
	dayAndMonthDigits := (day/10 + day%10 + month/10 + month%10) * 64

	// third
	firstSecondNameWordDigit := int(secondName[0]) - 'A' + 1
	if firstSecondNameWordDigit > 26 {
		firstSecondNameWordDigit = (int(secondName[0]) - 'a' + 1) * 256
	} else {
		firstSecondNameWordDigit = firstSecondNameWordDigit * 256
	}

	// fourth
	fourthStage := uniqueWordsCount + dayAndMonthDigits + firstSecondNameWordDigit

	// fifth
	hexFifthStage := fmt.Sprintf("%X", fourthStage)

	// sixth

	result := []byte{'0', '0', '0'}
	for jdx, idx := 0, len(hexFifthStage)-3; idx < len(hexFifthStage); jdx, idx = jdx+1, idx+1 {
		result[jdx] = hexFifthStage[idx]
	}
	return string(result)
}

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	numberOfCandidates, _ := strconv.Atoi(scaner.Text())
	hashes := make([]string, 0, numberOfCandidates)
	for idx := 0; idx < numberOfCandidates; idx++ {
		scaner.Scan()
		data := strings.Split(scaner.Text(), ",")
		day, _ := strconv.Atoi(data[3])
		month, _ := strconv.Atoi(data[4])
		year, _ := strconv.Atoi(data[5])
		hashes = append(hashes, calculateHash(data[0], data[1], data[2], day, month, year))
	}

	fmt.Println(strings.Join(hashes, " "))
}
