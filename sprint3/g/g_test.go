package main

import (
	"reflect"
	"testing"
)

func TestWardrobeSort(t *testing.T) {
	tests := []struct{
		testNumber string
		unsorted []int
		sorted []int
	}{
		{"1", []int{0, 2, 1, 2, 0, 0, 1}, []int{0, 0, 0, 1, 1, 2, 2}},
		{"2", []int{2, 1, 2, 0, 1}, []int{0, 1, 1, 2, 2}},
		{"3", []int{2, 1, 1, 2, 0, 2}, []int{0, 1, 1, 2, 2, 2}},
		{"4", []int{2, 0, 2, 0, 1, 1}, []int{0, 0, 1, 1, 2, 2}},
		{"5", []int{0}, []int{0}},
		{"6", []int{2, 1}, []int{1, 2}},
		{"7", []int{1, 1, 2}, []int{1, 1, 2}},
		{"8", []int{0, 2, 0}, []int{0, 0, 2}},
		{"9", []int{0, 1, 0, 2, 0, 1, 2}, []int{0, 0, 0, 1, 1, 2, 2}},
		{"10", []int{2, 1, 2, 1, 1, 1, 1, 2, 2, 0, 2, 2, 1, 0, 0, 2, 0, 1, 2, 1, 2, 1, 1, 2, 1, 1, 0, 0, 1, 1, 0, 1, 2, 2, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}},
		{"11", []int{2, 2, 1, 1, 2, 2, 0, 0, 2, 0, 1, 1, 0, 2, 2, 0, 0, 2, 1, 1, 0, 2, 2, 1, 1, 0, 2, 0, 1, 0, 2, 0}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}},
		{"13", []int{0, 2, 1, 0, 0, 2, 0, 2}, []int{0, 0, 0, 0, 1, 2, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.testNumber, func(t *testing.T) {
			got := WardrobeSort(tt.unsorted)
			if !reflect.DeepEqual(got, tt.sorted) {
				t.Errorf("got %v want %v", got, tt.sorted)
			}
		})
	}
}