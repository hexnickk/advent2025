package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData(filename string) ([][]rune, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	data := make([][]rune, 0)
	for _, line := range lines {
		runes := make([]rune, 0)
		for _, char := range line {
			runes = append(runes, char)
		}
		data = append(data, runes)
	}
	return data, nil
}

func findMaxIndex(data []rune, start int, end int) int {
	maxPos := start
	for i := start; i < end; i++ {
		num1, _ := strconv.Atoi(string(data[i]))
		num2, _ := strconv.Atoi(string(data[maxPos]))
		if num1 > num2 {
			maxPos = i
		}
	}
	return maxPos
}

func task1(data [][]rune) {
	fmt.Println(data)
	sum := 0
	for _, line := range data {
		maxPos1 := findMaxIndex(line, 0, len(line)-1)
		maxPos2 := findMaxIndex(line, maxPos1+1, len(line))
		str := string(line[maxPos1]) + string(line[maxPos2])
		num, _ := strconv.Atoi(str)
		fmt.Println(line, "\t", num)
		sum += num
	}
	fmt.Println("Sum ->", sum)
}

func task2(data [][]rune) {
	sum := 0
	for _, line := range data {
		lastMaxPos := 0
		numStr := ""
		for i := 11; i >= 0; i-- {
			newMaxPos := findMaxIndex(line, lastMaxPos, len(line)-i)
			numStr += string(line[newMaxPos])
			lastMaxPos = newMaxPos + 1
		}
		num, _ := strconv.Atoi(numStr)
		fmt.Println(line, "\t", num)
		sum += num
	}
	fmt.Println("Sum ->", sum)
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
