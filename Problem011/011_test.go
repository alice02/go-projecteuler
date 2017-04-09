package main

import "testing"

func TestProblem011(t *testing.T) {
	actual := Problem011()
	expected := 70600674
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestParseMatrix(t *testing.T) {
	matrix := `01 02 03 04 05
06 07 08 09 10
11 12 13 14 15`
	x_range = 5
	y_range = 3
	actual := ParseMatrix(matrix, x_range, y_range)
	expected := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}

	if len(actual) != len(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	for y := range actual {
		if len(actual[y]) != len(expected[y]) {
			t.Errorf("got %v\nwant %v", actual[y], expected[y])
		}
		for x := range actual[y] {
			if actual[y][x] != expected[y][x] {
				t.Errorf("got %v\nwant %v", actual[y][x], expected[y][x])
			}
		}
	}
}
