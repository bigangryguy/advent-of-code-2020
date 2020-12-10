package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	nums, err := getInput("day10_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := 35
	actual := part1(nums)
	if actual != expected {
		t.Errorf("part1 = %d, expected = %d\n", actual, expected)
	}

	nums, err = getInput("day10_test_input2.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected = 220
	actual = part1(nums)
	if actual != expected {
		t.Errorf("part1 = %d, expected = %d\n", actual, expected)
	}
}
