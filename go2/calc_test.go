package calc_test

import (
	calc "calctest"
	"calctest/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	got := calc.Add(2, 1)
	expected := 3

	if got != expected {
		t.Errorf("not expected result: %v expected: %v", got, expected)
	}
}

func TestSub(t *testing.T) {
	got := calc.Sub(2, 1)
	expected := 1

	if got != expected {
		t.Errorf("not expected result, got: %v expected: %v", got, expected)
	}
}

func TestSubTable(t *testing.T) {
	data := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 2, y: 1, expected: 1},
		{x: 2, y: 2, expected: 0},
		{x: 2, y: 3, expected: -1},
	}

	for _, val := range data {
		got := calc.Sub(val.x, val.y)
	}
}
