package main

import (
	"AdventOfCode2020/days"
	"AdventOfCode2020/game"
	"fmt"
	"os"
	"strconv"
)

func runDayCLI(day int, input []string) bool {
	solver := days.Get(day).CLI
	output := solver(input)

	for _, line := range output {
		fmt.Println(line)
	}

	return len(output) > 0
}

func cliMain(day int, inputFile string) bool {
	result := true
	fmt.Printf("Running day %d...", day)
	lines := []string{}

	if inputFile != "" {
		lines = days.GetInput(inputFile)
	} else {
		lines = days.GetInputLines(day)
	}

	if day == -1 {
		for dayIdx := 1; dayIdx < days.NumberDays(); dayIdx++ {
			result = result && runDayCLI(day, lines)
		}
	} else {
		result = result && runDayCLI(day, lines)
	}

	return result
}

func visualMain(givenDay int, inputFile string) bool {
	game.Init(givenDay, inputFile)
	return true
}

func main() {
	day := -1
	inputFile := ""

	if len(os.Args) > 1 {
		day, _ = strconv.Atoi(os.Args[1])
	}

	if len(os.Args) > 2 {
		inputFile = os.Args[2]
	}

	result := true

	if os.Getenv("AOC_VISUAL") == "ON" {
		result = visualMain(day, inputFile)
	} else {
		result = cliMain(day, inputFile)
	}

	if !result {
		os.Exit(1)
	}
}
