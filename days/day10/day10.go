package day10

import (
	"fmt"
	"sort"
	"strconv"
)

func dfs(dag map[int][]int, l int, memo map[int]int64) int64 {
	if cached, ok := memo[l]; ok {
		return cached
	}
	if vals, ok := dag[l]; ok {
		var total int64
		for _, val := range vals {
			total += dfs(dag, val, memo)
		}
		memo[l] = total
		return total
	} else {
		return 1
	}
}

func Day10(input []string) []string {
	ints := make([]int, 0, len(input)+1)
	ints = append(ints, 0)
	for _, i := range input {
		intVal, err := strconv.Atoi(i)

		if err != nil {
			panic("Bad input!")
		}
		ints = append(ints, intVal)
	}
	sort.Slice(ints, func(x, y int) bool {
		return ints[x] < ints[y]
	})
	ints = append(ints, ints[len(ints)-1]+3)
	// Go backwards with the shortest strides possible
	num := len(ints) - 1
	last := ints[num]
	dag := make(map[int][]int)
	ones := 0
	threes := 0

	for i := 0; i < len(ints)-1; i++ {
		diff := last - ints[num-1-i]
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
		last = ints[num-1-i]
		connected := make([]int, 0)
		for j := 1; j <= 3 && i+j <= num; j++ {
			if ints[i+j] <= ints[i]+3 {
				connected = append(connected, ints[i+j])
			}
		}
		dag[ints[i]] = connected
	}

	combos := dfs(dag, 0, make(map[int]int64))
	return []string{
		fmt.Sprintf("%d 1's x %d 3's = %d", ones, threes, ones*threes),
		fmt.Sprintf("Number of combos: %d", combos),
	}
}
