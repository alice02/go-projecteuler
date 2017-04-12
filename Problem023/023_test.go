package main

import "testing"

func TestProblem023(t *testing.T) {
	actual := Problem023()
	expected := 4179871
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestCalcActualDivisor(t *testing.T) {
	actual := CalcActualDivisor(12)
	expected := []int{1, 2, 3, 4, 6}
	if len(actual) != len(expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("got %v want %v", actual, expected)
		}
	}
}

func TestCalcArraySum(t *testing.T) {
	actual := CalcArraySum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := 55
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestCheckAbundantNumber(t *testing.T) {
	actual := CheckAbundantNumber(12)
	expected := true
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
