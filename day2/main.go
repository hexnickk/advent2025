package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData(filename string) [][]int {
	dat, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(dat)), ",")
	data := make([][]int, 0)

	for i := 0; i < len(lines); i++ {
		line := strings.Split(strings.TrimSpace(lines[i]), "-")
		start, _ := strconv.Atoi(line[0])
		end, _ := strconv.Atoi(line[1])
		data = append(data, []int{start, end})
	}

	return data
}

func isInvalidT1(num int) bool {
	str := strconv.Itoa(num)

	if len(str)%2 != 0 {
		return false
	}

	mid := len(str) / 2
	for i := range mid {
		if str[i] != str[mid+i] {
			return false
		}
	}

	return true
}

func task1(data [][]int) {
	sum := 0
	for entry := range data {
		fmt.Println(data[entry][0], "->", data[entry][1])
		for i := data[entry][0]; i <= data[entry][1]; i++ {
			if isInvalidT1(i) {
				sum += i
				fmt.Println(i)
			}
		}
	}
	fmt.Println("---")
	fmt.Println("Sum ->", sum)
}

func isInvalidN(num string, n int) bool {
	if len(num)%n != 0 {
		return false
	}
	for offset := range n {
		for j := 1; j < len(num)/n; j++ {
			if num[offset] != num[offset+j*n] {
				return false
			}
		}
	}
	return true
}

func isInvalidT2(num int) bool {
	str := strconv.Itoa(num)
	for n := 1; n < len(str)/2+1; n++ {
		if isInvalidN(str, n) {
			return true
		}
	}
	return false
}

func task2(data [][]int) {
	sum := 0
	for entry := range data {
		fmt.Println(data[entry][0], "->", data[entry][1])
		for i := data[entry][0]; i <= data[entry][1]; i++ {
			if isInvalidT2(i) {
				sum += i
				fmt.Println(i)
			}
		}
	}
	fmt.Println("---")
	fmt.Println("Sum ->", sum)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <task name> <input file>")
		return
	}

	taskname := os.Args[1]
	filename := os.Args[2]

	data := getData(filename)

	if taskname == "1" {
		task1(data)
	}
	if taskname == "2" {
		task2(data)
	}
}
