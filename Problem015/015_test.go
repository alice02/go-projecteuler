package main

import "testing"

func TestProblem015(t *testing.T) {
	actual := Problem015()
	expected := 137846528820
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
