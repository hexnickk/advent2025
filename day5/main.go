package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Data struct {
	Ranges [][2]int
	Inputs []int
}

func getData(filename string) (Data, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return Data{}, err
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	ranges := make([][2]int, 0)
	for index := 0; lines[index] != ""; index++ {
		parts := strings.Split(lines[index], "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, [2]int{start, end})
	}

	inputs := make([]int, 0)
	for i := len(ranges) + 1; i < len(lines); i++ {
		num, _ := strconv.Atoi(lines[i])
		inputs = append(inputs, num)
	}
	return Data{ranges, inputs}, nil
}

func isInRanges(ranges [][2]int, num int) bool {
	for _, r := range ranges {
		if num >= r[0] && num <= r[1] {
			return true
		}
	}
	return false
}

func task1(data Data) {
	count := 0
	for _, num := range data.Inputs {
		if isInRanges(data.Ranges, num) {
			count++
		}
	}
	fmt.Println(count)
}

func task2(data Data) {
	breakpoints := make([]int, 0)
	for _, r := range data.Ranges {
		breakpoints = append(breakpoints, r[0], r[1])
	}
	slices.Sort(breakpoints)

	count := 0
	for i := 0; i < len(breakpoints)-1; i++ {
		if breakpoints[i] == breakpoints[i+1] {
			continue
		}
		// calc start
		if isInRanges(data.Ranges, breakpoints[i]) {
			count += 1
		}

		// calc in between
		mid := (breakpoints[i] + breakpoints[i+1]) / 2
		if isInRanges(data.Ranges, mid) {
			count += breakpoints[i+1] - breakpoints[i] - 1
		}
	}

	// calc last point
	count += 1
	fmt.Println(count)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <task name> <input file>")
		os.Exit(1)
	}

	taskname := os.Args[1]
	filename := os.Args[2]

	data, err := getData(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading data:", err)
		os.Exit(1)
	}

	if taskname == "1" {
		task1(data)
	}
	if taskname == "2" {
		task2(data)
	}
}
