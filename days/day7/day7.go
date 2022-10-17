package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type contents map[string]int

// Part 2: Count how many bags a shiny gold bag needs to contain
func counter(toContents map[string]contents, container string, quantity int) int {
	total := quantity

	for bag, qty := range toContents[container] {
		total += quantity * counter(toContents, bag, qty)
	}

	return total
}
func Day7(input []string) []string {
	toContents := make(map[string]contents)
	toContainers := make(map[string][]string)
	topRe := regexp.MustCompile("(.+) bags contain (.+).")
	subRe := regexp.MustCompile("([0-9]+) (.+) bags?")

	for _, line := range input {
		matches := topRe.FindStringSubmatch(line)

		if len(matches) > 0 {
			container := matches[1]
			parts := strings.Split(matches[2], ", ")

			contains := make(contents)

			for _, part := range parts {
				subMatches := subRe.FindStringSubmatch(part)

				if subMatches != nil {
					qty, content := subMatches[1], subMatches[2]

					if _, ok := toContainers[content]; !ok {
						toContainers[content] = make([]string, 0)
					}

					containerList := toContainers[content]
					intval, err := strconv.Atoi(qty)

					if err == nil {
						contains[content] = intval
						toContainers[content] = append(containerList, container)
					}
				}
			}

			toContents[container] = contains
		}
	}

	// part1: Count number of different bags that might contain a shiny gold
	que := toContainers["shiny gold"]
	bagTypes := make(map[string]int)

	for len(que) > 0 {
		nextQue := make(map[string]int, 0)

		for len(que) > 0 {
			container := que[0]
			que = que[1:]

			bagTypes[container]++

			for _, parent := range toContainers[container] {
				nextQue[parent]++
			}
		}

		que = make([]string, 0, len(nextQue))

		for key, _ := range nextQue {
			que = append(que, key)
		}
	}

	return []string{
		fmt.Sprintf("%d different bag types can contain 'shiny gold'", len(bagTypes)),
		fmt.Sprintf("%d bags contained in 'shiny gold' bag", counter(toContents, "shiny gold", 1)-1),
	}
}
