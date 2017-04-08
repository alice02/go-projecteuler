package main

import (
	"testing"
)

func TestProblem010(t *testing.T) {
	actual := Problem010()
	expected := 142913828922
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCalculateArraySum(t *testing.T) {
	testArray := []int{1, 2, 3, 4, 5}
	actual := CalculateArraySum(testArray)
	expected := 15
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestEratosthenesSieve(t *testing.T) {
	actual := EratosthenesSieve(10)
	expected := []int{2, 3, 5, 7}
	if len(actual) != len(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}

func TestInitialFlag(t *testing.T) {
	actual := InitializeFlag(5)
	expected := []bool{false, false, true, true, true}
	if len(actual) != len(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
