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
	inputStrError = "something goes wrong while reading string. Error %s"
	parseStrError = "something goes wrong while parsing string to int. Error %s"
)

func main() {
	works, waitTime, err := getData()
	if err != nil {
		color.HiRed("%s", err.Error())
		return
	}
	color.Green("Result time is %d.", solve(works, waitTime))
}

func prepareData(works []string) (map[string]int, int) {
	res := make(map[string]int)
	max := 0
	for _, work := range works {
		if count, ok := res[work]; ok {
			count++
			res[work] = count
			if count > max {
				max = count
			}
		} else {
			res[work] = 1
		}
	}
	return res, max
}

func getData() ([]string, int, error) {
	reader := bufio.NewReader(os.Stdin)
	color.HiCyan("Input wait time between tasks:")
	waitTimeStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, 0, fmt.Errorf(inputStrError, err.Error())
	}

	waitTime, err := strconv.Atoi(strings.TrimSpace(waitTimeStr))
	if err != nil {
		return nil, 0, fmt.Errorf(parseStrError, err.Error())
	}

	color.HiCyan("Input works to do:")
	worksStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, 0, fmt.Errorf(inputStrError, err.Error())
	}

	works := strings.Split(worksStr, " ")
	resWorks := make([]string, len(works))
	for idx, work := range works {
		resWorks[idx] = strings.TrimSpace(work)
	}

	return resWorks, waitTime, nil
}

func solve(works []string, waitTime int) int {
	mapWorks := make(map[string]int)
	max := 0
	for _, work := range works {
		if count, ok := mapWorks[work]; ok {
			count++
			mapWorks[work] = count
			if count > max {
				max = count
			}
		} else {
			mapWorks[work] = 1
		}
	}
	totalTime := len(works)

	restWorksCount := len(works) - max

	countOfWaitTimes := max - restWorksCount

	if countOfWaitTimes > 1 {
		totalTime += (countOfWaitTimes - 1) * waitTime
	}
	return totalTime
}
