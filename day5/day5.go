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

func parseRow(row string) (result int, err error) {
	if len(row) != 7 {
		err = errors.New("Row section must be 7 characters")
		return
	}

	min := 0
	max := 127

	for _, c := range row {
		split := (min + max) / 2
		if c == 'F' {
			max = split
		} else if c == 'B' {
			min = split + 1
		} else {
			err = errors.New("Row section can only contain F or B characters")
		}
	}
	// At this point, min == max
	result = min
	return
}

func parseColumn(column string) (result int, err error) {
	if len(column) != 3 {
		err = errors.New("Column section must be 3 characters")
		return
	}

	min := 0
	max := 7

	for _, c := range column {
		split := (min + max) / 2
		if c == 'L' {
			max = split
		} else if c == 'R' {
			min = split + 1
		} else {
			err = errors.New("Column section can only contain L or R characters")
		}
	}
	// At this point, min == max
	result = min
	return
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
