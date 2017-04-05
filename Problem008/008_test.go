package main

import (
	"testing"
)

func TestProblem008(t *testing.T) {
	actual := Problem008()
	expected := 23514624000
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestParseText(t *testing.T) {
	text := `12345
67890`
	actual := ParseText(text)
	expected := "1234567890"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
