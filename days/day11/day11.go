package day11

import "fmt"

func adjacentStep(layout [][]rune) ([][]rune, int, int) {
	next := make([][]rune, len(layout))
	numChanges, numOccupied := 0, 0

	fmt.Println("")
	for i, line := range layout {
		next[i] = make([]rune, len(line))
		lineChanges := 0

		for j, state := range line {
			adjOccupied := 0
			next[i][j] = layout[i][j]

			if state == '.' {
				continue
			}

			for rowOffset := -1; rowOffset <= 1; rowOffset++ {
				aRow := i + rowOffset

				for colOffset := -1; colOffset <= 1; colOffset++ {
					aCol := j + colOffset

					if aRow == i && aCol == j {
						continue
					}

					if aRow >= 0 && aRow < len(layout) &&
						aCol >= 0 && aCol < len(line) &&
						layout[aRow][aCol] == '#' {
						adjOccupied++
					}
				}
			}

			if state == 'L' && adjOccupied == 0 {
				next[i][j] = '#'
				lineChanges++
			} else if state == '#' && adjOccupied >= 4 {
				next[i][j] = 'L'
				lineChanges++
			}

			if next[i][j] == '#' {
				numOccupied++
			}
		}
		numChanges += lineChanges

		fmt.Printf("%s => %s (+%d -> %d)\n", string(line), string(next[i]), lineChanges, numOccupied)
	}

	return next, numChanges, numOccupied
}

func nextInDir(layout [][]rune, i int, j int, rowStep int, colStep int) rune {
	for {
		if i < 0 || i >= len(layout) || j < 0 || j >= len(layout[i]) {
			return '.'
		}
		char := layout[i][j]
		if char == '#' || char == 'L' {
			return char
		}
		i += rowStep
		j += colStep
	}
}
func lineOfSightStep(layout [][]rune) ([][]rune, int, int) {
	next := make([][]rune, len(layout))
	numChanges, numOccupied := 0, 0

	fmt.Println("")
	for i, line := range layout {
		next[i] = make([]rune, len(line))
		lineChanges := 0

		for j, state := range line {
			adjOccupied := 0
			next[i][j] = layout[i][j]

			if state == '.' {
				continue
			}

			for rowOffset := -1; rowOffset <= 1; rowOffset++ {
				aRow := i + rowOffset

				for colOffset := -1; colOffset <= 1; colOffset++ {
					aCol := j + colOffset

					if aRow == i && aCol == j {
						continue
					}

					adjacent := nextInDir(layout, i+rowOffset, j+colOffset, rowOffset, colOffset)

					if adjacent == '#' {
						adjOccupied++
					}
				}
			}

			if state == 'L' && adjOccupied == 0 {
				next[i][j] = '#'
				lineChanges++
			} else if state == '#' && adjOccupied >= 5 {
				next[i][j] = 'L'
				lineChanges++
			}

			if next[i][j] == '#' {
				numOccupied++
			}
		}
		numChanges += lineChanges

		fmt.Printf("%s => %s (+%d -> %d)\n", string(line), string(next[i]), lineChanges, numOccupied)
	}

	return next, numChanges, numOccupied
}

func part1(layout [][]rune) string {
	next, numChanges, numOccupied := adjacentStep(layout)
	rounds := 0

	for numChanges > 0 {
		layout = next
		next, numChanges, numOccupied = adjacentStep(layout)
		rounds++
	}
	return fmt.Sprintf("Took %d rounds and ended up with %d occupied seats", rounds, numOccupied)
}

func part2(layout [][]rune) string {
	next, numChanges, numOccupied := lineOfSightStep(layout)
	rounds := 0

	for numChanges > 0 {
		layout = next
		next, numChanges, numOccupied = lineOfSightStep(layout)
		rounds++
	}
	return fmt.Sprintf("Took %d rounds and ended up with %d occupied seats", rounds, numOccupied)
}
func Day11(input []string) []string {
	layout := make([][]rune, len(input))
	for i, line := range input {
		layout[i] = make([]rune, len(line))

		for j, pos := range line {
			layout[i][j] = pos
		}
	}

	part1result := part1(layout)
	part2result := part2(layout)

	return []string{part1result, part2result}
}
