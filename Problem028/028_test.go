package main

import "testing"

func TestProblem028(t *testing.T) {
	actual := Problem028()
	expected := 669171001
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
