package day5

import "fmt"

func pow(x, y int) int {
	if x == 0 {
		return 0
	}
	if y == 0 {
		return 1
	}
	res := x
	for i := 1; i < y; i++ {
		res = res * x
	}
	return res
}

type empty struct{}

var exists empty

func Day5(input []string) []string {
	// FBFBBFF -> 0101100 -> 12 + 32 == 42
	// RLR -> 101 -> 4 + 1 = 5
	var seatIds []string
	seats := make(map[int]interface{})
	highest := 0
	for _, line := range input {
		seatId := 0
		for i := 0; i < len(line); i++ {
			pos := len(line) - i - 1

			if line[pos] == 'R' || line[pos] == 'B' {
				seatId += pow(2, i)
			}
		}
		if seatId > highest {
			highest = seatId
		}
		seats[seatId] = exists
	}

	mySeatId := 0

	for i := 0; i < highest; i++ {
		if _, ok := seats[i]; !ok {
			_, beforeExists := seats[i-1]
			_, afterExists := seats[i-1]

			if beforeExists && afterExists {
				mySeatId = i
			}
		}

	}
	return append(seatIds, fmt.Sprintf("Highest: %d, My Seat ID=%d", highest, mySeatId))
}
