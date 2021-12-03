package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {
	data, err := getData()
	if err != nil {
		color.HiRed(err.Error())
		return
	}

	res, err := solve(data)
	if err != nil {
		color.HiRed(err.Error())
		return
	}

	color.HiGreen("Result string is: %s", res)
}

func getData() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	color.HiCyan(("Input brackets string"))
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("something goes wrong while getting data from stdin. Error: %s", err.Error())
	}

	return strings.TrimSpace(data), nil
}

func solve(data string) (string, error) {
	var termsIndexes []int

	index := 0
	var err error = nil
	for ; index < len(data) && err == nil; {
		validTerm := false

		var stack = make([]byte, 0)
		for jdx, b := range data[index:] {
			if byte(b) == ')' && len(stack) == 0 {
				break
			}
			if byte(b) == '(' {
				stack = append(stack, byte(b))
			} else {
				stack = stack[:len(stack) - 1]
			}
			index++
			if jdx != index && len(stack) == 0 {
				validTerm = true
				break
			}
		}
		if validTerm {
			termsIndexes = append(termsIndexes, index)
			continue
		}

		if len(stack) != 0 {
			err = fmt.Errorf("error format of string in index '%d', not enoegh close brackets", index - 1)
		} else {
			err = fmt.Errorf("error format of string in index '%d', wrong place for close brackets", index)
		}
	}

	if err != nil {
		return "", err
	}

	res := make([]string, 0)
	start := 0
	for _, end := range termsIndexes {
		str := data[start + 1:end - 1]
		if str != "" {
			res = append(res, str)
		}
		start = end
	}

	return strings.Join(res, " + "), nil
}
