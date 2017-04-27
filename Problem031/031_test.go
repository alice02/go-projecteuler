package main

import "testing"

func TestProblem031(t *testing.T) {
	actual := Problem031()
	expected := 73682
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
