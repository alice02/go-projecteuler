package main

import "testing"

func TestProblem030(t *testing.T) {
	actual := Problem030()
	expected := 443839
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
