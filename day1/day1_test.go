package main

import (
	"fmt"
	"testing"
)

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
