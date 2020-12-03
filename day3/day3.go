package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Slope struct {
	Horizontal int
	Vertical int
}

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func part1(slope Slope, grid []string) (count int) {
	colLimit := len(grid[0])

	var col int
	for i := 0; i < len(grid); i += slope.Vertical {
		if grid[i][col % colLimit] == '#' {
			count++
		}

		col += slope.Horizontal
	}

	return
}

func part2(slopes []Slope, grid []string) (result int) {
	result = 1
	for _, v := range slopes {
		result *= part1(v, grid)
	}

	return
}

func main() {
	lines, err := getInput("day3_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result := part1(Slope{3, 1}, lines)

	slopes := []Slope {
		Slope{1, 1},
		Slope{3, 1},
		Slope{5, 1},
		Slope{7, 1},
		Slope{1, 2},
	}
	part2Result := part2(slopes, lines)

	fmt.Printf("Part 1 result: %d\n", part1Result)
	fmt.Printf("Part 2 result: %d\n", part2Result)
}
