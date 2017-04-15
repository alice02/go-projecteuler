package main

import "testing"

func TestProblem026(t *testing.T) {
	actual := Problem026()
	expected := 983
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestCalcRepeatendLength(t *testing.T) {
	var actual, expected int
	actual = CalcRepeatendLength(7)
	expected = 6
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
	actual = CalcRepeatendLength(6)
	expected = 1
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
