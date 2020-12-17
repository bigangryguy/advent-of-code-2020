package main

import (
	"fmt"
	"testing"
)

func areRangesSame(r1 Range, r2 Range) bool {
	return r1.Min == r2.Min && r1.Max == r2.Max
}

func areRulesSame(r1 Rule, r2 Rule) bool {
	return r1.Name == r2.Name &&
		areRangesSame(r1.Range1, r2.Range1) &&
		areRangesSame(r1.Range2, r2.Range2)
}

func areTicketsSame(t1 Ticket, t2 Ticket) bool {
	same := len(t1.Values) == len(t2.Values)
	if same {
		for i := 0; i < len(t1.Values); i++ {
			if t1.Values[i] != t2.Values[i] {
				same = false
				break
			}
		}
	}
	return same
}

func Test_rangeFromText(t *testing.T) {
	texts := []string {
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
	}
	expected := []Rule {
		Rule {
			Name: "class",
			Range1: Range{ Min: 1, Max: 3 },
			Range2: Range{ Min: 5, Max: 7 },
		},
		Rule {
			Name: "row",
			Range1: Range{ Min: 6, Max: 11 },
			Range2: Range{ Min: 33, Max: 44 },
		},
		Rule {
			Name: "seat",
			Range1: Range{ Min: 13, Max: 40 },
			Range2: Range{ Min: 45, Max: 50 },
		},
	}
	for i, text := range texts {
		actual, err := ruleFromText(text)
		if err != nil {
			t.Fatal("Received unexpected error")
		}
		if !areRulesSame(actual, expected[i]) {
			t.Errorf("rangeFromText = %v, expected %v\n", actual, expected[i])
		}
	}
}

func Test_ticketFromText(t *testing.T) {
	texts := []string {
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}
	expected := []Ticket {
		Ticket {
			Values: []int { 7, 3, 47 },
		},
		Ticket {
			Values: []int { 40, 4, 50 },
		},
		Ticket {
			Values: []int { 55, 2, 20 },
		},
		Ticket {
			Values: []int { 38, 6, 12 },
		},
	}
	for i, text := range texts {
		actual, err := ticketFromText(text)
		if err != nil {
			t.Fatal("Received unexpected error")
		}
		if !areTicketsSame(actual, expected[i]) {
			t.Errorf("ticketFromText = %v, expected %v\n", actual, expected[i])
		}
	}
}

func Test_parseInput(t *testing.T) {
	lines, err := getInput("day16_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	expectedRules := []Rule {
		Rule {
			Name: "class",
			Range1: Range{ Min: 1, Max: 3 },
			Range2: Range{ Min: 5, Max: 7 },
		},
		Rule {
			Name: "row",
			Range1: Range{ Min: 6, Max: 11 },
			Range2: Range{ Min: 33, Max: 44 },
		},
		Rule {
			Name: "seat",
			Range1: Range{ Min: 13, Max: 40 },
			Range2: Range{ Min: 45, Max: 50 },
		},
	}
	expectedYourTicket :=
		Ticket {
			Values: []int { 7, 1, 14 },
		}
	expectedNearbyTickets := []Ticket {
		Ticket {
			Values: []int { 7, 3, 47 },
		},
		Ticket {
			Values: []int { 40, 4, 50 },
		},
		Ticket {
			Values: []int { 55, 2, 20 },
		},
		Ticket {
			Values: []int { 38, 6, 12 },
		},
	}

	var rules []Rule
	var yourTicket Ticket
	var nearbyTickets []Ticket
	rules, yourTicket, nearbyTickets, err = parseInput(lines)
	if err != nil {
		t.Fatal("Received unexpected error")
	}

	for i, rule := range rules {
		if !areRulesSame(rule, expectedRules[i]) {
			t.Errorf("parseInput: rule = %v at index %d, expected %v\n",
				rule, i, expectedRules[i])
		}
	}
	if !areTicketsSame(yourTicket, expectedYourTicket) {
		t.Errorf("parseInput: yourTicket = %v, expected %v\n", yourTicket, expectedYourTicket)
	}
	for i, ticket := range nearbyTickets {
		if !areTicketsSame(ticket, expectedNearbyTickets[i]) {
			t.Errorf("parseInput: nearbyTicket = %v at index %d, expected %v\n",
				ticket, i, expectedNearbyTickets[i])
		}
	}
}

func Test_validForAnyRule(t *testing.T) {
	rules := []Rule {
		Rule {
			Name: "class",
			Range1: Range{ Min: 1, Max: 3 },
			Range2: Range{ Min: 5, Max: 7 },
		},
		Rule {
			Name: "row",
			Range1: Range{ Min: 6, Max: 11 },
			Range2: Range{ Min: 33, Max: 44 },
		},
		Rule {
			Name: "seat",
			Range1: Range{ Min: 13, Max: 40 },
			Range2: Range{ Min: 45, Max: 50 },
		},
	}
	values := []int { 7, 47, 4, 55 }
	expected := []bool { true, true, false, false }

	for i, value := range values {
		actual := validForAnyRule(value, rules)
		if actual != expected[i] {
			t.Errorf("validForAnyRule = %v for %d, expected %v\n", actual, value, expected[i])
		}
	}
}

func Test_getInvalidValues(t *testing.T) {
	rules := []Rule {
		Rule {
			Name: "class",
			Range1: Range{ Min: 1, Max: 3 },
			Range2: Range{ Min: 5, Max: 7 },
		},
		Rule {
			Name: "row",
			Range1: Range{ Min: 6, Max: 11 },
			Range2: Range{ Min: 33, Max: 44 },
		},
		Rule {
			Name: "seat",
			Range1: Range{ Min: 13, Max: 40 },
			Range2: Range{ Min: 45, Max: 50 },
		},
	}
	tickets := []Ticket {
		Ticket {
			Values: []int { 7, 1, 14 },
		},
		Ticket {
			Values: []int { 7, 3, 47 },
		},
		Ticket {
			Values: []int { 40, 4, 50 },
		},
		Ticket {
			Values: []int { 55, 2, 20, 56 },
		},
	}
	expected := [4][]int {
		{},
		{},
		{ 4 },
		{ 55, 56 },
	}

	for i, ticket := range tickets {
		actual := ticket.getInvalidValues(rules)
		if len(actual) != len(expected[i]) {
			t.Errorf("Length of actual does not match length of expected\nactual = %v, expected = %v\n",
				actual, expected[i])
		}
		for j, actualValue := range actual {
			if actualValue != expected[i][j] {
				t.Errorf("getInvalidValues = %d at index %d, expected %d\n",
					actualValue, j, expected[i][j])
			}
		}
	}
}
