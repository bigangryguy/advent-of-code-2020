package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func doPartitioning(input string, expectedLength int, lowerChar string, upperChar string, min int, max int) (result int, err error) {
	if len(input) != expectedLength {
		err = errors.New("Input length does not match expected length")
		return
	}

	for _, c := range input {
		cStr := string(c)
		split := (min + max) / 2
		if cStr == lowerChar {
			max = split
		} else if cStr == upperChar {
			min = split + 1
		} else {
			err = errors.New(fmt.Sprintf("Input can only contain %v or %v characters", lowerChar, upperChar))
			return
		}
	}
	// At this point, min == max
	result = min
	return
}

func parseRow(row string) (result int, err error) {
	return doPartitioning(row, 7, "F", "B", 0, 127)
}

func parseColumn(column string) (result int, err error) {
	return doPartitioning(column, 3, "L", "R", 0, 7)
}

func getSeatID(line string) (result int, err error) {
	if len(line) != 10 {
		err = errors.New("Boarding pass line must be 10 characters")
		return
	}

	var row, column int

	row, err = parseRow(line[:7])
	if err != nil {
		return
	}
	column, err = parseColumn(line[7:])
	if err != nil {
		return
	}

	result = (row * 8) + column
	return
}

func getSeatList(lines []string) (list []int, err error) {
	var seatID int
	for _, line := range lines {
		seatID, err = getSeatID(line)
		if err != nil {
			return
		}
		list = append(list, seatID)
	}
	return
}

func part1(lines []string) (result int, err error) {
	var seatList []int
	seatList, err = getSeatList(lines)
	if err != nil {
		return
	}

	for _, seatID := range seatList {
		if seatID > result {
			result = seatID
		}
	}
	return
}

func part2(lines []string) (result int, err error) {
	var seatList []int
	seatList, err = getSeatList(lines)
	if err != nil {
		return
	}

	sort.Ints(seatList)
	for i, seatID := range seatList {
		if seatID < 8 {
			continue
		} else if seatID >= 960 {
			err = errors.New("Could not find seat")
			return
		}

		if seatList[i+1] == seatID + 2 {
			result = seatID + 1
			return
		}
	}
	return
}

func main() {
	lines, err := getInput("day5_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result, err := part1(lines)
	if err != nil {
		fmt.Println("Error getting part 1 result")
	}

	part2Result, err := part2(lines)
	if err != nil {
		fmt.Println("Error getting part 2 result")
	}

	fmt.Printf("Part 1 answer: %d\n", part1Result)
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
