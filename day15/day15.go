package main

import "fmt"

func vanEck(init []int, length int) (sequence []int) {
	sequence = make([]int, length)
	copy(sequence, init)
	seen := make(map[int]int)
	for n := 0; n < length-1; n++ {
		if m, ok := seen[sequence[n]]; ok {
			sequence[n+1] = n - m
		}
		seen[sequence[n]] = n
	}
	return
}

func part1() int {
	length := 2020
	sequence := vanEck([]int { 10,16,6,0,1,17 }, length)
	return sequence[length-1]
}

func part2() int {
	length := 30000000
	sequence := vanEck([]int { 10,16,6,0,1,17 }, length)
	return sequence[length-1]
}

func main() {
	part1Result := part1()
	fmt.Printf("Part 1 answer = %d\n", part1Result)

	part2Result := part2()
	fmt.Printf("Part 2 answer = %d\n", part2Result)
}
