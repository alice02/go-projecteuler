package main

import "testing"

func TestProblem018(t *testing.T) {
	actual := Problem018()
	expected := 1074
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestMax(t *testing.T) {
	actual := Max(10, 11)
	expected := 11
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestStringToIntArray(t *testing.T) {
	actual := StringToIntArray([]string{"1", "2", "3", "4", "5"})
	expected := []int{1, 2, 3, 4, 5}
	if len(actual) != len(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
