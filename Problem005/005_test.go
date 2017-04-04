package main

import (
	"testing"
)

func TestProblem005(t *testing.T) {
	actual := Problem005()
	expected := 232792560
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestLcm(t *testing.T) {
	actual := Lcm(3, 4)
	expected := 12
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestGcd(t *testing.T) {
	actual := Gcd(12, 18)
	expected := 6
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
