package main

import "testing"

func TestProblem032(t *testing.T) {
	actual := Problem032()
	expected := 45228
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
