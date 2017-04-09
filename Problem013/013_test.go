package main

import "testing"

func TestProblem013(t *testing.T) {
	actual := Problem013()
	expected := 5537376230
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
