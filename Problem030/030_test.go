package main

import (
	"reflect"
	"testing"
)

func TestProblem030(t *testing.T) {
	actual := Problem030()
	expected := 443839
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestSplitDigits(t *testing.T) {
	actual := SplitDigits(1234)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}
