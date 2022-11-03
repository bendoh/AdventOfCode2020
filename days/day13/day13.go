package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day13(input []string) []string {
	departure, err := strconv.ParseInt(input[0], 10, 64)
	schedule := strings.Split(input[1], ",")

	if err != nil {
		return []string{}
	}

	buses := make([]int64, 0, len(schedule))
	for _, busNum := range schedule {
		if busNum != "x" {
			busNumInt, err := strconv.ParseInt(busNum, 10, 64)
			if err != nil {
				return []string{}
			}
			buses = append(buses, busNumInt)
		} else {
			buses = append(buses, 0)
		}
	}

	diff := int64(math.MaxInt64)
	winner := int64(0)

	for _, busNum := range buses {
		if busNum == 0 {
			continue
		}
		d := busNum*(departure/busNum+1) - departure

		if d < diff {
			diff = d
			winner = busNum
		}
	}

	d := int64(1)
	ts := int64(0)
	for idx, bus := range buses {
		if bus == 0 {
			continue
		}
		for {
			ts += d
			if (ts+int64(idx))%bus == 0 {
				d = d * bus
				break
			}
		}
	}

	return []string{
		fmt.Sprintf("Diff=%d * Winner=%d = %d", diff, winner, diff*winner),
		fmt.Sprintf("Earliest sequential ts=%d", ts),
	}
}
