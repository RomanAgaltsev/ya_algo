/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 5
Задача - A. Пирамидальная сортировка

Отчеты:
- Ревью 1 -

goos: windows
goarch: amd64
pkg: main/final1
cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkHeapsort/Size=10000/Iterative=true/cores-12                 1000000000               0.001506 ns/op
BenchmarkHeapsort/Size=100000/Iterative=true/cores-12                1000000000               0.02506 ns/op
BenchmarkHeapsort/Size=1000000/Iterative=true/cores-12               1000000000               0.4065 ns/op
BenchmarkHeapsort/Size=10000/Iterative=false/cores-12                1000000000               0.002000 ns/op
BenchmarkHeapsort/Size=100000/Iterative=false/cores-12               1000000000               0.02707 ns/op
BenchmarkHeapsort/Size=1000000/Iterative=false/cores-12              1000000000               0.4361 ns/op
PASS
ok      main/final1     33.288s

goos: windows
goarch: amd64
pkg: main/final1
cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkHeapsort/Size=10000/Iterative=true/cores-12                 1000000000               0.001504 ns/op
BenchmarkHeapsort/Size=100000/Iterative=true/cores-12                1000000000               0.02514 ns/op
BenchmarkHeapsort/Size=1000000/Iterative=true/cores-12               1000000000               0.3903 ns/op
BenchmarkHeapsort/Size=10000/Iterative=false/cores-12                1000000000               0.002001 ns/op
BenchmarkHeapsort/Size=100000/Iterative=false/cores-12               1000000000               0.02618 ns/op
BenchmarkHeapsort/Size=1000000/Iterative=false/cores-12              1000000000               0.4675 ns/op
PASS
ok      main/final1     32.995s
*/

package main

import (
	"bufio"
	"cmp"
	"os"
	"strconv"
	"strings"
)

// Participant - структура для хранения данных участника
type Participant struct {
	login string // Логин участника
	solve int // Количество решенных задач
	penalty int // Штраф
}

type Heap struct {
	arr []Participant
	length int
	size int
	iterative bool
	less func(a, b Participant) bool
}

func (h *Heap) build() {
	for i := h.length/2; i >= 1; i-- {
		if h.iterative {
			h.heapifyIterative(i)
		} else {
			h.heapifyRecursive(i)
		}
	}
}

func (h *Heap) heapifyRecursive(i int) {
	largest := h.largest(i)
	if largest != i {
		h.swap(i, largest)
		h.heapifyRecursive(largest)
	}
}

func (h *Heap) heapifyIterative(i int) {
	for {
		largest := h.largest(i)
		if largest != i {
			h.swap(i, largest)
			i = largest
			continue
		}
		break
	}
}

func (h *Heap) sort() {
	h.build()
	for i := h.length; i >= 2; i-- {
		h.swap(1, i)
		h.size -= 1
		if h.iterative {
			h.heapifyIterative(1)
		} else {
			h.heapifyRecursive(1)
		}
	}
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) largest(i int) int {
	largest := i
	l, r := left(i), right(i)
	if l <= h.size && h.less(h.arr[i], h.arr[l]) {
		largest = l
	}
	if r <= h.size && h.less(h.arr[largest], h.arr[r]) {
		largest = r
	}
	return largest
}

func newHeap(arr []Participant, iterative bool) *Heap {
	return &Heap{
		arr: append([]Participant{{}}, arr...),
		length: len(arr),
		size: len(arr),
		iterative: iterative,
		less: func(a, b Participant) bool {
			if res := cmp.Compare(b.solve, a.solve); res != 0 {
				return res == -1
			}
			if res := cmp.Compare(a.penalty, b.penalty); res != 0 {
				return res == -1
			}
			return strings.Compare(a.login, b.login) == -1
		},
	}
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func main() {
	// Создаем сканер
	scanner := makeScanner()
	// Читаем количество участников
	n := readInt(scanner)
	// Читаем участников в слайс структур
	participants := readParticipants(scanner, n)
	// Создаем новую кучу с участниками
	heap := newHeap(participants, true)
	// Сортируем кучу с участниками
	heap.sort()
	// Выводим участников после сортировки - только логины
	printParticipants(heap.arr)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
func readParticipants(scanner *bufio.Scanner, n int) []Participant {
	participants := make([]Participant, n)
	for i := 0; i < n; i++ {
		participants[i] = readParticipant(scanner)
	}
	return participants
}

func readParticipant(scanner *bufio.Scanner) Participant {
	scanner.Scan()
	listParticipant := strings.Split(scanner.Text(), " ")
	login := listParticipant[0]
	solve, _ := strconv.Atoi(listParticipant[1]) 
	penalty, _ := strconv.Atoi(listParticipant[2])
	return Participant{
		login,
		solve,
		penalty,
	}
}

func printParticipants(participants []Participant) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 1; i < len(participants); i++ {
		writer.WriteString(participants[i].login)
		writer.WriteString("\n")
	}
	writer.Flush()
}