package calc_test

import (
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
