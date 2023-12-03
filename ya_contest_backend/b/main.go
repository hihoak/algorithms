package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type log struct {
	day       int
	hour      int
	minute    int
	id        int
	orderType string
}

func sortLogs(logs []log) {
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].id != logs[j].id {
			return logs[i].id < logs[j].id
		}
		if logs[i].day < logs[j].day {
			return true
		}
		if logs[i].day > logs[j].day {
			return false
		}
		if logs[i].hour < logs[j].hour {
			return true
		}
		if logs[i].hour > logs[j].hour {
			return false
		}
		if logs[i].minute < logs[j].minute {
			return true
		}
		if logs[i].minute > logs[j].minute {
			return false
		}
		return false
	})
}

func calculateAllDistances(logs []log) []int {
	result := make([]int, 0)
	start := 0
	delta := 0
	currentID := -1
	for _, currentLog := range logs {
		if currentID == -1 || currentID != currentLog.id {
			if currentID != -1 {
				result = append(result, delta)
			}
			currentID = currentLog.id
			delta = 0
		}

		switch currentLog.orderType {
		case "A":
			start = currentLog.day*24*60 + currentLog.hour*60 + currentLog.minute
		case "S", "C":
			delta += currentLog.day*24*60 + currentLog.hour*60 + currentLog.minute - start
		}
	}
	return append(result, delta)
}

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	numberOfLogs, _ := strconv.Atoi(scaner.Text())
	logs := make([]log, 0, numberOfLogs)

	for idx := 0; idx < numberOfLogs; idx++ {
		scaner.Scan()
		data := strings.Split(scaner.Text(), " ")
		day, _ := strconv.Atoi(data[0])
		hour, _ := strconv.Atoi(data[1])
		minute, _ := strconv.Atoi(data[2])
		id, _ := strconv.Atoi(data[3])
		logs = append(logs, log{
			day:       day,
			hour:      hour,
			minute:    minute,
			id:        id,
			orderType: data[4],
		})
	}

	sortLogs(logs)

	for _, r := range calculateAllDistances(logs) {
		fmt.Printf("%d ", r)
	}
}
