package main

import "testing"

func TestProblem016(t *testing.T) {
	actual := Problem016()
	expected := 1366
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDigitSum(t *testing.T) {
	actual := DigitSum("123456789")
	expected := 45 // sum of 1,2,...,9
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
