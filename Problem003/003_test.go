package main

import (
	"testing"
)

func TestProblem003(t *testing.T) {
	actual := Problem003()
	expected := 6857
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
