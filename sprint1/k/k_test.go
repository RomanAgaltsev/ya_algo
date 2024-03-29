package main

import (
	"reflect"
	"testing"
)

func TestGetSum(t *testing.T) {
	sumTests := []struct{
		testNumber string
		bigNumber []int
		smallNumber int
		sum []int
	}{
		{"1", []int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
		{"2", []int{9, 5}, 17, []int{1, 1, 2}},
		{"4", []int{3}, 7991, []int{7, 9, 9, 4}},
		{"5", []int{3}, 3003, []int{3, 0, 0, 6}},
		{"6", []int{8}, 4311, []int{4, 3, 1, 9}},
		{"7", []int{5, 6}, 7285, []int{7, 3, 4, 1}},
		{"9", []int{2, 1}, 6636, []int{6, 6, 5, 7}},
		{"10", []int{7, 3}, 8516, []int{8, 5, 8, 9}},
		{"11", []int{5, 7}, 6528, []int{6, 5, 8, 5}},
		{"12", []int{1}, 3082, []int{3, 0, 8, 3}},
		{"13", []int{4, 9, 6}, 1289, []int{1, 7, 8, 5}},
		{"14", []int{6, 8}, 6632, []int{6, 7, 0, 0}},
		{"16", []int{3}, 2550, []int{2, 5, 5, 3}},
		{"17", []int{4, 5}, 3765, []int{3, 8, 1, 0}},
		{"18", []int{6, 1}, 4830, []int{4, 8, 9, 1}},
		{"19", []int{1, 8, 7}, 4403, []int{4, 5, 9, 0}},
		{"20", []int{6, 3, 2}, 3173, []int{3, 8, 0, 5}},
	}

		for _, tt := range sumTests {
			t.Run(tt.testNumber, func(t *testing.T) {
				got := GetSum(tt.bigNumber, tt.smallNumber)
				if !reflect.DeepEqual(got, tt.sum) {
					t.Errorf("bigNumber %#v smallNumber %v got %v want %v", tt.bigNumber, tt.smallNumber, got, tt.sum)
				}
			})
		}
}