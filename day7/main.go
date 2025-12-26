package main

import (
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Chars [][]string
}

func getData(filename string) (Data, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return Data{}, err
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	chars := make([][]string, 0)
	for _, line := range lines {
		chars = append(chars, strings.Split(line, ""))
	}

	return Data{chars}, nil
}

func task1(data Data) {
	sum := 0
	for i := 1; i < len(data.Chars); i++ {
		for j := range len(data.Chars[0]) {
			if data.Chars[i-1][j] == "S" && data.Chars[i][j] == "." {
				data.Chars[i][j] = "|"
			}
			if data.Chars[i-1][j] == "S" && data.Chars[i][j] == "^" {
				// too lazy to write this case
				panic("S into ^")
			}
			if data.Chars[i-1][j] == "|" {
				if data.Chars[i][j] == "." {
					data.Chars[i][j] = "|"
				}
				if data.Chars[i][j] == "^" {
					sum++
					if j > 0 && data.Chars[i][j-1] == "." {
						data.Chars[i][j-1] = "|"
					}
					if j < len(data.Chars[i])-1 && data.Chars[i][j+1] == "." {
						data.Chars[i][j+1] = "|"
					}
				}
			}
		}
	}

	for i := range data.Chars {
		for j := range data.Chars[i] {
			fmt.Print(data.Chars[i][j])
		}
		fmt.Println()
	}

	fmt.Println(sum)
}

func proceedLine(data Data, row int, col int, cache [][]int) int {
	if col < 0 && col == len(data.Chars[row]) {
		return 0
	}

	if cache[row][col] != -1 {
		return cache[row][col]
	}

	if row == len(data.Chars)-1 {
		return 1
	}

	if data.Chars[row][col] == "." {
		val := proceedLine(data, row+1, col, cache)
		cache[row][col] = val
		return val
	}

	if data.Chars[row][col] == "^" {
		val := proceedLine(data, row+1, col-1, cache) + proceedLine(data, row+1, col+1, cache)
		cache[row][col] = val
		return val
	}

	fmt.Println(data.Chars[row][col])

	panic("unreachable")
}

func task2(data Data) {
	startRow, startCol := -1, -1
startOut:
	for i := range data.Chars {
		for j := range data.Chars[i] {
			if data.Chars[i][j] == "S" {
				startRow = i
				startCol = j
				break startOut
			}
		}
	}

	visited := make([][]int, 0)
	for i := range data.Chars {
		visited = append(visited, make([]int, len(data.Chars[i])))
		for j := range data.Chars[i] {
			visited[i][j] = -1
		}
	}

	sum := proceedLine(data, startRow+1, startCol, visited)
	fmt.Println(sum)
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
