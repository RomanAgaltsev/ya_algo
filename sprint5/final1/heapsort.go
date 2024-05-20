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

// Heap - структура для реализации кучи
type Heap struct {
	arr []Participant // Слайс участников - сама куча
	length int // Длина всей кучи (слайса)
	size int // Оставшаяся длина кучи в слайсе
	iterative bool // Флаг режима для функции heapify. True - итеративный вариант, false - рекурсивный
	less func(a, b Participant) bool // Функция-компаратор для сравнения участников
}

// build - преобразовывает массив (слайс) участников в кучу (строит кучу)
func (h *Heap) build() {
	// Обходим массив, начиная с листьев, и выполняем обмены, чтобы построить невозрастающую кучу
	for i := h.length/2; i >= 1; i-- {
		// Проверяем, в каком режиме надо выполнять heapify
		if h.iterative {
			// Итеративный вариант
			// Выполняем перестановки для i-того элемента
			h.heapifyIterative(i)
		} else {
			// Рекурсивный вариант
			// Выполняем перестановки для i-того элемента
			h.heapifyRecursive(i)
		}
	}
}

// heapifyRecursive - рекурсивно выполняет перестановки i-того элемента массива
func (h *Heap) heapifyRecursive(i int) {
	// Получаем наибольший элемент из родителя (i), левого (2 * i) и правого (2*i + 1) ребенка
	largest := h.largest(i)
	// Проверяем, индекс наибольшего элемента равен ли i-тому
	if largest != i {
		// Если не равен, меняем меставами i-тый и наибольший элементы
		h.swap(i, largest)
		// Проваливаемся в рекурсивный вызов для дальнейших перестановок
		h.heapifyRecursive(largest)
	}
}

// heapifyIterative - итеративно выполняет перестановки i-того элемента массива
func (h *Heap) heapifyIterative(i int) {
	// Запускаем бесконечный цикл перестановок
	// Остановимся в тот момент, когда элемент займет своё место
	for {
		// Получаем наибольший элемент из родителя (i), левого (2 * i) и правого (2*i + 1) ребенка
		largest := h.largest(i)
		// Проверяем, индекс наибольшего элемента равен ли i-тому
		if largest != i {
			// Если не равен, меняем меставами i-тый и наибольший элементы
			h.swap(i, largest)
			// Подменяем i-тый элемент наибольшим для дальнейших перестановок
			i = largest
			continue
		}
		// Перестановки больше делать не надо - элемент занял своё место
		break
	}
}

// sort - выполняет пирамидальную сортировку массива (слайса) участников
func (h *Heap) sort() {
	// Сначала строим неубывающую кучу
	h.build()
	// Затем выполняем саму сортировку
	// Обходим массив в цикле, выполняя перестановки - идем с конца до 2-го элемента
	for i := h.length; i >= 2; i-- {
		// Так как куча была построена неубывающая, максимальный элемент находится на вершине - 1-ый элемент массива
		// Меняем местами 1-ый и i-тый элементы
		h.swap(1, i)
		// Уменьшаем размер кучи в массиве
		h.size -= 1
		// После обмена бывший 1-ый (максимальный) элемент занял "своё" место в отсортированной части массива
		// А бывший i-тый элемент занял 1-ое место - место максимального элемента в куче
		// Если этот элемент не является максимальным, надо переместить его на "своё" место в куче
		// Проверяем, в каком режиме надо выполнять heapify
		if h.iterative {
			// Итеративный вариант
			// Выполняем перестановки для нового 1-го элемента
			h.heapifyIterative(1)
		} else {
			// Рекурсивный вариант
			// Выполняем перестановки для нового 1-го элемента
			h.heapifyRecursive(1)
		}
	}
}

// swap - меняет местами элементы массива с переданными индексами
func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// largest - определяет и возвращает индекс наибольшего элемента из родителя (i), левого (2 * i) и правого (2*i + 1) ребенка
func (h *Heap) largest(i int) int {
	// Сначала за наибольший возьмем i-тый элемент
	largest := i
	// Получим индексы левого и правого ребенка
	l, r := left(i), right(i)
	// Проверяем, находится ли левый ребенок в пределах кучи и больше ли он своего родителя
	if l <= h.size && h.less(h.arr[i], h.arr[l]) {
		// Левый ребенок больше своего родителя
		largest = l
	}
	// Проверяем, находится ли правый ребенок в пределах кучи и больше ли он наибольшего из родителя и левого ребенка
	if r <= h.size && h.less(h.arr[largest], h.arr[r]) {
		// Правый ребенок больше наибольшего из родителя и левого ребенка
		largest = r
	}
	// Возвращаем индекс наибольшего элемента
	return largest
}

// newHeap - конструктор кучи
func newHeap(arr []Participant, iterative bool) *Heap {
	return &Heap{
		arr: append([]Participant{{}}, arr...), // Сортируемый массив участников, он же сама куча - начинается с индекса 1
		length: len(arr), // Размер массива
		size: len(arr), // Оставшийся размер кучи в массиве (слайсе)
		iterative: iterative, // Режим функции heapify
		// Функция-компаратор, сравнивает двух участников в соответствии с условиями задачи:
		// - сначала по количеству решенных задач, больше задач - меньше позиция участника
		// - затем по размеру штрафа, больше штраф - больше позиция участника
		// - в конце по логину в лексикографическом порядке
		// Возвращает:
		//	true, если участник a меньше участника b
		//	false, если наоборот
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

// left - возвращает индекс левого ребенка для i-того элемента
func left(i int) int {
	return 2 * i
}

// right - возвращает индекс правого ребенка для i-того элемента
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
	// Второй аргумент определяет вариант работы функции heapify:
	// - true, если используется итеративный вариант
	// - false, если рекурсивный
	heap := newHeap(participants, true)
	// Сортируем кучу с участниками пирамидальной сортировкой
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