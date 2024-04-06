/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 2
Задача - A. Дек

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/22781/run-report/111443350/

-- ПРИНЦИП РАБОТЫ --

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

-- ОПТИМИЗАЦИИ --
При решении задачи сделал следующие оптимизации:
1. Буферизированный вывод
Сначала сделал простой вывод, через fmt.Println.
 - Время теста - 213-217 ms
 - Память - 18,43-18,45 mb
Потом сделал буферизированный вывод через bufio.Writer.
 - Время теста - 50-57 ms
 - Память - 17,55-18,40 mb

2. Хранение команд
Сначала команды считывались с ввода и складывались в слайс. Потом слайс обходился в цикле и команды выполнялись.
- Время теста - 50-57 ms
- Память - 17,55-18,40 mb
Потом отказался от слайса для хранения команд - можно их выполнять сразу после чтения, без промежуточного хранения.
- Время теста - 28-32 ms
- Память - 4,96-5,95 mb

*/

package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Переменные для хранения ошибок
// С точки зрения вывода, какая из ошибок случилась - дек пустой или полный  - разницы нет, всегда выводим "error"
// Но для лучшей читаемости кода и для тестов (чтобы проверить тип ошибки) завел переменные
var (
	// Дек пустой
	errDequeIsEmpty = errors.New("deque is empty")
	// Дек полный
	errDequeIsFull = errors.New("deque is full")
)

// Deque - структура дека
type Deque struct {
	deque []int // Буфер храним в слайсе интов
	head int // Указатель на голову дека
	tail int // Указатель на хвост дека
	max int // Максимальный размер дека
	size int // Текущий размер дека
}

// newDeque - конструктор дека
// Инициализируем новый дек с использованием переданного размера буфера
func newDeque(n int) *Deque {
	return &Deque{
		deque: make([]int, n), // Инициализируем буфер (слайс интов) размера n
		head: 0, // Указатель на голову смотрит на 0
		tail: 0, // Указатель на хвост смотри на 0
		max: n, // Максимальный размер дека = n
		size: 0, // Текущий размер нового дека = 0
	}
}

// isEmpty - проверяет, пустой ли дек
func (d *Deque) isEmpty() bool {
	// Дек пустой, если его размер = 0
	return d.size == 0
}
// isFull - проверяет, полный ли дек
func (d *Deque) isFull() bool {
	// Дек полный, если его размер = максимальному
	return d.size == d.max
}

// moveHead - передвигает указатель на голову на переданное количество шагов:
// - положительное значение шага - двигаем вправо
// - отрицательное значение шага - двигаем влево
func (d *Deque) moveHead(step int) {
	// Формула одинаковая и для положительного, и для отрицательного шага
	d.head = (d.head + step + d.max) % d.max
}

// moveTail - передвигает указатель на хвост на переданное количество шагов:
// - положительное значение шага - двигаем вправо
// - отрицательное значение шага - двигаем влево
func (d *Deque) moveTail(step int) {
	// Формула одинаковая и для положительного, и для отрицательного шага
	d.tail = (d.tail + step + d.max) % d.max
}

// pushBack - добавляет переданное значение в хвост дека
func (d *Deque) pushBack(value int) error {
	// Проверяем, не полный ли дек
	if d.isFull() {
		// Дек полный, добавлять некуда - возвращаем ошибку
		return errDequeIsFull
	}
	// Записываем переданное значение в хвост - он всегда смотрит на свободную ячейку, если дек не полный
	d.deque[d.tail] = value
	// Двигаем хвост на 1 шаг вправо
	d.moveTail(1)
	// Увеличиваем текущий размер дека
	d.size += 1
	// Возвращаем ничего
	return nil
}

// pushFront - добавляет переданное значение в голову дека
func (d *Deque) pushFront(value int) error {
	// Проверяем, не полный ли дек
	if d.isFull() {
		// Дек полный, добавлять некуда - возвращаем ошибку
		return errDequeIsFull
	}
	// Двигаем голову на 1 шаг влево - голова всегда смотрит на занятую ячейку, если дек не пустой
	d.moveHead(-1)
	// Записываем переданное значение в голову
	d.deque[d.head] = value
	// Увеличиваем текущий размер дека
	d.size += 1
	// Возвращаем ничего
	return nil
}

