package day9

import (
	"fmt"
	"math"
	"strconv"
)

func inPreamble(pre []int, num int) bool {
	for i := 0; i < len(pre)-1; i++ {
		if pre[i] > num {
			continue
		}
		for j := i + 1; j < len(pre); j++ {
			if pre[i]+pre[j] == num {
				return true
			}
		}
	}

	return false
}

func findRange(num int, numbers []int) (int, int) {
	for i := range numbers {
		try := num
		for j := i; j < len(numbers); j++ {
			try -= numbers[j]

			if try == 0 {
				return i, j
			}
			if try < 0 {
				break
			}
		}
	}

	return -1, -1
}
func Day9(input []string) []string {
	numbers := make([]int, 0, len(input))
	pl := 25

	if input[0][0] == 'p' {
		if _pl, err := strconv.Atoi(input[0][1:]); err == nil {
			pl = _pl
		}
		input = input[1:]
	}

	for _, num := range input {
		if i, err := strconv.Atoi(num); err == nil {
			numbers = append(numbers, i)
		}
	}

	firstInvalid := 0

	for i := pl; i < len(numbers); i++ {
		num := numbers[i]
		if !inPreamble(numbers[i-pl:i], num) {
			firstInvalid = num
			break
		}
	}

	first, last := findRange(firstInvalid, numbers)
	min := math.MaxInt
	max := 0

	for i := first; i <= last; i++ {
		if numbers[i] < min {
			min = numbers[i]
		}

		if numbers[i] > max {
			max = numbers[i]
		}
	}

	return []string{
		fmt.Sprintf("First invalid number: %d", firstInvalid),
		fmt.Sprintf("range=numbers[%d:%d]; minmax(range) = %d, %d sum = %d", first, last, min, max, min+max),
	}
}
