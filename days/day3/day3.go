package day3

import "fmt"

func Day3(input []string) []string {
	return append(part1(input), part2(input)...)
}

func part1(input []string) []string {
	count := 0
	slope := 3

	for line := 0; line < len(input); line++ {
		c := rune(input[line][line*slope%len(input[line])])
		if c == '#' {
			count++
		}
	}

	result := fmt.Sprintf("Trees hit with slope 3,1: %d", count)

	return []string{result}
}

func part2(input []string) []string {
	slopes := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}
	var results []int

	for _, slope := range slopes {
		count := 0

		right, down := slope[0], slope[1]

		for line := 0; line < len(input); line += down {
			idx := right * (line / down)

			c := rune(input[line][idx%len(input[line])])
			if c == '#' {
				count++
			}
		}
		results = append(results, count)
	}

	total := 1

	output := []string{}
	for idx, result := range results {
		output = append(output,
			fmt.Sprintf("Trees hit at slope %d,%d: %d", slopes[idx][0], slopes[idx][1], result))
		total *= result
	}

	output = append(output, fmt.Sprintf("Product of trees hit: %d", total))

	return output
}
