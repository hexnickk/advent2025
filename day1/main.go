package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getData(filename string) []int {
	dat, _ := os.ReadFile(filename)
	lines := strings.Split(string(dat), "\n")
	data := make([]int, 0)

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		if strings.HasPrefix(lines[i], "L") {
			num, _ := strconv.Atoi(lines[i][1:])
			data = append(data, -num)
		}
		if strings.HasPrefix(lines[i], "R") {
			num, _ := strconv.Atoi(lines[i][1:])
			data = append(data, num)
		}
	}

	return data
}

func task1(data []int) {
	value := 50
	for _, rotation := range data {
		value = (value + rotation + 100) % 100
	}

	fmt.Println("Final value:", value)
}

func task2(data []int) {
	value := 50
	pass := 0
	for _, rotation := range data {
		// obvious passes through 0
		pass += int(math.Abs(float64(rotation) / 100))

		newValue := value + rotation%100
		if newValue < 0 && value > 0 || newValue > 100 || newValue%100 == 0 {
			pass++
		}

		value = (newValue + 100) % 100
		fmt.Println("rotation:", rotation, "\tvalue:", value, "\t pass:", pass)
	}

	fmt.Println("----- FINAL -----")
	fmt.Println("value:", value, "\t pass:", pass)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input file> <task name>")
		return
	}

	filename := os.Args[1]
	taskname := os.Args[2]

	data := getData(filename)

	if taskname == "1" {
		task1(data)
	}
	if taskname == "2" {
		task2(data)
	}
}
