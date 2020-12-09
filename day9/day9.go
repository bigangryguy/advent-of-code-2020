package main

import (
	"fmt"
	"io/ioutil"
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
	return
}

func canSum(nums []int, target int) bool {
	numsTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if _, ok := numsTable[diff]; ok {
			return true
		}
		numsTable[nums[i]] = i
	}
	return false
}

func findInvalidNumber(nums []int, preamble int) int {
	for i := 0; i < len(nums); i++ {
		if !canSum(nums[i:i+preamble], nums[i+preamble]) {
			return nums[i+preamble]
		}
	}
	return 0
}

func findContiguousSum(nums []int, target int) (result []int) {
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i+1; j < len(nums); j++ {
			if sum + nums[j] > target {
				break
			} else if sum + nums[j] == target {
				return nums[i:j+1]
			} else {
				sum += nums[j]
			}
		}
	}
	return nil
}

func part1(nums []int) int {
	return findInvalidNumber(nums, 25)
}

func part2(nums []int, target int) int {
	contiguous := findContiguousSum(nums, target)
	min := contiguous[0]
	max := contiguous[0]
	for i := 1; i < len(contiguous); i++ {
		if contiguous[i] < min {
			min = contiguous[i]
		} else if contiguous[i] > max {
			max = contiguous[i]
		}
	}
	return min + max
}

func main() {
	nums, err := getInput("day9_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result := part1(nums)
	part2Result := part2(nums, part1Result)

	fmt.Printf("part 1 answer: %d\n", part1Result)
	fmt.Printf("part 2 answer: %d\n", part2Result)
}
