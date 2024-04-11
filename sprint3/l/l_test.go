package main

import (
	"reflect"
	"testing"
)

func TestBikesSearch(t *testing.T) {
	bikesTests := []struct {
		testNumber string
		savings    []int
		price      int
		days       []int
	}{
		{"1", []int{1, 2, 4, 4, 6, 8}, 3, []int{3, 5}},
		{"2", []int{1, 2, 4, 4, 4, 4}, 3, []int{3, -1}},
		{"3", []int{1, 2, 4, 4, 4, 4}, 10, []int{-1, -1}},
	}
	for _, tt := range bikesTests {
		t.Run(tt.testNumber, func(t *testing.T) {
			got := getDays(tt.savings, tt.price)
			if !reflect.DeepEqual(got, tt.days) {
				t.Errorf("%#v got %v want %v", tt.savings, got, tt.days)
			}
		})
	}
}
