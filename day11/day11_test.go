package main

import (
	"fmt"
	"testing"
)

func Test_parseInput(t *testing.T) {
	lines, err := getInput("day11_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}
	expected := Layout {
		Width: 10,
		Height: 10,
		Seats: []string {
			"L", ".", "L", "L", ".", "L", "L", ".", "L", "L",
			"L", "L", "L", "L", "L", "L", "L", ".", "L", "L",
			"L", ".", "L", ".", "L", ".", ".", "L", ".", ".",
			"L", "L", "L", "L", ".", "L", "L", ".", "L", "L",
			"L", ".", "L", "L", ".", "L", "L", ".", "L", "L",
			"L", ".", "L", "L", "L", "L", "L", ".", "L", "L",
			".", ".", "L", ".", "L", ".", ".", ".", ".", ".",
			"L", "L", "L", "L", "L", "L", "L", "L", "L", "L",
			"L", ".", "L", "L", "L", "L", "L", "L", ".", "L",
			"L", ".", "L", "L", "L", "L", "L", ".", "L", "L",
		},
	}
	actual := parseInput(lines)
	if !areLayoutsSame(actual, expected) {
		t.Errorf("parseInput = %v, expected %v\n", actual, expected)
	}
}

func TestLayout_SeatAt(t *testing.T) {
	lines, err := getInput("day11_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}
	layout := parseInput(lines)

	type SeatAtTests struct {
		X, Y int
		Expected string
	}
	testCases := []SeatAtTests{
		{ 1, 0, "." },
		{ 4, 3, "." },
		{ 0, 9, "L" },
		{ 7, 7, "L" },
	}
	for _, testCase := range testCases {
		actual := layout.seatAt(testCase.X, testCase.Y)
		if actual != testCase.Expected {
			t.Errorf("seatAt = %v for %d, %d, expected %v\n",
				actual, testCase.X, testCase.Y, testCase.Expected)
		}
	}
}

func TestLayout_Copy(t *testing.T) {
	lines, err := getInput("day11_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}
	layout := parseInput(lines)

	clone := layout.copy()

	if !areLayoutsSame(clone, layout) {
		t.Errorf("copy failed")
	}
}

func Test_intMin(t *testing.T) {
	testCases := [6][3]int {
		{-1, 0, -1},
		{1, 0, 0},
		{0, 0, 0},
		{50, 66, 50},
		{83, 12, 12},
		{42, 42, 42},
	}
	for _, testCase := range testCases{
		actual := intMin(testCase[0], testCase[1])
		if actual != testCase[2] {
			t.Errorf("intMin = %d, expected %d\n", actual, testCase[2])
		}
	}
}

func Test_intMax(t *testing.T) {
	testCases := [6][3]int {
		{-1, 0, 0},
		{1, 0, 1},
		{0, 0, 0},
		{50, 66, 66},
		{83, 12, 83},
		{42, 42, 42},
	}
	for _, testCase := range testCases{
		actual := intMax(testCase[0], testCase[1])
		if actual != testCase[2] {
			t.Errorf("intMax = %d, expected %d\n", actual, testCase[2])
		}
	}
}

func Test_applyRules(t *testing.T) {
	lines, err := getInput("day11_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}
	layout := parseInput(lines)
	
	expected := Layout {
		Width: 10,
		Height: 10,
		Seats: []string {
			"#", ".", "#", "#", ".", "#", "#", ".", "#", "#",
			"#", "#", "#", "#", "#", "#", "#", ".", "#", "#",
			"#", ".", "#", ".", "#", ".", ".", "#", ".", ".",
			"#", "#", "#", "#", ".", "#", "#", ".", "#", "#",
			"#", ".", "#", "#", ".", "#", "#", ".", "#", "#",
			"#", ".", "#", "#", "#", "#", "#", ".", "#", "#",
			".", ".", "#", ".", "#", ".", ".", ".", ".", ".",
			"#", "#", "#", "#", "#", "#", "#", "#", "#", "#",
			"#", ".", "#", "#", "#", "#", "#", "#", ".", "#",
			"#", ".", "#", "#", "#", "#", "#", ".", "#", "#",
		},
	}
	actual := applyRules(layout, 4)
	if !areLayoutsSame(actual, expected) {
		t.Error("applyRules round 1 failed")
	}
	expected.Seats = []string {
		"#", ".", "L", "L", ".", "L", "#", ".", "#", "#",
		"#", "L", "L", "L", "L", "L", "L", ".", "L", "#",
		"L", ".", "L", ".", "L", ".", ".", "L", ".", ".",
		"#", "L", "L", "L", ".", "L", "L", ".", "L", "#",
		"#", ".", "L", "L", ".", "L", "L", ".", "L", "L",
		"#", ".", "L", "L", "L", "L", "#", ".", "#", "#",
		".", ".", "L", ".", "L", ".", ".", ".", ".", ".",
		"#", "L", "L", "L", "L", "L", "L", "L", "L", "#",
		"#", ".", "L", "L", "L", "L", "L", "L", ".", "L",
		"#", ".", "#", "L", "L", "L", "L", ".", "#", "#",
	}
	actual = applyRules(actual, 4)
	if !areLayoutsSame(actual, expected) {
		t.Error("applyRules round 2 failed")
	}
}

func Test_applyRulesUntilStable(t *testing.T) {
	lines, err := getInput("day11_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}
	layout := parseInput(lines)

	expected := Layout {
		Width: 10,
		Height: 10,
		Seats: []string {
			"#", ".", "#", "L", ".", "L", "#", ".", "#", "#",
			"#", "L", "L", "L", "#", "L", "L", ".", "L", "#",
			"L", ".", "#", ".", "L", ".", ".", "#", ".", ".",
			"#", "L", "#", "#", ".", "#", "#", ".", "L", "#",
			"#", ".", "#", "L", ".", "L", "L", ".", "L", "L",
			"#", ".", "#", "L", "#", "L", "#", ".", "#", "#",
			".", ".", "L", ".", "L", ".", ".", ".", ".", ".",
			"#", "L", "#", "L", "#", "#", "L", "#", "L", "#",
			"#", ".", "L", "L", "L", "L", "L", "L", ".", "L",
			"#", ".", "#", "L", "#", "L", "#", ".", "#", "#",
		},
	}
	actual := applyRulesUntilStable(layout, 4, false)
	if !areLayoutsSame(actual, expected) {
		t.Error("applyRulesUntilStable failed")
	}
}

func Test_part1(t *testing.T) {
	lines, err := getInput("day11_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	expected := 37
	actual := part1(lines)
	if actual != expected {
		t.Errorf("part1 = %d, expected %d\n", actual, expected)
	}
}

func Test_part2(t *testing.T) {
	lines, err := getInput("day11_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	expected := 26
	actual := part2(lines)
	if actual != expected {
		t.Errorf("part2 = %d, expected %d\n", actual, expected)
	}
}
