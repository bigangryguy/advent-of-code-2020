package main

import (
	"testing"
)

func areBagsSame(bag1 Bag, bag2 Bag) bool {
	if bag1.Color != bag2.Color {
		return false
	}
	for k, v := range bag1.Contains {
		if _, found := bag2.Contains[k]; found {
			if bag2.Contains[k] != v {
				return false
			}
		} else {
			return false
		}
	}
	for k, v := range bag1.ContainedBy {
		if _, found := bag2.ContainedBy[k]; found {
			if bag2.ContainedBy[k] != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func Test_parseLines(t *testing.T) {
	lines, err := getInput("day7_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := map[string]Bag {
		"light red": {
			Color: "light red",
			Contains: map[string]int {
				"bright white": 1,
				"muted yellow": 2,
			},
			ContainedBy: map[string]int {},
		},
		"dark orange": {
			Color: "dark orange",
			Contains: map[string]int {
				"bright white": 3,
				"muted yellow": 4,
			},
			ContainedBy: map[string]int {},
		},
		"bright white": {
			Color: "bright white",
			Contains: map[string]int {
				"shiny gold": 1,
			},
			ContainedBy: map[string]int {
				"light red": 1,
				"dark orange": 3,
			},
		},
		"muted yellow": {
			Color: "muted yellow",
			Contains: map[string]int {
				"shiny gold": 2,
				"faded blue": 9,
			},
			ContainedBy: map[string]int {
				"light red": 2,
				"dark orange": 4,
			},
		},
		"shiny gold": {
			Color: "shiny gold",
			Contains: map[string]int {
				"dark olive": 1,
				"vibrant plum": 2,
			},
			ContainedBy: map[string]int {
				"bright white": 1,
				"muted yellow": 2,
			},
		},
		"dark olive": {
			Color: "dark olive",
			Contains: map[string]int {
				"faded blue": 3,
				"dotted black": 4,
			},
			ContainedBy: map[string]int {
				"shiny gold": 1,
			},
		},
		"vibrant plum": {
			Color: "vibrant plum",
			Contains: map[string]int {
				"faded blue": 5,
				"dotted black": 6,
			},
			ContainedBy: map[string]int {
				"shiny gold": 2,
			},
		},
		"faded blue": {
			Color: "faded blue",
			Contains: map[string]int {},
			ContainedBy: map[string]int {
				"dark olive": 3,
				"vibrant plum": 5,
				"muted yellow": 9,
			},
		},
		"dotted black": {
			Color: "dotted black",
			Contains: map[string]int {},
			ContainedBy: map[string]int {
				"dark olive": 4,
				"vibrant plum": 6,
			},
		},
	}
	actual := parseLines(lines)
	for actualColor, actualBag := range actual {
		if expectedBag, found := expected[actualColor]; found {
			if !areBagsSame(actualBag, expectedBag) {
				t.Errorf("parseLines: Expected %v, received %v\n", expectedBag, actualBag)
			}
		} else {
			t.Errorf("parseLines: Expected to find %s in bag list, did not\n", actualColor)
		}
	}
}

func Test_canContainBagColor(t *testing.T) {
	lines, err := getInput("day7_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	bags := parseLines(lines)
	expected := 4
	actual := canContainBagColor("shiny gold", bags)
	if actual != expected {
		t.Errorf("canContainBagColor = %d, expected %d\n", actual, expected)
	}
}

func Test_countInnerBags(t *testing.T) {
	lines, err := getInput("day7_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	bags := parseLines(lines)
	expected := 32
	actual := countInnerBags("shiny gold", bags)
	if actual != expected {
		t.Errorf("countInnerBags = %d, expected %d\n", actual, expected)
	}
}

func Test_part1(t *testing.T) {
	lines, err := getInput("day7_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := 4
	actual := part1(lines)
	if actual != expected {
		t.Errorf("part1 = %d, expected %d\n", actual, expected)
	}
}

func Test_part2(t *testing.T) {
	lines, err := getInput("day7_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := 32
	actual := part2(lines)
	if actual != expected {
		t.Errorf("part2 = %d, expected %d\n", actual, expected)
	}
}
