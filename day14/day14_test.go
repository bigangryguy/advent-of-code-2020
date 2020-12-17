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

func Test_combinations(t *testing.T) {
	valuesList := [2][]int64 {
		{ 32, 1 },
		{ 8, 2, 1 },
	}
	expected := [2][]int64 {
		{ 0, 1, 32, 33 },
		{ 0, 1, 2, 3, 8, 9, 10, 11 },
	}
	for i, values := range valuesList {
		actual := combinations(values)
		if len(actual) != len(expected[i]) {
			t.Errorf("Length of actual does not match length of expected\nactual = %v, expected = %v\n",
				actual, expected[i])
		}
		for j, actualValue := range actual {
			if actualValue != expected[i][j] {
				t.Errorf("combinations: actual = %d at index %d, expected %d\n",
					actualValue, j, expected[i][j])
			}
		}
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

func Test_floatingBitsFromBitmask(t *testing.T) {
	bitmask := "000000000000000000000000000000X1001X"
	expected := []int64 { 0, 1, 32, 33 }
	actual := floatingBitsFromBitmask(bitmask)
	if len(actual) != len(expected) {
		t.Fatalf("Length mismatch in floatingBitsFromBitmask. actual = %v, expected = %v\n",
			actual, expected)
	}
	for i, actualValue := range actual {
		if actualValue != expected[i] {
			t.Errorf("floatingBitsFromBitmask = %d at index %d, expected %d\n",
				actualValue, i, expected[i])
		}
	}
}

func Test_fixedBitsFromBitmask(t *testing.T) {
	bitmask := "000000000000000000000000000000X1001X"
	var expected int64 = 18
	actual := fixedBitsFromBitmask(bitmask)
	if actual != expected {
		t.Errorf("fixedBitsFromBitmask = %d, expected %d\n", actual, expected)
	}
}

func Test_applyFloatingMask(t *testing.T) {
	bitmasks := []string {
		"000000000000000000000000000000X1001X",
		"00000000000000000000000000000000X0XX",
	}
	inputs := []int64 { 42, 26 }
	expected := [2][]int64 {
		{ 26, 27, 58, 59 },
		{ 16, 17, 18, 19, 24, 25, 26, 27 },
	}
	for i := 0; i < 2; i++ {
		actual := applyFloatingMask(inputs[i], bitmasks[i])
		if len(actual) != len(expected[i]) {
			t.Fatalf("Length mismatch in applyFloatingMask. actual = %v, expected = %v\n",
				actual, expected[i])
		}
		for j, actualValue := range actual {
			if actualValue != expected[i][j] {
				t.Errorf("applyFloatingMask = %d at index %d, expected %d\n",
					actualValue, j, expected[i][j])
			}
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

func Test_part2(t *testing.T) {
	lines, err := getInput("day14_test_input2.txt")
	if err != nil {
		t.Errorf("Error getting input: %v\n", err)
	}

	var expected int64 = 208
	actual, err := part2(lines)
	if err != nil {
		t.Errorf("Received unexpected error: %v\n", err)
	}
	if actual != expected {
		t.Errorf("part2 = %d, expected %d\n", actual, expected)
	}
}
