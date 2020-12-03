package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func splitInput(line string) (num1 int, num2 int, char string, password string) {
	parts := strings.Split(line, " ")

	numRange := strings.Split(parts[0], "-")
	num1, _ = strconv.Atoi(numRange[0])
	num2, _ = strconv.Atoi(numRange[1])

	char = string(parts[1][0])

	password = parts[2]

	return
}

func part1(lines []string) (result int) {
	for _, line := range lines {
		min, max, char, password := splitInput(line)
		var count int
		// Could use strings.Count() here but expanding to a for loop
		// allows us to short-circuit the logic if count > max
		for i := 0; i < len(password) && count <= max; i++ {
			if char == string(password[i]) {
				count++
			}
		}
		if count >= min && result <= max {
			result++
		}
	}
	return
}

func part2(lines []string) (result int) {
	for _, line := range lines {
		pos1, pos2, char, password := splitInput(line)
		pos1--
		pos2--

		inPos1 := pos1 < len(password) && string(password[pos1]) == char
		inPos2 := pos2 < len(password) && string(password[pos2]) == char
		if (inPos1 || inPos2) && (inPos1 != inPos2) {
			result ++
		}
	}
	return
}

func main() {
	lines, err := getInput("day2_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Valid := part1(lines)
	part2Valid := part2(lines)

	fmt.Printf("Valid passwords (Part 1): %d\n", part1Valid)
	fmt.Printf("Valid passwords (Part 2): %d\n", part2Valid)
}
