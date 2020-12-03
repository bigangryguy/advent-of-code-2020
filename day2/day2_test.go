package main

import (
	"fmt"
	"testing"
)

var lines []string

func TestPart1(t *testing.T) {
	actual := part1(lines)
	expected := 2
	if actual != expected {
		t.Errorf("part1 = %d, expected %d", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := part2(lines)
	expected := 1
	if actual != expected {
		t.Errorf("part2 = %d, expected %d", actual, expected)
	}
}

func TestMain(m *testing.M) {
	var err error
	lines, err = getInput("day2_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}
}
