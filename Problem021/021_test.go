package main

import "testing"

func TestProblem021(t *testing.T) {
	actual := Problem021()
	expected := 31626
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestCalcArraySum(t *testing.T) {
	actual := CalcArraySum([]int{1, 2, 3, 4, 5})
	expected := 15
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestCalcActualDivisor(t *testing.T) {
	actual := CalcActualDivisor(10)
	expected := []int{1, 2, 5}
	if len(actual) != len(expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("got %v want %v", actual, expected)
		}
	}
}
