package main

import (
	"testing"
)

func TestProblem009(t *testing.T) {
	actual := Problem009()
	expected := 31875000
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
