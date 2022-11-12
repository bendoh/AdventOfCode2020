package day17

import (
	"fmt"
)

type active interface{}

var isActive active

type bound struct {
	min int
	max int
}

func coord(coords []int) string {
	result := ""
	for i := 0; i < len(coords); i++ {
		result += fmt.Sprintf("x%d", coords[i])
	}
	return result[1:]
}

func countNeighbors(field map[string]active, dims []int, offsets []int, l int) int {
	total := 0

	if l == len(dims) {
		numZeros := 0

		next := make([]int, l)

		for i, val := range offsets {
			if val == 0 {
				numZeros++
			}
			next[i] = dims[i] + offsets[i]
		}

		if numZeros == l {
			return 0
		}

		key := coord(next)

		if _, ok := field[key]; ok {
			return 1
		} else {
			return 0
		}
	}

	for i := -1; i <= 1; i++ {
		offsets[l] = i

		total += countNeighbors(field, dims, offsets, l+1)
	}

	return total
}

func shouldBecomeActive(field map[string]active, pos []int) bool {
	numNeighbors := countNeighbors(field, pos, make([]int, len(pos)), 0)

	_, cellActive := field[coord(pos)]

	if cellActive && (numNeighbors == 2 || numNeighbors == 3) {
		return true
	} else if !cellActive && numNeighbors == 3 {
		return true
	}

	return false
}

func stepRecursive(field map[string]active, result *map[string]active, boundary []bound, boundaryResult *[]bound, pos []int, dim int) int {
	if dim == len(boundary) {
		if shouldBecomeActive(field, pos) {
			(*result)[coord(pos)] = true

			for i := 0; i < len(pos); i++ {
				if pos[i] < boundary[i].min {
					(*boundaryResult)[i].min = pos[i]
				}
				if pos[i] > boundary[i].max {
					(*boundaryResult)[i].max = pos[i]
				}
			}

			return 1
		} else {
			return 0
		}
	} else {
		numActive := 0

		for i := boundary[dim].min - 1; i <= boundary[dim].max+1; i++ {
			pos[dim] = i
			numActive += stepRecursive(field, result, boundary, boundaryResult, pos, dim+1)
		}

		return numActive
	}
}

func printField(field map[string]active, bounds []bound, prefix []int, depth int) {
	dimension := len(bounds) - depth - 1

	if dimension >= 2 {
		for i := bounds[dimension].min; i <= bounds[dimension].max; i++ {
			printField(field, bounds, append(prefix, i), depth+1)
		}
		return
	}

	prefixDimensions := []rune{'z', 'w'}
	for i := 0; i < len(prefix); i++ {
		fmt.Printf("%c=%d ", prefixDimensions[i], prefix[i])
	}
	fmt.Println("")
	for x := bounds[0].min; x <= bounds[0].max; x++ {
		for y := bounds[1].min; y <= bounds[1].max; y++ {
			pos := []int{x, y}
			_, cellActive := field[coord(append(pos, prefix...))]

			if cellActive {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}

}

type stepResult struct {
	field  map[string]active
	bounds []bound
	active int
}

func stepDimensions(field map[string]active, bounds []bound) stepResult {
	nextField := make(map[string]active)
	nextBounds := make([]bound, len(bounds))

	copy(nextBounds, bounds)

	numActive := stepRecursive(field, &nextField, bounds, &nextBounds, make([]int, len(bounds)), 0)

	return stepResult{
		field:  nextField,
		bounds: nextBounds,
		active: numActive,
	}
}

func Day17(input []string) []string {
	field3 := make(map[string]active)
	field4 := make(map[string]active)

	bounds3 := make([]bound, 3)
	bounds4 := make([]bound, 4)

	bounds3[0].max = len(input) - 1
	bounds3[1].max = len(input[0]) - 1

	bounds4[0].max = bounds3[0].max
	bounds4[1].max = bounds3[1].max

	for row, line := range input {
		for col := 0; col < len(line); col++ {
			if line[col] == '#' {
				field3[coord([]int{row, col, 0})] = isActive
				field4[coord([]int{row, col, 0, 0})] = isActive
			}
		}
	}

	result := make([]string, 0)
	fmt.Println("")
	//	printField(field3, bounds3, []int{}, 0)
	printField(field4, bounds4, []int{}, 0)
	for i := 0; i < 6; i++ {
		stepResult3 := stepDimensions(field3, bounds3)
		step3 := fmt.Sprintf("[3d] After %d steps, %d active", i+1, stepResult3.active)
		result = append(result, step3)
		field3 = stepResult3.field
		bounds3 = stepResult3.bounds

		stepResult4 := stepDimensions(field4, bounds4)
		step4 := fmt.Sprintf("[4d] After %d steps, %d active", i+1, stepResult4.active)
		result = append(result, step4)
		field4 = stepResult4.field
		bounds4 = stepResult4.bounds

		//		printField(field3, bounds3, []int{}, 0)
		printField(field4, bounds4, []int{}, 0)
	}
	return result
}
