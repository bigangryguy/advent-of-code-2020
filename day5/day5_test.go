package main

import (
	"fmt"
	"testing"
)

func Test_parseRow(t *testing.T) {
	var actual int
	var err error

	rows := []string { "BFFFBBF", "FFFBBBF", "BBFFBBF" }
	expected := []int { 70, 14, 102 }
	for i, row := range rows {
		actual, err = parseRow(row)
		if err != nil {
			t.Error("Received unexpected error")
		}
		if actual != expected[i] {
			t.Errorf("parseRow = %v for %v, expected %v", actual, row, expected[i])
		}
	}
}

func Test_parseColumn(t *testing.T) {
	var actual int
	var err error

	columns := []string { "RRR", "RLL" }
	expected := []int { 7, 4 }
	for i, column := range columns {
		actual, err = parseColumn(column)
		if err != nil {
			t.Error("Received unexpected error")
		}
		if actual != expected[i] {
			t.Errorf("parseColumn = %v for %v, expected %v", actual, column, expected[i])
		}
	}
}

func Test_getSeatID(t *testing.T) {
	var actual int
	var err error

	lines := []string { "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL" }
	expected := []int { 567, 119, 820 }
	for i, line := range lines {
		actual, err = getSeatID(line)
		if err != nil {
			t.Error("Received unexpected error")
		}
		if actual != expected[i] {
			t.Errorf("getSeatID = %v for %v, expected %v", actual, line, expected[i])
		}
	}
}

func Test_getSeatList(t *testing.T) {
	var actual []int
	var err error

	lines := []string { "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL" }
	expected := []int { 567, 119, 820 }
	actual, err = getSeatList(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}
	for i, seatID := range actual {
		if seatID != expected[i] {
			t.Errorf("getSeatList = %v for %v at index %v, expected %v", seatID, lines[i], i, expected[i])
		}
	}
}

func Test_part1(t *testing.T) {
	lines, err := getInput("day5_test_input.txt")
	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	actual, err := part1(lines)
	if err != nil {
		t.Error("Received unexpected error")
	}
	expected := 820
	if actual != expected {
		t.Errorf("part1 = %v, expected %v", actual, expected)
	}
}
