package main

import "testing"

func TestEvaluateFunction(t *testing.T) {
	got := evaluateFunction(-8, -2, 7, -5)
	want := -183
	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}

	got = evaluateFunction(8, 9, -10, 2)
	want = 40
	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}
	
	got = evaluateFunction(0, 9, -10, 3)
	want = 17
	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}
	
	got = evaluateFunction(7, 0, 5, 4)
	want = 117
	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}
	
	got = evaluateFunction(7, -5, 0, -2)
	want = 38
	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}
	
	got = evaluateFunction(-7, -5, 4, 0)
	want = 4
	if got != want {
		t.Errorf("want %d but got %d", want, got)
	}
}
