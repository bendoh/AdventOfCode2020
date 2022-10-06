package day1

import (
	"AdventOfCode2020/util"
	"fmt"
)

func Day1(input []string) bool {
	numbers := util.ParseIntList(input)

	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				fmt.Printf("2020 = %d + %d, %d x %d = %d\n", numbers[i], numbers[j], numbers[i], numbers[j], numbers[i]*numbers[j])
				return true
			}
		}
	}

	return false
}
