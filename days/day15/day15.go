package day15

import (
	"fmt"
	"strconv"
	"strings"
)

func compute(numbers []int64, nrounds int64) int64 {
	lastSeen := make(map[int64]int64)
	l := len(numbers)
	for i := 1; i < l; i++ {
		lastSeen[numbers[i-1]] = int64(i)
	}
	last := numbers[l-1]

	for i := int64(l); i < nrounds; i++ {
		before, seen := lastSeen[last]
		lastSeen[last] = i
		if seen {
			last = i - before
		} else {
			last = 0
		}
	}

	return last
}
func Day15(input []string) []string {
	result := make([]string, 0, len(input))

	part1steps := 2020
	part2steps := 30000000

	for caseNum, line := range input {
		nums := strings.Split(line, ",")
		numbers := make([]int64, len(nums))

		for i, num := range nums {
			var err error
			numbers[i], err = strconv.ParseInt(num, 10, 64)

			if err != nil {
				return []string{"Error!"}
			}
		}

		part1 := compute(numbers, int64(part1steps))
		part2 := compute(numbers, int64(part2steps))

		result = append(result,
			fmt.Sprintf(
				"Case %d: compute(%s, %d)=%d compute(%s, %d)=%d", caseNum,
				line, part1steps, part1,
				line, part2steps, part2))
	}

	return result
}
