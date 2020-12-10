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

type LineOfCode struct {
	LineNbr int
	Code OpCode
	Calls int
	CalledBy []int
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

func getLinesOfCode(lines []string) (locs []LineOfCode, err error) {
	var value int
	for i, line := range lines {
		tokens := strings.Split(line, " ")
		value, err = strconv.Atoi(tokens[1])
		if err != nil {
			return
		}
		op := tokens[0]
		nextLine := i + 1
		if op == "jmp" {
			nextLine = i + value
		}
		locs = append(locs, LineOfCode{
			LineNbr: i,
			Code: OpCode{ op, value },
			Calls: nextLine,
			CalledBy: []int {},
		})
	}

	// Set callers after all lines are added
	for _, loc := range locs {
		if loc.Calls < len(locs) {
			locs[loc.Calls].CalledBy = append(locs[loc.Calls].CalledBy, loc.LineNbr)
		}
	}

	// Eliminate dead code callers after all lines have callers set
	// Since index == line number, no need for map
	for i := 0; i < len(locs); i++ {
		var revisedCalledBy []int
		for _, caller := range locs[i].CalledBy {
			if caller == 0 || len(locs[caller].CalledBy) > 0 {
				revisedCalledBy = append(revisedCalledBy, caller)
			}
		}
		locs[i].CalledBy = revisedCalledBy
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

func canReachLineFrom(locs []LineOfCode, start int, end int) bool {
	alreadyChecked := make(map[int]int)
	for index := start; index < len(locs); {
		if _, found := alreadyChecked[index]; found {
			break
		}
		alreadyChecked[index]++
		if locs[index].Calls == end {
			return true
		} else {
			index = locs[index].Calls
		}
	}
	return false
}

func part1(lines []string) (result int, err error) {
	var opcodes []OpCode
	opcodes, err = getOpCodes(lines)
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

//func part2(lines []string) (result int, err error) {
//	var locs []LineOfCode
//	locs, err = getLinesOfCode(lines)
//	if err != nil {
//		return
//	}
//
//	// Find start of loop
//	alreadyRun := map[int]int { 0 : 1 }
//	for index := 0; index < len(locs); {
//		index, acc := executeOpCode()
//	}
//}

func main() {
	lines, err := getInput("day8_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	part1Result, err := part1(lines)
	if err != nil {
		fmt.Printf("Error getting part 1 answer: %v\n", err)
	}
	fmt.Printf("Part 1 answer: %d\n", part1Result)

	locs, _ := getLinesOfCode(lines)
	fmt.Printf("Lines of code: %v\n", locs)

	//part2Result, err := part2(lines)
	//if err != nil {
	//	fmt.Printf("Error getting part 2 answer: %v\n", err)
	//}
	//fmt.Printf("Part 2 answer: %d\n", part2Result)
}
