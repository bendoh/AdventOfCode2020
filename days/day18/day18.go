package day18

import (
	"fmt"
	"strconv"
	"strings"
)

type ast struct {
	op    rune
	val   int
	left  *ast
	right []ast
}

func compute(expr string) int {
	parts := strings.Split(expr, " ")
	val := 0

	for len(parts) > 0 {
		if parts[0] == "+" || parts[0] == "*" {
			opInt, err := strconv.Atoi(parts[1])

			if err != nil {
				panic("Unparseable")
			}

			if parts[0] == "+" {
				val = val + opInt
			} else if parts[0] == "*" {
				val = val * opInt
			}

			parts = parts[2:]
		} else {
			opInt, err := strconv.Atoi(parts[0])
			parts = parts[1:]
			val = opInt

			if err != nil {
				panic("Unparseable")
			}

		}
	}

	return val
}

func computeAdvanced(expr string) int {
	groups := strings.Split(expr, " * ")

	if len(groups) > 1 {
		product := computeAdvanced(groups[0])
		for _, group := range groups[1:] {
			var op int
			if strings.Contains(group, "+") {
				op = computeAdvanced(group)
			} else {
				var err error
				op, err = strconv.Atoi(group)
				if err != nil {
					panic(err)
				}
			}

			product *= op
		}

		return product
	}

	sum := 0
	parts := strings.Split(expr, " ")
	for len(parts) > 0 {
		if parts[0] == "+" {
			opInt, err := strconv.Atoi(parts[1])

			if err != nil {
				panic("Unparseable")
			}

			sum = sum + opInt

			parts = parts[2:]
		} else {
			opInt, err := strconv.Atoi(parts[0])
			parts = parts[1:]
			sum = opInt

			if err != nil {
				panic("Unparseable")
			}

		}
	}

	return sum
}

func expandAdvanced(expr string) string {
	if len(expr) == 0 {
		return ""
	}

	parenIdx := strings.Index(expr, "(")

	if parenIdx == -1 {
		return fmt.Sprintf("%d", computeAdvanced(expr))
	}

	for ; parenIdx >= 0; parenIdx = strings.Index(expr, "(") {
		depth := 0

		for i := parenIdx; i < len(expr); i++ {
			if expr[i] == '(' {
				depth++
			} else if expr[i] == ')' {
				depth--
			}

			if depth == 0 {
				expr = expr[0:parenIdx] + expandAdvanced(expr[parenIdx+1:i]) + expr[i+1:]
				break
			}
		}
	}

	return fmt.Sprintf("%d", computeAdvanced(expr))
}
func expand(expr string) string {
	if len(expr) == 0 {
		return ""
	}

	parenIdx := strings.Index(expr, "(")

	if parenIdx == -1 {
		return fmt.Sprintf("%d", compute(expr))
	}

	for ; parenIdx >= 0; parenIdx = strings.Index(expr, "(") {
		depth := 0

		for i := parenIdx; i < len(expr); i++ {
			if expr[i] == '(' {
				depth++
			} else if expr[i] == ')' {
				depth--
			}

			if depth == 0 {
				expr = expr[0:parenIdx] + expand(expr[parenIdx+1:i]) + expr[i+1:]
				break
			}
		}
	}

	return fmt.Sprintf("%d", compute(expr))
}

func Day18(input []string) []string {
	result := make([]string, 0)
	var sum, advancedSum int64

	for _, line := range input {
		lineResult := expand(line)
		advancedLineResult := expandAdvanced(line)

		result = append(result, fmt.Sprintf(
			"%s=%s (simple) %s (advanced)",
			line,
			lineResult,
			advancedLineResult))

		intVal, err := strconv.Atoi(lineResult)
		advancedIntVal, err := strconv.Atoi(advancedLineResult)

		if err == nil {
			sum += int64(intVal)
			advancedSum += int64(advancedIntVal)
		}
	}

	result = append(result, fmt.Sprintf("Total: %d, Advanced Total: %d", sum, advancedSum))
	return result
}
