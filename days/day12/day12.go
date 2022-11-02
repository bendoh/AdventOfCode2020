package day12

import (
	"fmt"
	"math"
)

func Day12(input []string) []string {
	return []string{
		part1(input),
		part2(input),
	}
}

func part1(input []string) string {
	pos := []int{0, 0}
	angle := 0

	for _, line := range input {
		var op rune
		var val int
		nparsed, err := fmt.Sscanf(line, "%c%d", &op, &val)

		if err != nil {
			return "Error!"
		}

		if nparsed > 0 {
			switch op {
			case 'N':
				pos[1] += val
			case 'S':
				pos[1] -= val
			case 'E':
				pos[0] += val
			case 'W':
				pos[0] -= val
			case 'F':
				var stepX, stepY int
				switch ((angle % 360) + 360) % 360 {
				case 0:
					stepX = 1
				case 180:
					stepX = -1
				case 90:
					stepY = 1
				case 270:
					stepY = -1
				}
				pos[0] += stepX * val
				pos[1] += stepY * val
			case 'L':
				angle += val
			case 'R':
				angle -= val
			}

		}
	}

	return fmt.Sprintf(
		"Distance: %.0f + %.0f = %.0f",
		math.Abs(float64(pos[0])), math.Abs(float64(pos[1])),
		math.Abs(float64(pos[0]))+math.Abs(float64(pos[1])),
	)
}

func part2(input []string) string {
	pos := []int{0, 0}
	wp := []int{10, 1}

	for _, line := range input {
		var op rune
		var val int
		nparsed, err := fmt.Sscanf(line, "%c%d", &op, &val)

		if err != nil {
			return "Error!"
		}

		pre := []int{wp[0], wp[1]}
		if nparsed > 0 {
			switch op {
			case 'E':
				wp[0] += val
			case 'W':
				wp[0] -= val
			case 'N':
				wp[1] += val
			case 'S':
				wp[1] -= val
			case 'F':
				pos[0] += wp[0] * val
				pos[1] += wp[1] * val
			case 'L':
				switch val % 360 {
				case 90:
					wp[0] = -pre[1]
					wp[1] = pre[0]
				case 180:
					wp[0] = -pre[0]
					wp[1] = -pre[1]
				case 270:
					wp[0] = pre[1]
					wp[1] = -pre[0]
				}
			case 'R':
				switch val % 360 {
				case 90:
					wp[0] = pre[1]
					wp[1] = -pre[0]
				case 180:
					wp[0] = -pre[0]
					wp[1] = -pre[1]
				case 270:
					wp[0] = -pre[1]
					wp[1] = pre[0]
				}

			}

		}
	}

	return fmt.Sprintf(
		"Distance: %.0f + %.0f = %.0f",
		math.Abs(float64(pos[0])), math.Abs(float64(pos[1])),
		math.Abs(float64(pos[0]))+math.Abs(float64(pos[1])),
	)

}
