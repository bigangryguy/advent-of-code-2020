package main

import (
	"fmt"
	"testing"
)

var nums []int

func TestPart1(t *testing.T) {
	actual := part1(nums)
	expected := 514579
	if actual != expected {
		t.Errorf("part1 = %d, expected %d", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := part2(nums)
	expected := 241861950
	if actual != expected {
		t.Errorf("part2 = %d, expected %d", actual, expected)
	}
}

func TestRecursivePart1(t *testing.T) {
	actual := recursive(nums, 2020, make([]int, 0), 2)
	expected := 514579
	if actual != expected {
		t.Errorf("recursive part1 = %d, expected %d", actual, expected)
	}
}

func TestRecursivePart2(t *testing.T) {
	actual := recursive(nums, 2020, make([]int, 0), 3)
	expected := 241861950
	if actual != expected {
		t.Errorf("recursive part2 = %d, expected %d", actual, expected)
	}
}

func TestMain(m *testing.M) {
	var err error
	nums, err = getInput("day1_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}
}

func BenchmarkPart1(b *testing.B) {
	nums, err := getInput("day1_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	_ = part1(nums)
}

func BenchmarkPart2(b *testing.B) {
	nums, err := getInput("day1_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	_ = part2(nums)
}

func BenchmarkRecursivePart1(b *testing.B) {
	nums, err := getInput("day1_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	_ = recursive(nums, 2020, make([]int, 0), 2)
}

func BenchmarkRecursivePart2(b *testing.B) {
	nums, err := getInput("day1_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	_ = recursive(nums, 2020, make([]int, 0), 3)
}
