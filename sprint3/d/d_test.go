package main

import (
	"testing"
)

func TestWardrobeSort(t *testing.T) {
	tests := []struct{
		testNumber string
		kids []int
		cookies []int
		satisfied int
	}{
		{"1", []int{1, 2}, []int{2, 1, 3}, 2},
		{"2", []int{2, 1, 3}, []int{1, 1}, 1},
	
	}

	for _, tt := range tests {
		t.Run(tt.testNumber, func(t *testing.T) {
			got := FeedKids(tt.kids, tt.cookies)
			if got != tt.satisfied {
				t.Errorf("kids %v cookies %v got %v want %v", tt.kids, tt.cookies, got, tt.satisfied)
			}
		})
	}
}