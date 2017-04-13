package main

import (
	"math/big"
	"testing"
)

func TestProblem025(t *testing.T) {
	actual := Problem025()
	expected := 4782
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for _, e := range expected {
		act := f()
		expect := big.NewInt(int64(e))
		if act.String() != expect.String() {
			t.Errorf("got %v want %v", act, expect)
		}
	}

}
