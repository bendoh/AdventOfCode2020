package main

import (
	"AdventOfCode2020/day0"
	"AdventOfCode2020/day1"
	"bufio"
	"os"
	"strconv"
)

func main() {
	day := 0
	if len(os.Args) > 1 {
		day, _ = strconv.Atoi(os.Args[1])
	}
	inputFile := "/dev/stdin"
	if len(os.Args) > 2 {
		inputFile = os.Args[2]
	}

	input, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(input)
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if len(line) > 0 {
			lines = append(lines, line[:len(line)-1])
		} else {
			break
		}
	}

	days := []func([]string) bool{day0.Day0, day1.Day1}
	days[day](lines)

	os.Exit(0)
}
