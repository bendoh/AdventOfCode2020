package day16

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var fields map[string][][]int

type empty interface{}

var isEmpty empty

func valInField(val int, field string) bool {
	for _, interval := range fields[field] {
		if val >= interval[0] && val <= interval[1] {
			return true
		}
	}

	return false
}

func scanTicket(ticket []int) (error, int) {
	for _, val := range ticket {
		foundValue := false

		for field, _ := range fields {
			if valInField(val, field) {
				foundValue = true
				break
			}
		}

		if !foundValue {
			return errors.New("invalid ticket"), val
		}
	}

	return nil, 0
}

func printTicket(label string, ticket []int) {
	var ticketStr []string
	for _, val := range ticket {
		ticketStr = append(ticketStr, fmt.Sprintf("%d", val))
	}
	fmt.Printf("%s: %s\n", label, strings.Join(ticketStr, ","))
}

func reduceTicketFields(ticketFields []map[string]interface{}) []string {
	result := make([]string, len(ticketFields))
	foundFields := make(map[string]interface{})

	found := 0

	for found < len(ticketFields) {
		for i, fields := range ticketFields {
			for field, _ := range foundFields {
				delete(fields, field)
			}

			if len(fields) == 1 {
				for field, _ := range fields {
					foundFields[field] = isEmpty
					result[i] = field
					found++
				}
			}
		}
	}

	return result
}

func Day16(input []string) []string {
	fieldPat := regexp.MustCompile("^([\\w\\s]+): (.*)")
	fields = make(map[string][][]int)

	inMyTicket := true

	myTicket := make([]int, 0)
	otherTickets := make([][]int, 0)

	for _, line := range input {
		matches := fieldPat.FindStringSubmatch(line)

		if len(matches) > 0 {
			field := matches[1]
			ranges := matches[2]

			for _, r := range strings.Split(ranges, " or ") {
				pair := strings.Split(r, "-")

				start, err := strconv.Atoi(pair[0])
				if err != nil {
					return []string{"Error"}
				}
				end, err := strconv.Atoi(pair[1])
				if err != nil {
					return []string{"Error"}
				}
				fields[field] = append(fields[field], []int{start, end})
			}
			continue
		} else if line == "your ticket:" {
			continue
		} else if line == "nearby tickets:" {
			inMyTicket = false
			continue
		} else if line == "" {
			continue
		}

		var ticket []int
		for _, tp := range strings.Split(line, ",") {
			intVal, _ := strconv.Atoi(tp)
			ticket = append(ticket, intVal)
		}

		if inMyTicket {
			myTicket = ticket
		} else {
			otherTickets = append(otherTickets, ticket)
		}
	}

	printTicket("My ticket", myTicket)

	validTickets := make([]int, 0)
	invalidDigits := make([]int, 0)
	for ticketNum, ticket := range otherTickets {
		printTicket(fmt.Sprintf("Ticket %d", ticketNum), ticket)

		err, invalidDigit := scanTicket(ticket)

		if err != nil {
			invalidDigits = append(invalidDigits, invalidDigit)
		} else {
			validTickets = append(validTickets, ticketNum)
		}
	}

	results := make([]string, 0)
	invalidSum := 0

	for _, invalidDigit := range invalidDigits {
		results = append(results, fmt.Sprintf("+ %d", invalidDigit))
		invalidSum += invalidDigit
	}
	results = append(results, fmt.Sprintf("sum of invalid tickets: %d", invalidSum))

	ticketFields := make([]map[string]interface{}, len(myTicket))

	for ticketPos, val := range myTicket {
		ticketFields[ticketPos] = make(map[string]interface{})

		for field, ranges := range fields {
			for _, r := range ranges {
				if val >= r[0] && val <= r[1] {
					ticketFields[ticketPos][field] = isEmpty
				}
			}
		}
	}
	for _, ticketNum := range validTickets {
		for ticketPos, val := range otherTickets[ticketNum] {
			for field, ranges := range fields {
				for _, r := range ranges {
					if val >= r[0] && val <= r[1] {
						ticketFields[ticketPos][field] = isEmpty
					}
				}
			}
		}
	}

	for _, ticketNum := range validTickets {
		for pos, val := range otherTickets[ticketNum] {

			keptFields := make(map[string]interface{}, 0)
			for field, _ := range ticketFields[pos] {
				if valInField(val, field) {
					keptFields[field] = isEmpty
				}
			}
			ticketFields[pos] = keptFields
		}
	}

	positionFields := reduceTicketFields(ticketFields)

	result := int64(1)
	departureResult := int64(1)
	for pos, field := range positionFields {
		results = append(results, fmt.Sprintf("Adding %s=%d, result=%d", field, myTicket[pos], result))
		result *= int64(myTicket[pos])

		if strings.Contains(field, "departure") {
			departureResult *= int64(myTicket[pos])
			step := fmt.Sprintf("Departure field %s=%d result=%d", field, myTicket[pos], departureResult)
			fmt.Println(step)
			results = append(results, step)
		}
	}

	// Then iterate through valid tickets with the base set of fields
	return results
}
