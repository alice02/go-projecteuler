package main

import "testing"

func TestProblem012(t *testing.T) {
	actual := Problem012()
	expected := 76576500
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCalcDivisorLength(t *testing.T) {
	actual := CalcDivisorLength(12)
	expected := 5 // 1, 2, 3, 4, 6
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestGenerateTriangularNumber(t *testing.T) {
	actual := GenerateTriangularNumber(10)
	expected := 55 // sum of 1,2,...,10
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
