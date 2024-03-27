package main

import "testing"

func TestGetWeatherRandomness(t *testing.T) {
	randomnessTests := []struct{
		number string
		temperatures []int
		randomness int
	}{
		{"1",[]int{-1, -10, -8, 0, 2, 0, 5}, 3},
		{"2",[]int{1, 2, 5, 4, 8}, 2},
		{"3", []int{0, -1, -1, 10, 0, 0}, 2},
		{"4", []int{0}, 1},
		{"5", []int{0, 0, 0, 0, 0, 0, 0}, 0},
		{"6", []int{-20, -15, -16, -16, -8, -1}, 2},
	}
	
	for _, tt := range randomnessTests {
		t.Run(tt.number, func(t *testing.T) {
			got := getWeatherRandomness(tt.temperatures)
			if got != tt.randomness {
				t.Errorf("%#v got %d want %d", tt.temperatures, got, tt.randomness)
			}
		})
	}
}