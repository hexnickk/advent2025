package main

import (
	"fmt"
	"os"
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

func countNeighbours(data [][]rune, row int, col int) int {
	count := 0

	// top row
	if row > 0 && data[row-1][col] == '@' {
		count++
	}
	if row > 0 && col > 0 && data[row-1][col-1] == '@' {
		count++
	}
	if row > 0 && col < len(data[0])-1 && data[row-1][col+1] == '@' {
		count++
	}

	// same row
	if col > 0 && data[row][col-1] == '@' {
		count++
	}
	if col < len(data[0])-1 && data[row][col+1] == '@' {
		count++
	}

	// bottom row
	if row < len(data)-1 && col > 0 && data[row+1][col-1] == '@' {
		count++
	}
	if row < len(data)-1 && data[row+1][col] == '@' {
		count++
	}
	if row < len(data)-1 && col < len(data[0])-1 && data[row+1][col+1] == '@' {
		count++
	}

	return count
}

func task1(data [][]rune) {
	count := 0
	for row := range data {
		for col := range data[0] {
			if data[row][col] == '@' && countNeighbours(data, row, col) < 4 {
				fmt.Print("#")
				count++
			}
			fmt.Print(string(data[row][col]))
		}
		fmt.Println()
	}

	fmt.Println(count)
}

func transformData(data [][]rune) ([][]rune, int) {
	deleted := 0
	newData := make([][]rune, len(data))
	for i := range data {
		newData[i] = make([]rune, len(data[0]))
		for j := range data[0] {
			if data[i][j] == '@' && countNeighbours(data, i, j) < 4 {
				deleted++
				newData[i][j] = '.'
			} else {
				newData[i][j] = data[i][j]
			}
		}
	}

	return newData, deleted
}

func task2(data [][]rune) {
	count := 0
	deleted := 0
	for {
		data, deleted = transformData(data)
		count += deleted
		if deleted == 0 {
			break
		}
	}

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
