package day1

import (
	"AdventOfCode2020/util"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"strings"
)

var numbers []int

var part1done bool
var part2done bool

var finalResults []string

func VisualSetup(_input []string) {
	numbers = util.ParseIntList(_input)
	part1indices = []int{0, 1}
	part1done, part2done = false, false
	finalResults = []string{}
}

func VisualStep(screen *ebiten.Image, timeElapsed int64) ([]string, bool) {
	var results []string

	if !part1done {
		results = append(results, part1step(screen, timeElapsed)...)

		if part1done {
			finalResults = append(finalResults, results...)
		}
	} else if part1done && !part2done {
		results = append(results, part2step(screen, timeElapsed)...)

		if part2done {
			finalResults = append(finalResults, results...)
		}
	} else {
		ebitenutil.DebugPrint(screen, strings.Join(finalResults, "\n"))
	}

	if part1done && part2done {
		return finalResults, true
	} else {
		return results, false
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var part1indices []int

var lastStep int64

func part1step(screen *ebiten.Image, nsElapsed int64) []string {
	// Speed = 50 -> 50 steps per second -> 1/50 seconds / step

	maxSteps := nsElapsed / 1e5

	last := len(numbers)

	for part1indices[0] < last {
		i, j := part1indices[0], part1indices[1]
		steps := int64(i*len(numbers) + j)
		stepString := fmt.Sprintf("Steps: %d / %d\n", steps, maxSteps)
		fmt.Print(stepString)
		if steps > maxSteps {
			ebitenutil.DebugPrint(screen, stepString)
			return []string{}
		}

		if numbers[i]+numbers[j] == 2020 {
			part1done = true
			result := fmt.Sprintf("2020 = %d + %d, %d * %d = %d",
				numbers[i], numbers[j],
				numbers[i], numbers[j],
				numbers[i]*numbers[j],
			)
			ebitenutil.DebugPrint(screen, result)
			return []string{result}
		}

		part1indices[1]++

		if part1indices[1] == last {
			part1indices[0]++
			part1indices[1] = part1indices[0] + 1
		}
	}

	return []string{}
}

func part2step(screen *ebiten.Image, timeElapsed int64) []string {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j < len(numbers); j++ {
			remainder := 2020 - (numbers[i] + numbers[j])

			if remainder > 0 && remainder < 2020 {
				for k := j; k < len(numbers); k++ {
					if numbers[k] == remainder {
						part2done = true
						result := fmt.Sprintf("2020 = %d + %d + %d, %d * %d * %d = %d",
							numbers[i], numbers[j], numbers[k],
							numbers[i], numbers[j], numbers[k],
							numbers[i]*numbers[j]*numbers[k],
						)
						ebitenutil.DebugPrint(screen, result)
						return []string{result}
					}
				}
			}
		}
	}

	return []string{}
}
