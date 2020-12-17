package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
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

func combinations(values []int64) (output []int64) {
	for i := 1; i < 1 << len(values); i++ {
		var value int64
		for j := 0; j < len(values); j++ {
			if (i >> j) & 1 == 1 {
				value += values[j]
			}
		}
		output = append(output, value)
	}
	output = append(output, 0)
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	return
}

func floatingBitsFromBitmask(bitmask string) (set []int64) {
	for i, bit := range bitmask {
		if bit == 'X' {
			set = append(set, 1 << (35 - i))
		}
	}
	set = combinations(set)
	return
}

func fixedBitsFromBitmask(bitmask string) (result int64) {
	for i, bit := range bitmask {
		if bit == '1' {
			result += 1 << (35 - i)
		}
	}
	return
}

func applyMasks(value int64, orMask int64, andMask int64) int64 {
	return (value | orMask) & andMask
}

func applyFloatingMask(value int64, bitmask string) (result []int64) {
	fixedMask := fixedBitsFromBitmask(bitmask)
	floatingMasks := floatingBitsFromBitmask(bitmask)
	base := value | fixedMask
	for _, mask := range floatingMasks {
		result = append(result, base ^ mask)
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return
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
	for _, value := range memory {
		result += value
	}
	return
}

func part2(lines []string) (result int64, err error) {
	memory := make(map[int64]int64)
	var bitmask string
	for _, line := range lines {
		tokens := strings.Split(line, " = ")
		if tokens[0] == "mask" {
			bitmask = tokens[1]
		} else {
			var originalMemloc int64
			originalMemloc, err = strconv.ParseInt(tokens[0][4:len(tokens[0])-1], 10, 64)
			if err != nil {
				return
			}
			var value int64
			value, err = strconv.ParseInt(tokens[1], 10, 64)
			if err != nil {
				return
			}
			memlocs := applyFloatingMask(originalMemloc, bitmask)
			for _, memloc := range memlocs {
				memory[memloc] = value
			}
		}
	}
	for _, value := range memory {
		result += value
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

	part2Result, err := part2(lines)
	if err != nil {
		fmt.Printf("Received unexpected error: %v\n", err)
	}
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
