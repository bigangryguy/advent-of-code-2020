package main

import (
	"testing"
)

func Test_getOpCodes(t *testing.T) {
	lines, err := getInput("day8_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := []OpCode {
		{
			Op: "nop",
			Value: 0,
		},
		{
			 Op: "acc",
			 Value: 1,
		},
		{
			Op: "jmp",
			Value: 4,
		},
		{
			Op: "acc",
			Value: 3,
		},
		{
			Op: "jmp",
			Value: -3,
		},
		{
			Op: "acc",
			Value: -99,
		},
		{
			Op: "acc",
			Value: 1,
		},
		{
			Op: "jmp",
			Value: -4,
		},
		{
			Op: "acc",
			Value: 6,
		},
	}
	actual, err := getOpCodes(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}
	if len(actual) != len(expected) {
		t.Error("getOpCodes: Actual length does not match expected length")
	}
	for i, actualOpCode := range actual {
		if actualOpCode.Op != expected[i].Op || actualOpCode.Value != expected[i].Value {
			t.Errorf("getOpCodes = %v at index %d for %v, expected %v", actualOpCode, i, lines[i], expected[i])
		}
	}
}

func Test_executeOpCode(t *testing.T) {
	opcodes := []OpCode {
		{
			Op: "acc",
			Value: 1,
		},
		{
			Op: "acc",
			Value: 3,
		},
		{
			Op: "acc",
			Value: -1,
		},
		{
			Op: "acc",
			Value: -9,
		},
		{
			Op: "nop",
			Value: 0,
		},
		{
			Op: "jmp",
			Value: 1,
		},
		{
			Op: "jmp",
			Value: 12,
		},
		{
			Op: "jmp",
			Value: -1,
		},
		{
			Op: "jmp",
			Value: -80,
		},
	}

	type OpCodeResults struct {
		Index int
		Accumulator int
	}
	expected := []OpCodeResults {
		{
			Index: 1,
			Accumulator: 1,
		},
		{
			Index: 2,
			Accumulator: 4,
		},
		{
			Index: 3,
			Accumulator: 3,
		},
		{
			Index: 4,
			Accumulator: -6,
		},
		{
			Index: 5,
			Accumulator: -6,
		},
		{
			Index: 6,
			Accumulator: -6,
		},
		{
			Index: 18,
			Accumulator: -6,
		},
		{
			Index: 6,
			Accumulator: -6,
		},
		{
			Index: -72,
			Accumulator: -6,
		},
	}
	var index, acc int
	for i, opcode := range opcodes {
		index, acc = executeOpCode(opcodes, i, acc)
		if index != expected[i].Index || acc != expected[i].Accumulator {
			t.Errorf("executeOpCode: Index = %d, Accumulator = %d for %v, expected %d, %d",
				index, acc, opcode, expected[i].Index, expected[i].Accumulator)
		}
	}
}

func areLinesOfCodeSame(loc1 LineOfCode, loc2 LineOfCode) bool {
	same := loc1.LineNbr == loc2.LineNbr &&
		loc1.Code.Op == loc2.Code.Op &&
		loc1.Code.Value == loc2.Code.Value &&
		loc1.Calls == loc2.Calls &&
		len(loc1.CalledBy) == len(loc2.CalledBy)
	if same {
		for i := 0; i < len(loc1.CalledBy); i++ {
			if loc1.CalledBy[i] != loc2.CalledBy[i] {
				same = false
				break
			}
		}
	}
	return same
}

func Test_getLinesOfCode(t *testing.T) {
	lines, err := getInput("day8_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}

	expected := []LineOfCode {
		{
			LineNbr: 0,
			Code: OpCode {
				Op: "nop",
				Value: 0,
			},
			Calls: 1,
			CalledBy: []int {},
		},
		{
			LineNbr: 1,
			Code: OpCode {
				Op: "acc",
				Value: 1,
			},
			Calls: 2,
			CalledBy: []int { 0, 4 },
		},
		{
			LineNbr: 2,
			Code: OpCode {
				Op: "jmp",
				Value: 4,
			},
			Calls: 6,
			CalledBy: []int { 1 },
		},
		{
			LineNbr: 3,
			Code: OpCode {
				Op: "acc",
				Value: 3,
			},
			Calls: 4,
			CalledBy: []int { 7 },
		},
		{
			LineNbr: 4,
			Code: OpCode {
				Op: "jmp",
				Value: -3,
			},
			Calls: 1,
			CalledBy: []int { 3 },
		},
		{
			LineNbr: 5,
			Code: OpCode {
				Op: "acc",
				Value: -99,
			},
			Calls: 6,
			CalledBy: []int {},
		},
		{
			LineNbr: 6,
			Code: OpCode {
				Op: "acc",
				Value: 1,
			},
			Calls: 7,
			CalledBy: []int { 2 },
		},
		{
			LineNbr: 7,
			Code: OpCode {
				Op: "jmp",
				Value: -4,
			},
			Calls: 3,
			CalledBy: []int { 6 },
		},
		{
			LineNbr: 8,
			Code: OpCode {
				Op: "acc",
				Value: 6,
			},
			Calls: 9,
			CalledBy: []int {},
		},
	}
	actual, err := getLinesOfCode(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}
	for i, actualLoc := range actual {
		if !areLinesOfCodeSame(actualLoc, expected[i]) {
			t.Errorf("getLinesOfCode = %v at index %d, expected %v\n", actualLoc, i, expected[i])
		}
	}
}

func Test_canReachLineFrom(t *testing.T) {
	lines, err := getInput("day8_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}
	locs, err := getLinesOfCode(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}

	type ReachTestValues struct {
		Expected bool
		Start int
		End int
	}

	testValues := []ReachTestValues {
		{true,1,7, },
		{ true, 6, 4, },
		{ false, 0, 8, },
		{ false, 3, 5, },
	}
	for _, testValue := range testValues {
		actual := canReachLineFrom(locs, testValue.Start, testValue.End)
		if actual != testValue.Expected {
			t.Errorf("canReachLineFrom = %v for start %d and end %d, expected %v\n",
				actual, testValue.Start, testValue.End, testValue.Expected)
		}
	}
}

func Test_part1(t *testing.T) {
	lines, err := getInput("day8_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}

	expected := 5
	actual, err := part1(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}
	if actual != expected {
		t.Errorf("part1 = %d, expected %d\n", actual, expected)
	}
}

//func Test_part2(t *testing.T) {
//	lines, err := getInput("day8_test_input.txt")
//	if err != nil {
//		t.Error("Received unexpected error")
//	}
//
//	expected := 8
//	actual, err := part2(lines)
//	if err != nil {
//		t.Error("Received unexpected error")
//	}
//	if actual != expected {
//		t.Errorf("part2 = %d, expected %d\n", actual, expected)
//	}
//}
