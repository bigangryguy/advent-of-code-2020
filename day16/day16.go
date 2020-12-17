package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Range struct {
	Min, Max int
}

type Rule struct {
	Name string
	Range1, Range2 Range
}

func (r Rule) isValueValid(value int) bool {
	return (value >= r.Range1.Min && value <= r.Range1.Max) ||
		(value >= r.Range2.Min && value <= r.Range2.Max)
}

type Ticket struct {
	Values []int
}

func (t Ticket) getInvalidValues(rules []Rule) (invalids []int) {
	for _, value := range t.Values {
		if !validForAnyRule(value, rules) {
			invalids = append(invalids, value)
		}
	}
	return
}

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func ruleFromText(text string) (r Rule, err error) {
	tokens := strings.Split(text, ": ")

	name := tokens[0]

	ranges := strings.Split(tokens[1], " or ")

	range1Parts := strings.Split(ranges[0], "-")
	var range1Min, range1Max int
	range1Min, err = strconv.Atoi(range1Parts[0])
	if err != nil {
		return
	}
	range1Max, err = strconv.Atoi(range1Parts[1])
	if err != nil {
		return
	}

	range2Parts := strings.Split(ranges[1], "-")
	var range2Min, range2Max int
	range2Min, err = strconv.Atoi(range2Parts[0])
	if err != nil {
		return
	}
	range2Max, err = strconv.Atoi(range2Parts[1])
	if err != nil {
		return
	}

	r = Rule {
		Name: name,
		Range1: Range { Min: range1Min, Max: range1Max },
		Range2: Range { Min: range2Min, Max: range2Max },
	}
	return
}

func ticketFromText(text string) (t Ticket, err error) {
	tokens := strings.Split(text, ",")

	var values []int
	for _, token := range tokens {
		var num int
		num, err = strconv.Atoi(token)
		if err != nil {
			return
		}
		values = append(values, num)
	}
	t = Ticket {
		Values: values,
	}
	return
}

func parseInput(lines []string) (rules []Rule, yourTicket Ticket, nearbyTickets []Ticket, err error) {
	nextLineIsYourTicket := false
	nextLinesAreNearbyTickets := false

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line == "your ticket:" {
			nextLineIsYourTicket = true
			continue
		}
		if line == "nearby tickets:" {
			nextLinesAreNearbyTickets = true
			continue
		}

		if nextLineIsYourTicket {
			nextLineIsYourTicket = false
			yourTicket, err = ticketFromText(line)
			if err != nil {
				return
			}
		} else if nextLinesAreNearbyTickets {
			var ticket Ticket
			ticket, err = ticketFromText(line)
			if err != nil {
				return
			}
			nearbyTickets = append(nearbyTickets, ticket)
		} else {
			var rule Rule
			rule, err = ruleFromText(line)
			if err != nil {
				return
			}
			rules = append(rules, rule)
		}
	}
	return
}

func validForAnyRule(value int, rules []Rule) bool {
	for _, rule := range rules {
		if rule.isValueValid(value) {
			return true
		}
	}
	return false
}

func sum(nums []int) (result int) {
	for _, num := range nums {
		result += num
	}
	return
}

func part1(lines []string) (result int, err error) {
	var rules []Rule
	var nearbyTickets []Ticket
	rules, _, nearbyTickets, err = parseInput(lines)
	if err != nil {
		return
	}

	for _, ticket := range nearbyTickets {
		invalid := ticket.getInvalidValues(rules)
		result += sum(invalid)
	}
	return
}

func main() {
	var lines []string
	var err error
	lines, err = getInput("day16_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	var part1Result int
	part1Result, err = part1(lines)
	fmt.Printf("Part 1 answer: %d\n", part1Result)
}
