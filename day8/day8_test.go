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

func Test_part2(t *testing.T) {
	lines, err := getInput("day8_test_input.txt")
	if err != nil {
		t.Error("Received unexpected error")
	}

	expected := 8
	actual, err := part2(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}
	if actual != expected {
		t.Errorf("part2 = %d, expected %d\n", actual, expected)
	}
}
