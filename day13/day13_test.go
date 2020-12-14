package main

import (
	"testing"
)

func Test_getNextDeparture(t *testing.T) {
	timestamp := 939
	busIds := []int { 7, 13, 59, 31, 19 }
	expected := []int { 945, 949, 944, 961, 950 }
	for i, busId := range busIds {
		actual := getNextDeparture(timestamp, busId)
		if actual != expected[i] {
			t.Errorf("getNextDeparture = %d for busId = %d, expected %d\n",
				actual, busId, expected[i])
		}
	}
}

func Test_getEarliestNextDeparture(t *testing.T) {
	timestamp := 939
	busIds := []int { 7, 13, 59, 31, 19 }
	expectedBusId := 59
	expectedDeparture := 944
	actualBusId, actualDeparture := getEarliestNextDeparture(timestamp, busIds)
	if actualBusId != expectedBusId {
		t.Errorf("getEarliestNextDeparture: busId = %d, expected %d\n",
			actualBusId, expectedBusId)
	}
	if actualDeparture != expectedDeparture {
		t.Errorf("getEarliestNextDeparture: busId = %d, expected %d\n",
			actualDeparture, expectedDeparture)
	}
}

func Test_moduloInverse(t *testing.T) {
	testCases := [4][3]int {
		{ 35, 3, 2 },
		{ 21, 5, 1 },
		{ 15, 7, 1 },
		{ 3, 11, 4 },
	}
	for _, testCase := range testCases {
		actual := moduloInverse(testCase[0], testCase[1])
		if actual != testCase[2] {
			t.Errorf("moduloInverse = %d for %d, %d, expected %d",
				actual, testCase[0], testCase[1], testCase[2])
		}
	}
}

func Test_crt(t *testing.T) {
	nums := []int { 17, 13, 19 }
	remainders := []int { 0, 11, 16 }
	expected := 3417
	actual := crt(nums, remainders)
	if actual != expected {
		t.Errorf("crt = %d for %v, %v, expected %d",
			actual, nums, remainders, expected)
	}
}

func Test_part1(t *testing.T) {
	timestamp := 939
	busIds := []int { 7, 13, 59, 31, 19 }
	expected := 295
	actual := part1(timestamp, busIds)
	if actual != expected {
		t.Errorf("part1 = %d, expected %d\n", actual, expected)
	}
}

func Test_part2(t *testing.T) {
	busIds := []int { 7, 13, -1, -1, 59, -1, 31, 19 }
	expected := 1068781
	actual := part2(busIds)
	if actual != expected {
		t.Errorf("part2 = %d, expected %d\n", actual, expected)
	}
}
