package day4

import (
	"fmt"
	"regexp"
	"strconv"
)

func isPresent(fields map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range required {
		_, present := fields[field]

		if !present {
			return false
		}
	}

	return true
}

func isValid(fields map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range required {
		value, present := fields[field]

		if !present {
			return false
		}

		if field == "byr" {
			year, err := strconv.Atoi(value)

			if err != nil || year < 1920 || year > 2002 {
				return false
			}
		} else if field == "iyr" {
			year, err := strconv.Atoi(value)

			if err != nil || year < 2010 || year > 2020 {
				return false
			}
		} else if field == "eyr" {
			year, err := strconv.Atoi(value)

			if err != nil || year < 2020 || year > 2030 {
				return false
			}
		} else if field == "hgt" {
			val, unit := 0, ""
			fmt.Sscanf(value, "%d%s", &val, &unit)

			if (unit != "cm" && unit != "in") ||
				(unit == "cm" && (val < 150 || val > 193)) ||
				(unit == "in" && (val < 59 || val > 76)) {
				return false
			}
		} else if field == "hcl" {
			if matched, err := regexp.MatchString("^#[0-9a-f]{6}$", value); !matched || err != nil {
				return false
			}
		} else if field == "ecl" {
			if matched, err := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", value); !matched || err != nil {
				return false
			}
		} else if field == "pid" {
			if matched, err := regexp.MatchString("^\\d{9}$", value); !matched || err != nil {
				return false
			}
		}

	}

	return true
}
func Day4(input []string) []string {
	count := 0

	input = append(input, "")

	fields := make(map[string]string)
	numPresent := 0
	numValid := 0

	for i := 0; i < len(input); i++ {
		line := input[i] + " "

		if line == " " {
			count++
			if isPresent(fields) {
				numPresent++
			}
			if isValid(fields) {
				numValid++
			}
			fields = map[string]string{}
			continue
		}

		field := ""
		val := ""

		for c := 0; c < len(line); c++ {
			chr := rune(line[c])

			if chr == ':' {
				field = val
				val = ""
			} else if chr == ' ' {
				fields[field] = val
				field = ""
				val = ""
			} else {
				val += string(chr)
			}
		}
	}

	return []string{fmt.Sprintf("Of %d entries, %d have required fields, %d are valid", count, numPresent, numValid)}
}
