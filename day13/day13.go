package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput(filename string) (timestamp int, busIds []int, err error) {
	var data []byte
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error reading file: %q", err))
		return
	}

	lines := strings.Split(string(data), "\n")
	timestamp, err = strconv.Atoi(lines[0])
	if err != nil {
		return
	}
	parts := strings.Split(lines[1], ",")
	for _, part := range parts {
		if part == "x" {
			busIds = append(busIds, -1)
			continue
		}
		var id int
		id, err = strconv.Atoi(part)
		if err != nil {
			return
		}
		busIds = append(busIds, id)
	}
	return
}

func getNextDeparture(timestamp int, busId int) int {
	return timestamp + busId - (timestamp % busId)
}

func getEarliestNextDeparture(timestamp int, busIds []int) (earliestBusId int, departure int) {
	earliestBusId = busIds[0]
	departure = getNextDeparture(timestamp, busIds[0])
	for _, busId := range busIds[1:] {
		if busId == -1 {
			continue
		}
		value := getNextDeparture(timestamp, busId)
		if value < departure {
			earliestBusId = busId
			departure = value
		}
	}
	return
}

func moduloInverse(a int, m int) int {
	if m == 1 {
		return 0
	}
	m0 := m
	x0 := 0
	x1 := 1
	for {
		if a <= 1 {
			break
		}
		q := a / m
		t := m
		m = a % m
		a = t
		t = x0
		x0 = x1 - q * x0
		x1 = t
	}

	if x1 < 0 {
		x1 += m0
	}

	return x1
}

func product(nums []int) (result int) {
	result = 1
	for _, num := range nums {
		result *= num
	}
	return
}

func sum(nums []int) (result int) {
	for _, num := range nums {
		result += num
	}
	return
}

// Gauss Chinese Remainder Theorem
func crt(nums []int, remainders []int) int {
	n := product(nums)
	var xs []int
	for i, num := range nums {
		ni := n / num
		xs = append(xs, remainders[i] * ni * moduloInverse(ni, num))
	}
	xcon := sum(xs)
	return xcon % n
}

func part1(timestamp int, busIds []int) int {
	busId, departure := getEarliestNextDeparture(timestamp, busIds)
	return (departure - timestamp) * busId
}

func part2(busIds []int) int {
	var nums []int
	var remainders []int
	for i, num := range busIds {
		if num > 0 {
			nums = append(nums, num)
			remainders = append(remainders, num - i)
		}
	}
	return crt(nums, remainders)
}

func main() {
	timestamp, busIds, err := getInput("day13_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result := part1(timestamp, busIds)
	fmt.Printf("Part 1 answer: %d\n", part1Result)

	part2Result := part2(busIds)
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
