package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type OpCode struct {
	Op string
	Value int
}

func getInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file: %q", err))
	}

	return strings.Split(string(data), "\n"), nil
}

func getOpCodes(lines []string) (opcodes []OpCode, err error) {
	var value int
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		value, err = strconv.Atoi(tokens[1])
		if err != nil {
			return
		}
		opcodes = append(opcodes, OpCode{ tokens[0], value })
	}
	return
}

func executeOpCode(opcodes []OpCode, indexToExecute int, currentAcc int) (nextIndex int, nextAcc int) {
	opcode := opcodes[indexToExecute]
	switch opcode.Op {
	case "acc":
		nextIndex = indexToExecute + 1
		nextAcc = currentAcc + opcode.Value
	case "jmp":
		nextIndex = indexToExecute + opcode.Value
		nextAcc = currentAcc
	case "nop":
		nextIndex = indexToExecute + 1
		nextAcc = currentAcc
	}
	return
}

func part1(lines []string) (result int, err error) {
	opcodes, err := getOpCodes(lines)
	if err != nil {
		return
	}

	alreadyRun := map[int]int { 0: 1 }
	for index := 0; index < len(opcodes); {
		index, result = executeOpCode(opcodes, index, result)
		if _, found := alreadyRun[index]; found {
			break
		}
		alreadyRun[index]++
	}
	return
}

func part2(lines []string) (result int, err error) {
	opcodes, err := getOpCodes(lines)
	if err != nil {
		return
	}

	// Seriously, trial and error?
	// Ah, duh, figured out the graph-based solution after giving in and doing this.
	// Will implement after completing day 9.
	for i := 0; i < len(opcodes); i++ {
		oldOp := opcodes[i].Op
		if opcodes[i].Op == "jmp" {
			opcodes[i].Op = "nop"
		} else if opcodes[i].Op == "nop" {
			opcodes[i].Op = "jmp"
		} else {
			continue
		}

		success := true
		result = 0
		alreadyRun := map[int]int{0: 1}
		for index := 0; index < len(opcodes); {
			index, result = executeOpCode(opcodes, index, result)
			if _, found := alreadyRun[index]; found {
				success = false
				break
			}
			alreadyRun[index]++
		}
		if success {
			break
		}
		opcodes[i].Op = oldOp
	}
	return
}

func main() {
	lines, err := getInput("day8_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result, err := part1(lines)
	if err != nil {
		fmt.Printf("Error getting part 1 answer: %v\n", err)
	}
	fmt.Printf("Part 1 answer: %d\n", part1Result)

	part2Result, err := part2(lines)
	if err != nil {
		fmt.Printf("Error getting part 2 answer: %v\n", err)
	}
	fmt.Printf("Part 2 answer: %d\n", part2Result)
}
