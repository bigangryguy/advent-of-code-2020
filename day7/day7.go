package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bag struct {
	Color string
	Contains map[string]int
	ContainedBy map[string]int
}

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func parseLines(lines []string) (bags map[string]Bag) {
	bags = make(map[string]Bag)
	for _, line := range lines {
		line = strings.TrimRight(line, ".")
		parts := strings.Split(line, " bags contain ")
		color := parts[0]
		inners := strings.Split(parts[1], ", ")

		if _, found := bags[color]; !found {
			bags[color] = Bag{
				Color: color,
				Contains: make(map[string]int),
				ContainedBy: make(map[string]int),
			}
		}
		bag, _ := bags[color]
		if inners[0] == "no other bags" {
			continue
		}
		for _, inner := range inners {
			tokens := strings.Split(inner, " ")
			num, _ := strconv.Atoi(tokens[0])
			innerColor := strings.Join(tokens[1:len(tokens)-1], " ")
			bag.Contains[innerColor] = num

			if innerBag, ok := bags[innerColor]; ok {
				innerBag.ContainedBy[color] = num
			} else {
				bags[innerColor] = Bag{
					Color: innerColor,
					Contains: make(map[string]int),
					ContainedBy: map[string]int{
						color: num,
					},
				}
			}
		}

	}
	return
}

func canContainBagColor(color string, bags map[string]Bag) int {
	var colorsToCheck []string
	if bag, found := bags[color]; found {
		for k, _ := range bag.ContainedBy {
			colorsToCheck = append(colorsToCheck, k)
		}
	} else {
		return 0
	}

	outsideBags := make(map[string]int)
	for {
		if len(colorsToCheck) == 0 {
			break
		}
		color := colorsToCheck[0]
		colorsToCheck = colorsToCheck[1:]
		if bag, found := bags[color]; found {
			outsideBags[color]++
			for k, _ := range bag.ContainedBy {
				colorsToCheck = append(colorsToCheck, k)
			}
		}
	}
	return len(outsideBags)
}

func countInnerBags(color string, bags map[string]Bag) (result int) {
	var colorsToCheck []string
	if bag, found := bags[color]; found {
		for k, v := range bag.Contains {
			for i := 0; i < v; i++ {
				colorsToCheck = append(colorsToCheck, k)
			}
		}
	} else {
		return
	}

	for {
		if len(colorsToCheck) == 0 {
			break
		}
		color := colorsToCheck[0]
		colorsToCheck = colorsToCheck[1:]
		if bag, found := bags[color]; found {
			result++
			for k, v := range bag.Contains {
				for i := 0; i < v; i++ {
					colorsToCheck = append(colorsToCheck, k)
				}
			}
		}
	}
	return
}

func part1(lines []string) int {
	parsed := parseLines(lines)
	return canContainBagColor("shiny gold", parsed)
}

func part2(lines []string) int {
	parsed := parseLines(lines)
	return countInnerBags("shiny gold", parsed)
}

func main() {
	lines, err := getInput("day7_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result := part1(lines)
	part2Result := part2(lines)

	fmt.Printf("Part 1 answer: %d\n", part1Result)
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
