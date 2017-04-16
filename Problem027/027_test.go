package main

import "testing"

func TestProblem027(t *testing.T) {
	actual := Problem027()
	expected := -59231
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestIsPrime(t *testing.T) {
	primes := EratosthenesSieve(100)
	actual := IsPrime(53, primes)
	expected := true
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
