package main

import (
	"testing"
)

func Test_vanEck(t *testing.T) {
	expected := []int { 0, 0, 1, 0, 2, 0, 2, 2, 1, 6, 0, 5, 0, 2, 6, 5, 4, 0, 5, 3 }
	actual := vanEck([]int{ 0, 0 }, 20)
	if len(actual) != len(expected) {
		t.Fatalf("Length of actual does not match length of expected\nactual = %v, expected = %v\n",
			actual, expected)
	}
	for i, actualValue := range actual {
		if actualValue != expected[i] {
			t.Errorf("vanEck: actual = %d at index %d, expected %d\n",
				actualValue, i, expected[i])
		}
	}
}
