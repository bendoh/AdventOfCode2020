package day20

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
  5   2
1 # . . 4
  . # #
3 # # . 6
  5   2
*/

type body [][]bool
type tile struct {
	borders  [8]int
	body     body
	contents body
}
type tileSet map[int]tile

var tiles tileSet
var sideLen int

func parseBorder(line []bool) (int, int) {
	w := len(line)

	var b0, b1 int
	for _, val := range line {
		b0 = b0 << 1
		b1 = b1 >> 1
		if val {
			b0 |= 1
			b1 |= 1 << (w - 1)
		}
	}

	return b0, b1
}

func parseBorders(borders *[8]int, b body) {
	h, w := len(b), len(b[0])
	left := make([]bool, h)
	right := make([]bool, h)
	for i, line := range b {
		if i == 0 {
			borders[0], borders[7] = parseBorder(line)
		}
		if i == h-1 {
			borders[5], borders[2] = parseBorder(line)
		}
		left[i] = line[0]
		right[i] = line[w-1]
	}
	borders[1], borders[6] = parseBorder(right)
	borders[4], borders[3] = parseBorder(left)
}

func parseContent(b body) body {
	h, w := len(b), len(b[0])
	res := make(body, h-2)
	for i := 1; i < h-1; i++ {
		res[i-1] = make([]bool, w-2)
		for j := 1; j < w-1; j++ {
			if b[i][j] {
				res[i-1][j-1] = true
			}
		}
	}
	return res
}

func parseBody(lines []string) body {
	h, w := len(lines), len(lines[0])
	b := make(body, h)

	for i, line := range lines {
		b[i] = make([]bool, w)
		for k := 0; k < w; k++ {
			if line[k] == '#' {
				b[i][k] = true
			}
		}
	}

	return b
}

func MakeTile(lines []string) tile {
	var t tile
	t.body = parseBody(lines)
	t.contents = parseContent(t.body)
	parseBorders(&t.borders, t.body)

	return t
}

func ParseTiles(input []string) {
	var tileNum int

	tiles = make(tileSet)
	tileLines := make([]string, 0)

	for _, line := range input {
		if line == "" {
			if tileNum > 0 {
				tiles[tileNum] = MakeTile(tileLines)
				tileLines = make([]string, 0)
			}
			continue
		}

		if line[0:4] == "Tile" {
			var err error
			tileNum, err = strconv.Atoi(line[5 : len(line)-1])

			if err != nil {
				panic(err)
			}
		} else {
			tileLines = append(tileLines, line)
		}
	}

	t := MakeTile(tileLines)
	tiles[tileNum] = t
}

func Day20(input []string) []string {
	ParseTiles(input)

	/**
	 * Find edges on the tiles in flipped / rotated tiles to find an arrangement where they all line up
	 *
	 * This will be a breadth first search for each tile where each root tile is the top left, and then
	 * each remaining tile should be able to fit in
	 */

	sideFloat := math.Sqrt(float64(len(tiles)))

	if sideFloat-float64(int(sideFloat)) != 0 {
		// Not actually a square!
		panic(fmt.Sprintf("Don't have a square number of tiles: %d", len(tiles)))
	}

	sideLen = int(sideFloat)

	var tileIds []int
	i := 0

	for tileId := range tiles {
		i++
		result := findSquare(tileId, field{})

		if len(result) > 0 {
			for _, c := range result {
				tileIds = append(tileIds, c.tileId)
			}
			break
		}
	}

	if len(tileIds) == 0 {
		return []string{"Couldn't find a square with the given tile set :("}
	}

	result := int64(1)
	factorStrings := make([]string, 4)
	for i, pos := range []int{0, sideLen - 1, len(tileIds) - sideLen, len(tileIds) - 1} {
		result *= int64(tileIds[pos])
		factorStrings[i] = fmt.Sprintf("%d", tileIds[pos])
	}

	return []string{
		fmt.Sprintf("%s=%d", strings.Join(factorStrings, "*"), result),
	}
}

func flip(field [][]bool) [][]bool {
	w := len(field)
	result := make([][]bool, w)

	for i := 0; i < w; i++ {
		result[i] = make([]bool, w)

		for j := 0; j < w; j++ {
			result[i][j] = field[i][w-j-1]
		}
	}

	return result
}

func rotate(t body) body {
	l := len(t[0])
	tn := make(body, l)

	for i := 0; i < l; i++ {
		tn[i] = make([]bool, l)
		for j := 0; j < l; j++ {
			tn[i][j] = t[l-j-1][i]
		}
	}

	return tn
}

func PrintTile(t tile) {
	h := len(t.body)
	var line string
	line += fmt.Sprintf("\n      %-5d%5d\n", t.borders[4], t.borders[1])
	for i := 0; i < h; i++ {
		w := len(t.body[i])
		if i == 0 {
			line += fmt.Sprintf("%5d ", t.borders[0])
		} else if i == h-1 {
			line += fmt.Sprintf("%5d ", t.borders[5])
		} else {
			line += "      "
		}
		for j := 0; j < w; j++ {
			if t.body[i][j] {
				line += fmt.Sprintf("#")
			} else {
				line += fmt.Sprintf(".")
			}
		}
		if i == 0 {
			line += fmt.Sprintf(" %-5d", t.borders[7])
		} else if i == h-1 {
			line += fmt.Sprintf(" %-5d", t.borders[2])
		}
		line += "\n"
	}
	line += fmt.Sprintf("      %-5d%5d\n", t.borders[3], t.borders[6])
	fmt.Print(line)
}

