package main

import (
	"testing"
)

func TestProblem007(t *testing.T) {
	actual := Problem007()
	expected := 104743
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsPrime(t *testing.T) {
	actual := IsPrime(5)
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsPrime(9)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
