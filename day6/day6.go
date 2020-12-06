package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func getGroups(lines []string) (groups []string) {
	var group string
	for _, line := range lines {
		if line != "" {
			group += line + "\n"
		} else {
			group = strings.TrimRight(group, "\n")
			groups = append(groups, group)
			group = ""
		}
	}
	if group != "" && group != "\n" {
		group = strings.TrimRight(group, "\n")
		groups = append(groups, group)
	}
	return
}

func countUniqueQuestionsInGroup(group string) int {
	questions := make(map[string]int)
	people := strings.Split(group, "\n")
	for _, person := range people {
		for _, answer := range person {
			questions[string(answer)]++
		}
	}
	return len(questions)
}

func countUnanimousQuestionsInGroup(group string) (result int) {
	questions := make(map[string]int)
	people := strings.Split(group, "\n")
	for _, person := range people {
		for _, answer := range person {
			questions[string(answer)]++
		}
	}
	for _, v := range questions {
		if v == len(people) {
			result++
		}
	}
	return
}

func part1(lines []string) (result int) {
	groups := getGroups(lines)
	for _, group := range groups {
		result += countUniqueQuestionsInGroup(group)
	}
	return
}

func part2(lines []string) (result int) {
	groups := getGroups(lines)
	for _, group := range groups {
		result += countUnanimousQuestionsInGroup(group)
	}
	return
}

func main() {
	lines, err := getInput("day6_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result := part1(lines)
	part2Result := part2(lines)

	fmt.Printf("Part 1 answer: %d\n", part1Result)
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
