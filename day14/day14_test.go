package main

import (
	"testing"
)

func Test_generateMasks(t *testing.T) {
	bitmask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	var expectedOrMask int64 = 64
	var expectedAndMask int64 = 68719476733
	actualOrMask, actualAndMask := generateMasks(bitmask)
	if actualOrMask != expectedOrMask {
		t.Errorf("generateMasks: orMask = %d, expected %d\n", actualOrMask, expectedOrMask)
	}
	if actualAndMask != expectedAndMask {
		t.Errorf("generateMasks: andMask = %d, expected %d\n", actualAndMask, expectedAndMask)
	}
}

func Test_applyMasks(t *testing.T) {
	var orMask int64 = 64
	var andMask int64 = 68719476733
	testCases := [3][2]int64 {
		{ 11, 73 },
		{ 101, 101 },
		{ 0, 64 },
	}
	for _, testCase := range testCases {
		actual := applyMasks(testCase[0], orMask, andMask)
		if actual != testCase[1] {
			t.Errorf("applyMasks = %d for %d, expected %d", actual, testCase[0], testCase[1])
		}
	}
}

func Test_part1(t *testing.T) {
	lines, err := getInput("day14_test_input.txt")
	if err != nil {
		t.Errorf("Error getting input: %v\n", err)
	}

	var expected int64 = 165
	actual, err := part1(lines)
	if err != nil {
		t.Errorf("Received unexpected error: %v\n", err)
	}
	if actual != expected {
		t.Errorf("part1 = %d, expected %d\n", actual, expected)
	}
}
