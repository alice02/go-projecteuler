package main

import (
	"testing"
)

func TestProblem002(t *testing.T) {
	actual := Problem002()
	expected := 4613732
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
