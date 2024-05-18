package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func getRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func BenchmarkHeapsort(b *testing.B) {
	RunBenchmarkHeapsort(b, true)
	RunBenchmarkHeapsort(b, false)
}

func RunBenchmarkHeapsort(b *testing.B, iterative bool) {
	for _, size := range []int{10_000, 100_000, 1_000_000} {
		b.Run(fmt.Sprintf("Size=%d/Iterative=%v/", size, iterative), func(b *testing.B) {
			participants := make([]Participant, size)
			for i := range participants {
				participants[i] = Participant{
					getRandomString(20),
					rand.Intn(1_000_000_000),
					rand.Intn(1_000_000_000),
				}
			}
			b.ResetTimer()
			heap := newHeap(participants, iterative)
			heap.sort()
		})
	}
}