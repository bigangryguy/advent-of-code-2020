package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput(filename string) ([]int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	lines := strings.Fields(string(data))
	nums := make([]int, len(lines), len(lines))
	for i := 0; i < len(lines); i++ {
		nums[i], _ = strconv.Atoi(lines[i])
	}
	return nums, nil
}

func part1(nums []int) (result int) {
	numsTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := 2020 - nums[i]
		if _, ok := numsTable[diff]; ok {
			result = nums[i] * diff
			break
		}
		numsTable[nums[i]] = i
	}
	return result
}

func part2(nums []int) (result int) {
	numsTable := make(map[int]int)
	for i := 0; i < len(nums) - 1; i++ {
		first := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if first + nums[j] >= 2020 {
				continue
			}
			diff := 2020 - first - nums[j]
			if _, ok := numsTable[diff]; ok {
				result = first * nums[j] * diff
				break
			}
			numsTable[first] = i
			numsTable[nums[j]] = j
		}
	}
	return result
}

func sum(nums []int) int {
	var result int
	for _, v := range nums {
		result += v
	}
	return result
}

func product(nums []int) int {
	result := 1
	for _, v := range nums {
		result *= v
	}
	return result
}

func recursive(nums []int, target int, parts []int, needed int) int {
	if len(parts) == needed {
		if target == 0 {
			return product(parts)
		}
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	return recursive(nums[1:], target - nums[0], append(parts, nums[0]), needed) |
		recursive(nums[1:], target, parts, needed)
}

func main() {
	nums, err := getInput("day1_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result := part1(nums)
	part2Result := part2(nums)

	recursiveResultPart1 := recursive(nums, 2020, make([]int, 0), 2)
	recursiveResultPart2 := recursive(nums, 2020, make([]int, 0), 3)

	fmt.Printf("Part 1 result: %d\n", part1Result)
	fmt.Printf("Part 2 result: %d\n", part2Result)
	fmt.Printf("Recursive Part 1 result: %d\n", recursiveResultPart1)
	fmt.Printf("Recursive Part 2 result: %d\n", recursiveResultPart2)
}
