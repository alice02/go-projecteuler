package main

import (
	"testing"
)

func TestProblem006(t *testing.T) {
	actual := Problem006()
	expected := 25164150
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
