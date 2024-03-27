package main

import (
	"reflect"
	"testing"
)

func TestGetNeighbours(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{0, 2, 6},
		{7, 4, 1},
		{2, 7, 0},
	}
	got := getNeighbours(matrix, 3, 0)
	want := []int{7, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
	
	matrix = [][]int{
		{1, 2, 3},
		{0, 2, 6},
		{7, 4, 1},
		{2, 7, 0},
	}
	got = getNeighbours(matrix, 0, 0)
	want = []int{0, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
	
	matrix = [][]int{
		{4, -10, 4, -9, 9, 5, -7, 1, 4, -3},
		{-3, 0, -1, -6, -6, 2, 3, 3, 4, 0},
		{-1, -5, 1, -9, -9, -6, 3, -1, -10, -7},
	}
	got = getNeighbours(matrix, 1, 0)
	want = []int{-1, 0, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
	matrix = [][]int{
		{3, 3, -9, 7, -5, 8, -6, -10, -4},
		{5, -2, -6, -9, 8, -4, 5, -5, 0},
		{-9, -3, 3, 2, 1, -4, -6, 3, -9},
		{-7, 1, -2, 4, -2, 1, -5, 4, -8},
		{-2, 5, 5, 7, -7, 2, 3, -4, -4},
		{-1, 7, -10, 7, 4, 5, -7, 1, 5},
		{-1, 3, 0, -8, -10, -2, 5, 1, 7},
		{10, 4, -9, 5, 3, -1, 7, 10, -5},
	}
	got = getNeighbours(matrix, 3, 0)
	want = []int{-9, -2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
}