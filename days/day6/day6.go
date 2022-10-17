package day6

import "fmt"

func Day6(input []string) []string {
	type answerSet map[rune]int

	var numPassengers []int

	answers := make([]answerSet, 0)

	input = append(input, "")
	group := make(answerSet)

	groupSize := 0
	for _, line := range input {
		if line == "" {
			numPassengers = append(numPassengers, groupSize)
			answers = append(answers, group)
			group = make(answerSet)
			groupSize = 0
			continue
		}

		for j := 0; j < len(line); j++ {
			group[rune(line[j])] += 1
		}
		groupSize++
	}

	totalAny := 0
	totalEvery := 0
	for i, ans := range answers {
		totalAny += len(ans)

		for _, cnt := range ans {
			if cnt == numPassengers[i] {
				totalEvery++
			}
		}
	}
	return []string{fmt.Sprintf("Any answers=%d, Every Answer=%d", totalAny, totalEvery)}
}
