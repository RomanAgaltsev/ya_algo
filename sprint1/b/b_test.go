package main

import "testing"

func TestCheckParity(t *testing.T) {
	got := checkParity(1, 2, -3)
	if got != false {
		t.Errorf("want %v but got %v", false, got)
	}
	
	got = checkParity(7, 11, 7)
	if got != true {
		t.Errorf("want %v but got %v", true, got)
	}
	
	got = checkParity(6, -2, 0)
	if got != true {
		t.Errorf("want %v but got %v", true, got)
	}
}