// popFront - возвращает значение с головы и удаляет его из дека
func (d *Deque) popFront() (int, error) {
	// Проверяем, не пустой ли дек
	if d.isEmpty() {
		// Дек пустой, возвращать нечего - возвращаем ошибку
		return 0, errDequeIsEmpty
	}
	// Берем значение с головы - голова всегда смотрит на занятую ячейку, если дек не пустой
	value := d.deque[d.head]
	// Двигаем голову на 1 шаг право
	d.moveHead(1)
	// Уменьшаем текущий размер дека
	d.size -= 1
	// Возвращаем значение
	return value, nil
}

// popBack - возвращает значение с хвоста и удаляет его из дека
func (d *Deque) popBack() (int, error) {
	// Проверяем, не пустой ли дек
	if d.isEmpty() {
		// Дек пустой, возвращать нечего - возвращаем ошибку
		return 0, errDequeIsEmpty
	}
	// Двигаем хвост на 1 шаг влево - он всегда смотрит на свободную ячейку, если дек не полный
	d.moveTail(-1)
	// Берем значение с хвоста
	value := d.deque[d.tail]
	// Уменьшаем текущий размер дека
	d.size -= 1
	// Возвращаем значение
	return value, nil
}

// Command - структура команды
type Command struct {
	name string // Имя команды
	parameter int // Параметр команды
}

func main() {
	// Создаем сканер для чтения ввода
	scanner := makeScanner()
	// Читаем количество команд и максимальный размер дека
	commandsNumber, dequeMax := readInt(scanner), readInt(scanner)
	// Создаем новый дек - он сразу инициализируется
	deque := newDeque(dequeMax)
	// Переменная для результата выполнения команды
	result := ""
	// Создаем писатель для вывода
	writer := bufio.NewWriter(os.Stdout)
	// В цикле по считанному ранее количеству читаем команды с ввода, выполняем, результат кладем в писатель
	for i := 0; i < commandsNumber; i++ {
		// readCommand - читаем команду с ввода сканером
		// executeCommand - выполняем команду, получаем результат
		result = executeCommand(deque, readCommand(scanner))
		// Выводим только непустые результаты
		if result != "" {
			writer.WriteString(result)
			writer.WriteString("\n")
		}
	}
	writer.Flush()
}

// makeScanner - создает новый сканер
func makeScanner() *bufio.Scanner {
	const maxCapacity = 10 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

// readInt - читает целое число сканером с ввода
func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

// readCommand - читает команду
func readCommand(scanner *bufio.Scanner) Command {
	// Создаем новую структур команды
	command := Command{}
	// Читаем строку команды
	scanner.Scan()
	// Разбиваем строку команды на имя и параметр
	listCommand := strings.Split(scanner.Text(), " ")
	// Проверяем, есть ли вообще что-то
	if len(listCommand) >= 1 {
		// Что-то есть, имя команды берем из первого элемента
		command.name = listCommand[0]
	}
	// Проверяем, есть ли второй элемент - в нем должен лежать параметр команды
	if len(listCommand) == 2 {
		// Второй элемент есть, берем его, конвертируя в целое число
		command.parameter, _ = strconv.Atoi(listCommand[1])
	}
	// Возвращаем структуру команды
	return command
}

// executeCommand - выполняет переданную команду с переданным же деком
func executeCommand(deque *Deque, command Command) string {
	// Проверяем имя команды
	if command.name == "push_back" {
		// Это добавление в хвост
		err := deque.pushBack(command.parameter)
		// Проверяем, есть ли ошибка добавления
		if err != nil {
			// Ошибка есть, возвращаем "error"
			return "error"
		}
	}
	// Проверяем имя команды
	if command.name == "push_front" {
		// Это добавление в голову
		err := deque.pushFront(command.parameter)
		// Проверяем, если ли ошибка добавления
		if err != nil {
			// Ошибка есть, возвращаем "error"
			return "error"
		}
	}
	// Проверяем имя команды
	if command.name == "pop_front" {
		// Это вывод с удалением с головы
		value, err := deque.popFront()
		// Проверяем, есть ли ошибка
		if err != nil {
			// Ошибка есть, возвращаем "error"
			return "error"
		} else {
			// Ошибки нет, возвращаем значение, конвертируя его в строку
			return strconv.Itoa(value)
		}
	}
	// Проверяем имя команды
	if command.name == "pop_back" {
		// Это вывод с удалением с хвоста
		value, err := deque.popBack()
		// Проверяем, есть ли ошибка
		if err != nil {
			// Ошибка есть, возвращаем "error"
			return "error"
		} else {
			// Ошибки нет, возвращаем значение, конвертируя его в строку
			return strconv.Itoa(value)
		}
	}
	// Все предыдущие return сидят в if-ах, компилятор требует добавить return тут
	return ""
}