package main

import (
	"fmt"
	"testing"
)

var grid []string

func TestPart1(t *testing.T) {
	actual := part1(Slope{3, 1}, grid)
	expected := 7
	if actual != expected {
		t.Errorf("part1 = %d, expected %d", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := part2(grid)
	expected := 336
	if actual != expected {
		t.Errorf("part1 = %d, expected %d", actual, expected)
	}
}

func TestMain(m *testing.M) {
	var err error
	grid, err = getInput("day3_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}
}
