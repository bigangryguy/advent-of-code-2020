package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func getInput(filename string) (nums []int, err error) {
	var data []byte
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		var num int
		num, err = strconv.Atoi(line)
		if err != nil {
			return
		}
		nums = append(nums, num)
	}
	nums = append(nums, 0)
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)
	return
}

func part1(nums []int) int {
	var diff1, diff3 int
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
	}
	return diff1 * diff3
}

func part2(nums []int) int {
	max := nums[len(nums)-1]
	adapters := map[int][]int { 0: []int { nums[0] }, max: []int { max+3 }}
	for i := 0; i < len(nums)-1; i++ {
		adapters[nums[i]] = []int {}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] - nums[i] <= 3 {
				adapters[nums[i]] = append(adapters[nums[i]], nums[j])
			} else {
				break
			}
		}
	}
	cache := map[int]int { max+3: 1 }
	for i := len(nums)-1; i >= 0; i-- {
		val := nums[i]
		var combos int
		for j := 0; j < len(adapters[val]); j++ {
			combos += cache[adapters[val][j]]
		}
		cache[val] = combos
	}
	return cache[0]
}

func main() {
	nums, err := getInput("day10_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result := part1(nums)
	fmt.Printf("Part 1 answer: %d\n", part1Result)

	part2Result := part2(nums)
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
