package main

import "testing"

func TestProblem017(t *testing.T) {
	actual := Problem017()
	expected := 21124
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCalcNumberLength(t *testing.T) {
	var actual, expected int
	actual = CalcNumberLength(1)
	expected = 3 // one
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	actual = CalcNumberLength(21)
	expected = 9 // twenty one
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	actual = CalcNumberLength(112)
	expected = 19 // one hundred and twelve
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
