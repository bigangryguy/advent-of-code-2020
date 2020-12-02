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

func part1(min int, max int, char string, password string) bool {
	var count int
	for i := 0; i < len(password) && count <= max; i++ {
		if char == string(password[i]) {
			count++
		}
	}
	return count >= min && count <= max
}

func part2(pos1 int, pos2 int, char string, password string) bool {
	pos1--
	pos2--

	inPos1 := pos1 < len(password) && string(password[pos1]) == char
	inPos2 := pos2 < len(password) && string(password[pos2]) == char
	return (inPos1 || inPos2) && (inPos1 != inPos2)
}

func main() {
	lines, err := getInput("day2_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	var part1Valid int
	var part2Valid int
	for i := 0; i < len(lines); i++ {
		num1, num2, char, password := splitInput(lines[i])

		if part1(num1, num2, char, password) {
			part1Valid++
		}
		if part2(num1, num2, char, password) {
			part2Valid++
		}
	}

	fmt.Printf("Valid passwords (Part 1): %d\n", part1Valid)
	fmt.Printf("Valid passwords (Part 2): %d\n", part2Valid)
}
