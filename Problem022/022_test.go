package main

import "testing"

func TestProblem022(t *testing.T) {
	actual := Problem022()
	expected := 871198282
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestCalcScore(t *testing.T) {
	actual := CalcScore("COLIN", 938)
	expected := 49714
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestSortStringArray(t *testing.T) {
	actual := SortStringArray([]string{"FUGA", "BINGO", "ZOO", "HOGE", "AI"})
	expected := []string{"AI", "BINGO", "FUGA", "HOGE", "ZOO"}
	if len(actual) != len(expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("got %v want %v", actual, expected)
		}
	}
}

func TestParseText(t *testing.T) {
	actual := ParseText("hoge,fuga,piyo,bar,foo")
	expected := []string{"hoge", "fuga", "piyo", "bar", "foo"}
	if len(actual) != len(expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("got %v want %v", actual, expected)
		}
	}
}

func TestLoadData(t *testing.T) {
	actual := LoadData("test.txt")
	expected := `"hogehoge"`
	if len(actual) != len(expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}
