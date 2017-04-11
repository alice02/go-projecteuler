package main

import "testing"

func TestProblem019(t *testing.T) {
	actual := Problem019()
	expected := 171
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
