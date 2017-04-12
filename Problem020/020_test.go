package main

import "testing"

func TestProblem020(t *testing.T) {
	actual := Problem020()
	expected := 648
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestDigitSum(t *testing.T) {
	actual := DigitSum("12345")
	expected := 15
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
