package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		got := []int{1, 4, 2, 10, 1, 2}
		merge_sort(got, 0, 6)
		want := []int{1, 1, 2, 2, 4, 10}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		got := []int{-6, -12, -14, 14}
		merge_sort(got, 0, 4)
		want := []int{-14, -12, -6, 14}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		got := []int{17, 7}
		merge_sort(got, 0, 2)
		want := []int{7, 17}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("6", func(t *testing.T) {
		got := []int{18, -19, 15, -8, 14, 6, -6, 8, 17}
		merge_sort(got, 0, 9)
		want := []int{-19, -8, -6, 6, 8, 14, 15, 17, 18}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("8", func(t *testing.T) {
		got := []int{71, 78, -76, 16, 53, 27, 55, -96, 31, 10, 47, -5, 59, 40, 93, 11, 0, -55, -86, -73, -69, -4, -61, -76, -59, 31, 26, -68, 97, -1, 52, 29, -31, 47, -54, -4, -84, -88, 48, -85, -53, -85, -58, -25, -30, -81, 33, 59}
		merge_sort(got, 0, 48)
		want := []int{-96, -88, -86, -85, -85, -84, -81, -76, -76, -73, -69, -68, -61, -59, -58, -55, -54, -53, -31, -30, -25, -5, -4, -4, -1, 0, 10, 11, 16, 26, 27, 29, 31, 31, 33, 40, 47, 47, 48, 52, 53, 55, 59, 59, 71, 78, 93, 97}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("9", func(t *testing.T) {
		got := []int{49, 67, 66, 73, -85, -76, 59, 17, 65, 85, -16, -16, 35, 70}
		merge_sort(got, 0, 14)
		want := []int{-85, -76, -16, -16, 17, 35, 49, 59, 65, 66, 67, 70, 73, 85}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("11", func(t *testing.T) {
		got := []int{80, 88, -87, 71, -41, -24, -44, -12, -52, 98, -45, -2, -1, -43, -85, 17, 28, -85, -76, -94, -94, 39, 56, 30, 28}
		merge_sort(got, 0, 25)
		want := []int{-94, -94, -87, -85, -85, -76, -52, -45, -44, -43, -41, -24, -12, -2, -1, 17, 28, 28, 30, 39, 56, 71, 80, 88, 98}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
