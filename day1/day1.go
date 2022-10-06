package day1

import (
	"AdventOfCode2020/util"
	"fmt"
)

func part1(numbers []int) bool {
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
func part2(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j < len(numbers); j++ {
			remainder := 2020 - (numbers[i] + numbers[j])

			if remainder > 0 && remainder < 2020 {
				for k := j; k < len(numbers); k++ {
					if numbers[k] == remainder {
						fmt.Printf("2020 = %d + %d + %d, %d * %d * %d = %d",
							numbers[i], numbers[j], numbers[k],
							numbers[i], numbers[j], numbers[k],
							numbers[i]*numbers[j]*numbers[k],
						)
						return true
					}
				}
			}
		}
	}

	return false
}
func Day1(input []string) bool {
	numbers := util.ParseIntList(input)

	return part1(numbers) && part2(numbers)
}
