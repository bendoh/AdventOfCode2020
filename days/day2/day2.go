package day2

import "fmt"

func Day2(input []string) []string {
	return append(part1(input), part2(input)...)
}

func part1(input []string) []string {
	numValid := 0

	for lineNum := 0; lineNum < len(input); lineNum++ {
		line := input[lineNum]

		var min, max int
		var char uint8
		var password string

		fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &char, &password)

		count := 0

		for i := 0; i < len(password); i++ {
			if password[i] == char {
				count++
			}
		}

		if count >= min && count <= max {
			numValid++
		}
	}

	result := fmt.Sprintf("Valid passwords: %d", numValid)
	return []string{result}
}

func part2(input []string) []string {
	numValid := 0

	for lineNum := 0; lineNum < len(input); lineNum++ {
		line := input[lineNum]

		var min, max int
		var char uint8
		var password string

		fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &char, &password)

		count := 0

		if password[min-1] == char {
			count++
		}

		if password[max-1] == char {
			count++
		}

		if count == 1 {
			numValid++
		}
	}

	result := fmt.Sprintf("Valid passwords: %d", numValid)
	return []string{result}
}
