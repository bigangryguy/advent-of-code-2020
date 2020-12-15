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

func generateMasks(bitmask string) (orMask int64, andMask int64) {
	if len(bitmask) != 36 {
		return
	}

	for i, bit := range bitmask {
		var val int64 = 1 << (35 - i)
		if bit == '1' {
			andMask += val
			orMask += val
		} else if bit == 'X' {
			andMask += val
		}
	}
	return
}

func applyMasks(value int64, orMask int64, andMask int64) int64 {
	return (value | orMask) & andMask
}

func part1(lines []string) (result int64, err error) {
	memory := make(map[int]int64)
	var orMask int64
	var andMask int64
	for _, line := range lines {
		tokens := strings.Split(line, " = ")
		if tokens[0] == "mask" {
			orMask, andMask = generateMasks(tokens[1])
		} else {
			var memloc int
			memloc, err = strconv.Atoi(tokens[0][4:len(tokens[0])-1])
			if err != nil {
				return
			}
			var value int64
			value, err = strconv.ParseInt(tokens[1], 10, 64)
			if err != nil {
				return
			}
			value = applyMasks(value, orMask, andMask)
			memory[memloc] = value
		}
	}
	for _, values := range memory {
		result += values
	}
	return
}

func main() {
	lines, err := getInput("day14_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result, err := part1(lines)
	if err != nil {
		fmt.Printf("Received unexpected error: %v\n", err)
	}

	fmt.Printf("Part 1 answer: %d\n", part1Result)
}
