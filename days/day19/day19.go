package day19

import (
	"fmt"
	"strings"
)

type rule [][]string

var rules map[string]rule

func Day19(input []string) []string {
	rules = make(map[string]rule)

	results := make([]string, 0)

	messages := parseRules(input)

	results = append(results, matchRules(messages)...)

	rules["8"] = [][]string{{"42"}, {"42", "8"}}
	rules["11"] = [][]string{{"42", "31"}, {"42", "11", "31"}}

	results = append(results, matchRules(messages)...)

	return results
}

func parseRules(input []string) []string {
	messages := make([]string, 0)
	for _, line := range input {
		colon := strings.Index(line, ":")

		if colon > 0 {
			ruleNo := line[0:colon]
			rules[ruleNo] = parseRule(line[colon+1:])
		} else if len(line) > 0 {
			messages = append(messages, line)
		}
	}
	return messages
}

func parseRule(input string) rule {
	groups := strings.Split(input, "|")
	result := make(rule, 0, len(groups))

	for _, group := range groups {
		ruleNums := strings.Split(strings.Trim(group, " "), " ")
		groupRules := make([]string, 0, len(ruleNums))

		for _, ruleNo := range ruleNums {
			groupRules = append(groupRules, ruleNo)
		}
		result = append(result, groupRules)
	}

	return result
}

func matchRules(messages []string) []string {
	matches := 0
	results := make([]string, 0)

	for _, line := range messages {
		lineResult := fmt.Sprintf("%s... ", line)
		ruleMatches, idx := matchesRule(line, "0")

		if ruleMatches && idx == len(line) {
			lineResult += fmt.Sprintf("MATCHES: %d=%d", idx, len(line))
			matches++
		}

		results = append(results, lineResult)
	}

	return append(results, fmt.Sprintf("%d total matches", matches))
}

func matchesRule(part string, ruleNum string) (bool, int) {
	thisRule := rules[ruleNum]

	if thisRule[0][0] == `"a"` || thisRule[0][0] == `"b"` {
		chr := thisRule[0][0][1]

		if part[0] == chr {
			return true, 1
		}

		return false, 0
	} else {
		matchingRules := make(map[int]int)

		if ruleNum == "8" && len(thisRule) == 2 {
			idx := 0

			for idx < len(part) {
				if ruleMatches, offset := matchesRule(part[idx:], "42"); ruleMatches {
					idx += offset
				} else {
					break
				}
			}

			if idx > 0 {
				return true, idx
			} else {
				return false, 0
			}
		} else if ruleNum == "11" && len(thisRule) == 2 {
			idx := 0
			numMatches := 0

			for idx < len(part) {
				if ruleMatches, o := matchesRule(part[idx:], "42"); ruleMatches {
					idx += o
					numMatches++
				} else {
					break
				}
			}

			for numMatches > 0 && idx < len(part) {
				if ruleMatches, o := matchesRule(part[idx:], "31"); ruleMatches {
					idx += o
					numMatches--
				} else {
					break
				}
			}

			if idx > 0 && numMatches == 0 {
				return true, idx
			} else {
				return false, 0
			}
		}

		for i, nextRules := range thisRule {
			matches := true
			idx := 0
			for _, nextRule := range nextRules {
				ruleMatches, offset := matchesRule(part[idx:], nextRule)

				if !ruleMatches {
					matches = false
					break
				}
				idx += offset
			}
			if matches {
				matchingRules[i] = idx
				return true, idx
			}
		}

		return false, 0
	}
}
