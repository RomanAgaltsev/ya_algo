package main

import (
	"reflect"
	"testing"
)

func TestGetBinaryNumber(t *testing.T) {
	binaryNumberTests := []struct{
		testNumber string
		number int
		binNumber []int
	}{
		{"1", 5, []int{1, 0, 1}},
		{"2", 14, []int{1, 1, 1, 0}},
		{"3", 0, []int{0}},
	}

	for _, tt := range binaryNumberTests {
			t.Run(tt.testNumber, func(t *testing.T) {
				got := GetBinaryNumber(tt.number)
				if !reflect.DeepEqual(got, tt.binNumber) {
					t.Errorf("%q got %v want %v", tt.number, got, tt.binNumber)
				}
			})
		}
}