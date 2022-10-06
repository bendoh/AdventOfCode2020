package util

import "strconv"

func ParseIntList(input []string) []int {
	var numbers []int

	for _, line := range input {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	return numbers
}
