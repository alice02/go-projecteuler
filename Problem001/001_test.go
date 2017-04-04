package main

import (
	"testing"
)

func TestProblem001(t *testing.T) {
	actual := Problem001()
	expected := 233168
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
