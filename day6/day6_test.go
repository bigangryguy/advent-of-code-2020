package main

import (
	"fmt"
	"testing"
)

func Test_getGroups(t *testing.T) {
	lines := []string { "abc", "", "a", "b", "c", "", "ab", "ac", "", "a", "a", "a", "a", "", "b"}
	expected := []string { "abc", "a\nb\nc", "ab\nac", "a\na\na\na", "b" }
	actual := getGroups(lines)
	for i, group := range actual {
		if group != expected[i] {
			t.Errorf("getGroups = %v at index %d, expected %v", group, i, expected[i])
		}
	}
}

func Test_countUniqueQuestionsInGroup(t *testing.T) {
	var actual int

	groups := []string { "abc", "a\nb\nc", "ab\nac", "a\na\na\na", "b" }
	expected := []int { 3, 3, 3, 1, 1 }
	for i, group := range groups {
		actual = countUniqueQuestionsInGroup(group)
		if actual != expected[i] {
			t.Errorf("countUniqueQuestionsInGroup = %d for %v at index %d, expected %d", actual, group, i, expected[i])
		}
	}
}

func Test_countUnanimousQuestionsInGroup(t *testing.T) {
	var actual int

	groups := []string { "abc", "a\nb\nc", "ab\nac", "a\na\na\na", "b" }
	expected := []int { 3, 0, 1, 1, 1 }
	for i, group := range groups {
		actual = countUnanimousQuestionsInGroup(group)
		if actual != expected[i] {
			t.Errorf("countUnanimousQuestionsInGroup = %d for %v at index %d, expected %d", actual, group, i, expected[i])
		}
	}
}

func Test_part1(t *testing.T) {
	lines, err := getInput("day6_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	actual := part1(lines)
	expected := 11
	if actual != expected {
		t.Errorf("part1 = %d, expected %d", actual, expected)
	}
}

func Test_part2(t *testing.T) {
	lines, err := getInput("day6_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	actual := part2(lines)
	expected := 6
	if actual != expected {
		t.Errorf("part2 = %d, expected %d", actual, expected)
	}
}