func printField(stack field) {
	fmt.Println()

	for i := 0; i < len(stack); i += sideLen {
		rowTiles := make([]body, sideLen)

		for j := 0; j < sideLen && i+j < len(stack); j++ {
			p := i + j
			rowTiles[j] = transform(tiles[stack[p].tileId].body, stack[p].rot)

			var t tile
			t.body = rowTiles[j]
			t.contents = parseContent(t.body)
			parseBorders(&t.borders, t.body)
			PrintTile(t)
		}

		lines := make([]string, len(tiles[stack[0].tileId].body)+1)

		for t, rowTile := range rowTiles {
			if rowTile == nil {
				continue
			}
			p := i + t
			lines[0] += fmt.Sprintf("%-3d %4d/%d ", p, stack[p].tileId, stack[p].rot)

			for k := 0; k < len(rowTile); k++ {
				for l := 0; l < len(rowTile[k]); l++ {
					if rowTile[k][l] {
						lines[k+1] += "#"
					} else {
						lines[k+1] += "."
					}
				}
				lines[k+1] += " "
			}
		}

		for _, line := range lines {
			fmt.Println(line)
		}
	}
}

// 0-3 => rotated
// 4-7 => rotated + flipped
func transform(t body, kind int) body {
	if kind == 0 {
		return t
	}

	w := len(t)
	if kind > 3 {
		return flip(transform(t, kind-3))
	}

	result := make([][]bool, w)
	result = t

	for i := 1; i <= kind; i++ {
		result = rotate(result)
	}

	return result
}

type cell struct {
	tileId int
	rot    int
}

type field []cell

func findSquare(tileId int, stack field) field {
	for _, c := range stack {
		if c.tileId == tileId {
			return nil
		}
	}

	for rot, border := range tiles[tileId].borders {
		matchesAbove := false
		matchesLeft := false

		pos := len(stack)

		nextStack := append(stack, cell{tileId: tileId, rot: rot})
		printField(nextStack)

		if pos%sideLen == 0 {
			matchesLeft = true
		} else {
			compare := stack[pos-1]

			if tiles[compare.tileId].borders[compare.rot] == border {
				matchesLeft = true
			}
		}

		if pos < sideLen {
			matchesAbove = true
		} else {
			compare := stack[pos-sideLen]

			if tiles[compare.tileId].borders[(compare.rot+1)%8] == border {
				matchesAbove = true
			}
		}

		if matchesLeft && matchesAbove {
			if len(nextStack) == len(tiles) {
				return nextStack
			}

			for nextTileId, _ := range tiles {
				result := findSquare(nextTileId, nextStack)

				if len(result) > 0 {
					return result
				}
			}
		}
	}

	return field{}
}

/**
	0	 0,0=# 0,1=. 0,2=. 1,0=. 1,1=# 1,2=# 2,0=# 2,1=# 2,2=.
	1	 0,0=# 0,1=. 0,2=# 1,0=# 1,1=# 1,2=. 2,0=. 2,1=# 2,2=.

	  j   0 1 2       0 1 2       0 1 2       0 1 2
   	  i 0 # . .     0 # . #     0 . # #     0 . # .
		1 . # #  => 1 # # .  => 1 # # .  => 1 . # #
		2 # # .     2 . # .     2 . . #     2 # . #

	1
		0,0 <= 2,0; 0,1 <= 1,0; 0,2 <= 0,0
		1,0 <= 2,1; 1,1 <= 1,1; 1,2 <= 0,1
		2,0 <= 2,2; 2,1 <= 1,2; 2,2 <= 0,2

	2
		0,0 <= 2,2; 0,1 <= 2,1; 0,2 <= 2,0
		1,0 <= 1,2; 1,1 <= 1,1; 1,2 <= 1,0
		2,0 <= 0,2; 2,1 <= 0,1; 2,2 <= 0,0

	3
		0,0 <= 0,2; 0,1 <= 1,2; 0,2 <= 2,2
		1,0 <= 0,1; 1,1 <= 1,1; 1,2 <= 2,1
		2,0 <= 0,0; 2,1 <= 1,0; 2,2 <= 2,0


			flip h, rotate
	  j   0 1 2       0 1 2       0 1 2       0 1 2
	  i 0 . . #     0 . # .     0 # # .     0 # . #
		1 # # .  => 1 # # .  => 1 . # #  => 1 . # #
		2 . # #     2 # . #     2 # . .     2 . # .
	flip v, rotate (same as flip h and rotate, offset by 2)
	  j   0 1 2       0 1 2       0 1 2       0 1 2
	  i 0 # # .     0 # . #     0 . . #     0 . # .
		1 . # #  => 1 . # #  => 1 # # .  => 1 # # .
		2 # . .     2 . # .     2 . # #     2 # . #
*/
