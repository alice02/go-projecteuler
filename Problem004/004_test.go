package main

import (
	"testing"
)

func TestProblem004(t *testing.T) {
	actual := Problem004()
	expected := 906609
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsPalindrome(t *testing.T) {
	var actual, expected bool
	palindromeNum := 245542
	actual = IsPalindrome(palindromeNum)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	notPalindromeNum := 245541
	actual = IsPalindrome(notPalindromeNum)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
