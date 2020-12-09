package main

import (
	"testing"
)

func Test_canSum(t *testing.T) {
	nums, err := getInput("day9_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	targets := []int { 55, 164, 853, 128 }
	expected := []bool { true, true, true, false }
	for i, target := range targets {
		actual := canSum(nums, target)
		if actual != expected[i] {
			t.Errorf("canSum = %v for %d, expected %v\n", actual, target, expected[i])
		}
	}
}

func Test_findInvalidNumber(t *testing.T) {
	nums, err := getInput("day9_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := 127
	actual := findInvalidNumber(nums, 5)
	if actual != expected {
		t.Errorf("findInvalidNumber = %d, expected %d\n", actual, expected)
	}
}

func Test_findContiguousSum(t *testing.T) {
	nums, err := getInput("day9_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := []int { 15, 25, 47, 40 }
	actual := findContiguousSum(nums, 127)
	for i, actualElement := range actual {
		if actualElement != expected[i] {
			t.Errorf("findContiguousSum = %d at index %d, expected %d. Full actual = %v\n",
				actualElement, i, expected[i], actual)
		}
	}
}

func Test_part2(t *testing.T) {
	nums, err := getInput("day9_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := 62
	actual := part2(nums, 127)
	if actual != expected {
		t.Errorf("part2 = %d, expected %d\n", actual, expected)
	}
}
