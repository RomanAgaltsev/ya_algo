package main

import (
	"reflect"
	"testing"
)

func TestGetNearestZeros(t *testing.T) {
	nearestZerosTests := []struct{
		testNumber string
		houseNumbers []int
		nearetsZeros []int
	}{
		{"1", []int{0, 1, 4, 9, 0}, []int{0, 1, 2, 1, 0}},
		{"2", []int{0, 7, 9, 4, 8, 20}, []int{0, 1, 2, 3, 4, 5}},
		{"3", []int{98, 0, 10, 77, 0, 59, 28, 0, 94}, []int{1, 0, 1, 1, 0, 1, 1, 0, 1}},
		{"4", []int{99, 0, 100, 72, 43, 49, 0, 51, 19, 61, 93, 31}, []int{1, 0, 1, 2, 2, 1, 0, 1, 2, 3, 4, 5}},
		{"5", []int{64, 68, 37, 11, 77, 80, 48, 82, 0}, []int{8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{"6", []int{0, 3, 41, 0, 0, 0, 0, 0, 49, 0, 0, 56, 0, 88}, []int{0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1}},
		{"7", []int{0, 0, 20, 0, 0, 0, 0, 40, 0, 0, 65, 73, 77, 0, 79, 0, 82, 0, 0, 0}, []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 2, 1, 0, 1, 0, 1, 0, 0, 0}},
		{"8", []int{5, 8, 9, 12, 15, 26, 30, 0, 0, 55, 0, 0, 67, 0, 76, 80, 82, 0, 0, 98}, []int{7, 6, 5, 4, 3, 2, 1, 0, 0, 1, 0, 0, 1, 0, 1, 2, 1, 0, 0, 1}},
		{"9", []int{10, 13, 31, 35, 39, 0, 0, 59, 0, 66, 68, 73, 74, 0, 0, 0, 87, 89, 96, 99}, []int{5, 4, 3, 2, 1, 0, 0, 1, 0, 1, 2, 2, 1, 0, 0, 0, 1, 2, 3, 4}},
		{"10", []int{3, 15, 0, 22, 31, 32, 0, 41, 0, 0, 50, 0, 0, 66, 0, 76, 77, 82, 0, 89}, []int{2, 1, 0, 1, 2, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 2, 1, 0, 1}},
		{"27", []int{0}, []int{0}},
		{"28", []int{0, 1}, []int{0, 1}},
		{"29", []int{1, 0}, []int{1, 0}},
		{"31", []int{1, 2, 3, 4, 0, 6, 7, 8, 9, 10}, []int{4, 3, 2, 1, 0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range nearestZerosTests {
		t.Run(tt.testNumber, func(t *testing.T) {
			got := GetNearestZeros(&tt.houseNumbers)
			if !reflect.DeepEqual(got, tt.nearetsZeros){
				t.Errorf("%#v got %v want %v", tt.houseNumbers, got, tt.nearetsZeros)
			}
		})
	}
}