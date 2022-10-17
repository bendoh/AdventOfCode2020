package day8

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type instruction struct {
	instr string
	arg   int
}

func Day8(input []string) []string {
	reg := regexp.MustCompile("(acc|jmp|nop) ([+-]?\\d+)$")
	instructions := make([]instruction, 0, len(input))

	for _, line := range input {
		matches := reg.FindStringSubmatch(line)

		if len(matches) == 0 {
			panic("Bad command")
		}

		val, err := strconv.Atoi(matches[2])

		if err != nil {
			panic("Bad argument: " + matches[2])
		}
		instructions = append(instructions, instruction{matches[1], val})
	}

	part1result := part1(instructions)
	part2result := part2(instructions)

	return []string{part1result, part2result}
}

func intcode(instructions []instruction) (int, int, int) {
	pc := 0
	acc := 0
	steps := 0
	seen := make(map[int]struct{})
	empty := struct{}{}

	for pc < len(instructions) {
		if _, visited := seen[pc]; visited {
			break
		}
		seen[pc] = empty

		instr := instructions[pc]
		steps++

		if steps == math.MaxInt {
			panic("Too many steps!")
		}

		switch instr.instr {
		case "acc":
			acc += instr.arg
			pc++
		case "jmp":
			pc += instr.arg
		case "nop":
			pc++
		}
	}

	return acc, steps, pc
}

func part1(instructions []instruction) string {
	acc, steps, pc := intcode(instructions)

	return fmt.Sprintf("After %d steps, pc=%d, acc=%d", steps, pc, acc)
}

func part2(instructions []instruction) string {
	nextInstr := make([]instruction, len(instructions))

	var acc, steps, pc int

	for i := 0; i < len(instructions); i++ {
		nextInstr = append([]instruction(nil), instructions...)

		if instructions[i].instr == "nop" {
			nextInstr[i].instr = "jmp"
		} else if instructions[i].instr == "jmp" {
			nextInstr[i].instr = "nop"
		} else {
			continue
		}

		acc, steps, pc = intcode(nextInstr)

		if pc == len(instructions) {
			break
		}
	}

	return fmt.Sprintf("After %d steps, pc=%d=%d, acc=%d", steps, pc, len(instructions), acc)
}
