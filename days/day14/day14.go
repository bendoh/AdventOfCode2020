package day14

import (
	"fmt"
	"regexp"
	"strconv"
)

func applyValueMask(mask string, val int64) int64 {
	l := len(mask) - 1
	for i := 0; i <= l; i++ {
		switch mask[l-i] {
		case '0':
			val = val &^ (1 << i)
		case '1':
			val = val | (1 << i)
		}
	}

	return val
}

func maskAddr(mask string, addr int64) []int64 {
	l := len(mask) - 1
	floaters := make([]int, 0)
	baseAddr := addr
	for i := 0; i <= l; i++ {
		switch mask[l-i] {
		case '1':
			baseAddr = baseAddr | (1 << i)
		case 'X':
			floaters = append(floaters, i)
		}
	}

	addrs := []int64{baseAddr}

	for _, i := range floaters {
		nextSet := make([]int64, 0)

		for _, na := range addrs {
			nextSet = append(nextSet, na|(1<<i), na&^(1<<i))
		}

		addrs = nextSet
	}

	return addrs
}

// 1898179721949
// 11140949342941
// 11179633149677
func Day14(input []string) []string {
	var mask string
	reg, err := regexp.Compile("(mask|mem\\[(\\d+)\\]) = (.*)$")

	mem1 := make(map[int64]int64)
	mem2 := make(map[int64]int64)

	if err != nil {
		return []string{"Error!"}
	}

	fmt.Println("")
	for _, line := range input {
		matches := reg.FindStringSubmatch(line)

		if matches == nil {
			continue
		}

		if matches[1] == "mask" {
			mask = matches[3]
			fmt.Printf("m%s\n", mask)
		} else {
			addr, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				return []string{"Error"}
			}
			intVal, err := strconv.ParseInt(matches[3], 10, 64)
			if err != nil {
				return []string{"Error"}
			}
			mem1[addr] = applyValueMask(mask, intVal)
			fmt.Printf("%037b = mem1[%d]\n", mem1[addr], addr)

			for _, maskedAddr := range maskAddr(mask, addr) {
				mem2[maskedAddr] = intVal
			}
		}
	}

	var total1 int64
	for _, val := range mem1 {
		total1 += val
	}
	var total2 int64
	for _, val := range mem2 {
		total2 += val
	}

	return []string{
		fmt.Sprintf("1: %d => %064b", total1, total1),
		fmt.Sprintf("2: %d => %064b", total2, total2),
	}
}
