package main

import "testing"

func TestProblem029(t *testing.T) {
	actual := Problem029()
	expected := 9183
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
