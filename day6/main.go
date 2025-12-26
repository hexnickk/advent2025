package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Numbers    [][]int
	Operations []rune
}

func getChunks(line string) []string {
	parts := strings.Split(line, " ")
	chunks := make([]string, 0)
	for _, part := range parts {
		if part != "" {
			chunks = append(chunks, part)
		}
	}
	return chunks
}

func getData(filename string) (Data, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return Data{}, err
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	numbers := make([][]int, 0)
	for i := 0; i < len(lines)-1; i++ {
		chunks := getChunks(lines[i])
		row := make([]int, 0)
		for _, chunk := range chunks {
			num, _ := strconv.Atoi(chunk)
			row = append(row, num)
		}
		numbers = append(numbers, row)
	}

	operations := make([]rune, 0)
	chunks := getChunks(lines[len(lines)-1])
	for _, chunk := range chunks {
		operations = append(operations, []rune(chunk)[0])
	}

	return Data{numbers, operations}, nil
}

func getOp(char rune) (func(a int, b int) int, int) {
	switch char {
	case '+':
		return func(a int, b int) int {
			return a + b
		}, 0
	case '*':
		return func(a int, b int) int {
			return a * b
		}, 1
	default:
		panic("Unknown operation")
	}
}

func task1(data Data) {
	sum := 0
	for i := range len(data.Numbers[0]) {
		operator, def := getOp(data.Operations[i])
		result := def
		for j := range len(data.Numbers) {
			result = operator(data.Numbers[j][i], result)
		}
		sum += result
	}
	fmt.Println(sum)
}

func getData2(filename string) (Data, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return Data{}, err
	}
	rawLines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	rawCharsLines := make([][]string, 0)
	for _, line := range rawLines {
		rawCharsLines = append(rawCharsLines, strings.Split(line, ""))
	}

	maxLen := 0
	for _, line := range rawCharsLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	numbers := make([][]int, 0)
	operators := make([]rune, 0)
	numbersCol := make([]int, 0)

	for col := maxLen - 1; col >= 0; col-- {
		numStr := ""
		for row := 0; row < len(rawCharsLines)-1; row++ {
			if len(rawCharsLines[row]) > col && rawCharsLines[row][col] != " " {
				numStr += rawCharsLines[row][col]
			}
		}
		if numStr != "" {
			num, _ := strconv.Atoi(numStr)
			numStr = ""
			numbersCol = append(numbersCol, num)
		}
		if len(rawCharsLines[len(rawCharsLines)-1]) > col && rawCharsLines[len(rawCharsLines)-1][col] != " " {
			operators = append(operators, []rune(rawCharsLines[len(rawCharsLines)-1][col])[0])
			numbers = append(numbers, numbersCol)
			numbersCol = make([]int, 0)
		}
	}

	// fmt.Println(numbers)
	// for _, operator := range operators {
	// 	fmt.Print(string(operator), " ")
	// }
	// fmt.Println()

	return Data{numbers, operators}, nil
}

func task2(data Data) {
	sum := 0
	for i := range len(data.Numbers) {
		operator, def := getOp(data.Operations[i])
		result := def
		for j := range len(data.Numbers[i]) {
			result = operator(data.Numbers[i][j], result)
		}
		sum += result
	}
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

	data2, err := getData2(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading data:", err)
		os.Exit(1)
	}
	if taskname == "2" {
		task2(data2)
	}
}
