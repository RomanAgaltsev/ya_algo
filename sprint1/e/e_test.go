package main

import "testing"

func TestGetLongestWord(t *testing.T) {
	randomnessTests := []struct{
		testNumber string
		line string
		longestWord string
		longestWordLenght int
	}{
		{"1", "i love segment tree", "segment", 7},
		{"2", "frog jumps from river", "jumps", 5},
		{"3", "", "", 0},
		{"4", "     ", "", 0},
		{"5", "  g  ", "g", 1},
		{"6", " mymbg", "mymbg", 5},
		{"7", " wt", "wt", 2},
		{"8", " ligf kaakm", "kaakm", 5},
		{"9", "frog jump from rive", "frog", 4},
		{"10", " pfc", "pfc", 3},
		{"11", " ozmfudehppmivpkrmkvtnnzb ranigcyhpvooglyahlcjtusipmebtblqqquuweugqlkaakfdlcwzxmtcnzrcgyhsdbehy igrzfcsjbjeuhkkcbvbuvfdepdvnsiipkivrmzewqpryndfmvyhrzikzzddlqoovzmoikpwxirwn", "igrzfcsjbjeuhkkcbvbuvfdepdvnsiipkivrmzewqpryndfmvyhrzikzzddlqoovzmoikpwxirwn", 76},
		{"12", " izqcthct", "izqcthct", 8},
		
	}
	
	for _, tt := range randomnessTests {
		t.Run(tt.testNumber, func(t *testing.T) {
			got := getLongestWord(tt.line)
			if got != tt.longestWord || len(got) != tt.longestWordLenght {
				t.Errorf("%q got %q lenght %d want %q lenght %d", tt.line, got, len(got), tt.longestWord, tt.longestWordLenght)
			}
		})
	}
}