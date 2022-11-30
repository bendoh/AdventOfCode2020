package day22

import (
	"fmt"
	"strconv"
	"strings"
)

type empty interface{}

var isEmpty empty

var p1, p2 []int

func ParseInput(input []string) {
	p1 = make([]int, 0)
	p2 = make([]int, 0)

	var ref = &p1

	for _, line := range input {
		if line == "Player 1:" {
			continue
		}
		if line == "Player 2:" {
			ref = &p2
		}
		if intVal, err := strconv.Atoi(line); err == nil {
			*ref = append(*ref, intVal)
		}
	}
}

func Play() (int, []int) {
	rounds := 0
	for {
		rounds++
		if len(p1) == 0 {
			return rounds, p2
		} else if len(p2) == 0 {
			return rounds, p1
		}

		if p1[0] > p2[0] {
			p1 = append(p1[1:], p1[0], p2[0])
			p2 = p2[1:]
		} else {
			p2 = append(p2[1:], p2[0], p1[0])
			p1 = p1[1:]
		}
	}
}

func getHash(p []int) string {
	res := make([]string, len(p))
	for i, v := range p {
		res[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(res, ",")
}

func PlayRecursive(ip1, ip2 []int) (int, int, []int) {
	rounds := 0

	p1, p2 := make([]int, len(ip1)), make([]int, len(ip2))

	copy(p1, ip1)
	copy(p2, ip2)

	seen := make(map[string]empty)

	for {
		if len(p1) == 0 {
			return rounds, 2, p2
		} else if len(p2) == 0 {
			return rounds, 1, p1
		}

		rounds++
		_, seen1 := seen["1:"+getHash(p1)]
		_, seen2 := seen["2:"+getHash(p2)]

		seenConfig := seen1 && seen2

		var winner int
		if seenConfig {
			return 0, 1, []int{}
		} else if p1[0] < len(p1) && p2[0] < len(p2) {
			_, winner, _ = PlayRecursive(p1[1:p1[0]+1], p2[1:p2[0]+1])
		} else if p1[0] > p2[0] {
			winner = 1
		} else {
			winner = 2
		}

		seen["1:"+getHash(p1)] = isEmpty
		seen["2:"+getHash(p2)] = isEmpty

		if winner == 1 {
			p1 = append(p1[1:], p1[0], p2[0])
			p2 = p2[1:]
		} else if winner == 2 {
			p2 = append(p2[1:], p2[0], p1[0])
			p1 = p1[1:]
		}
	}
}

func getResult(w []int) int {
	res := 0
	for i, val := range w {
		res += val * (len(w) - i)
	}
	return res
}

func Day22(input []string) []string {
	ParseInput(input)

	rp1 := make([]int, len(p1))
	rp2 := make([]int, len(p2))

	copy(rp1, p1)
	copy(rp2, p2)

	nRounds1, winner1 := Play()
	result1 := getResult(winner1)

	nRounds2, _, deck := PlayRecursive(rp1, rp2)

	result2 := getResult(deck)

	return []string{
		fmt.Sprintf("Result after %d rounds: %d", nRounds1, result1),
		fmt.Sprintf("Result after %d recursive rounds: %d", nRounds2, result2),
	}
}
