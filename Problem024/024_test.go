package main

import "testing"

func TestProblem024(t *testing.T) {
	actual := Problem024()
	expected := 2783915460
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestElementToInt(t *testing.T) {
	actual := ElementsToInt([]int{0, 1, 2, 3, 4, 5})
	expected := 12345
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestRemoveElement(t *testing.T) {
	actual := RemoveElement([]int{1, 2, 3, 4, 5}, 0)
	expected := []int{2, 3, 4, 5}
	if len(actual) != len(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}

func TestPermutationSize(t *testing.T) {
	actual := PermutationSize(4)
	expected := 24 // 1 * 2 * 3 * 4
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
