package main

import "testing"

func TestProblem014(t *testing.T) {
	actual := Problem014()
	expected := 837799
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestMakeSequence(t *testing.T) {
	actual := MakeSequence(13)
	expected := []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	if len(actual) != len(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
