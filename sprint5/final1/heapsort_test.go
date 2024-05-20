package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

// getRandomString - возвращает случайную строку длины n
func getRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

// BenchmarkHeapsort - запускает бенчмарк
func BenchmarkHeapsort(b *testing.B) {
	// Бенчмарк итеративного варианта heapify
	RunBenchmarkHeapsort(b, true)
	// Бенчмарк рекурсивного варианта heapify
	RunBenchmarkHeapsort(b, false)
}

func RunBenchmarkHeapsort(b *testing.B, iterative bool) {
	// Тестируем на 10 000, 100 000 и 1 000 000 элементах
	for _, size := range []int{10_000, 100_000, 1_000_000} {
		b.Run(fmt.Sprintf("Size=%d/Iterative=%v/cores", size, iterative), func(b *testing.B) {
			// Инициируем слайс заданного размера
			participants := make([]Participant, size)
			// Наполяем слайс структурами участников
			for i := range participants {
				// Создаем нового участника
				participants[i] = Participant{
					getRandomString(20), // Логин всегда 20 символов
					rand.Intn(1_000_000_000), // Количество решенных задач - случайное число от 0 до 10 в 9 степени
					rand.Intn(1_000_000_000), // Размер штрафа - случайное число от 0 до 10 в 9 степени
				}
			}
			// Сбрасываем таймер
			b.ResetTimer()
			// Создаем новую кучу
			heap := newHeap(participants, iterative)
			// Выполняем сортировку
			heap.sort()
		})
	}
}