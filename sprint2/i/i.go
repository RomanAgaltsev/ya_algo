/*
I. Ограниченная очередь

Астрологи объявили день очередей ограниченного размера.
Тимофею нужно написать класс MyQueueSized, который принимает параметр max_size,
означающий максимально допустимое количество элементов в очереди.

Помогите ему — реализуйте программу, которая будет эмулировать работу такой очереди.
Функции, которые надо поддержать, описаны в формате ввода.

Формат ввода
В первой строке записано одно число — количество команд, оно не превосходит 5000.
Во второй строке задан максимально допустимый размер очереди, он не превосходит 5000.
Далее идут команды по одной на строке. Команды могут быть следующих видов:
- push(x) — добавить число x в очередь;
- pop() — удалить число из очереди и вывести на печать;
- peek() — напечатать первое число в очереди;
- size() — вернуть размер очереди;

При превышении допустимого размера очереди нужно вывести «error».
При вызове операций pop() или peek() для пустой очереди нужно вывести «None».

Формат вывода
Напечатайте результаты выполнения нужных команд, по одному на строке.

Пример 1
Ввод
8
2
peek
push 5
push 2
peek
size
size
push 1
size

Вывод
None
5
2
2
error
2

Пример 2
Ввод
10
1
push 1
size
push 3
size
push 1
pop
push 1
pop
push 3
push 3

Вывод
1
error
1
error
1
1
error
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrQueueIsEmpty = errors.New("queue is empty")
	ErrQueueIsFull = errors.New("queue is full")
)

type Queue struct {
	queue []int
	head int
	tail int
	max int
	size int
}

func NewQueue(n int) *Queue {
	return &Queue{
		queue: make([]int, n),
		head: 0,
		tail: 0,
		max: n,
		size: 0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Push(item int) error {
	if q.size == q.max {
		return ErrQueueIsFull
	}
	q.queue[q.tail] = item
	q.tail = (q.tail + 1) % q.max
	q.size += 1
	return nil
}

func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, ErrQueueIsEmpty
	}
	item := q.queue[q.head]
	q.queue[q.head] = 0
	q.head = (q.head + 1) % q.max
	q.size -= 1
	return item, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, ErrQueueIsEmpty
	}
	return q.queue[q.head], nil
}

type Command struct {
	name string
	parameter int
}

func main() {
	scanner := makeScanner()
	commandsNumber := readInt(scanner)
	queue := NewQueue(readInt(scanner))
	commands := make([]Command, commandsNumber)
	for i := 0; i < commandsNumber; i++ {
		commands = append(commands, readCommand(scanner))
	}
	for _, command := range commands {
		executeCommand(queue, command)
	}
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

func readCommand(scanner *bufio.Scanner) Command {
	command := Command{}
	scanner.Scan()
	listCommand := strings.Split(scanner.Text(), " ")
	if len(listCommand) >= 1 {
		command.name = listCommand[0]
	}
	if len(listCommand) == 2 {
		command.parameter, _ = strconv.Atoi(listCommand[1])
	}
	return command
}

func executeCommand(queue *Queue, command Command) {
	if command.name == "push" {
		err := queue.Push(command.parameter)
		if err != nil {
			fmt.Println("error")
		}
	}
	if command.name == "pop" {
		item, err := queue.Pop()
		if err != nil {
			fmt.Println("None")
		} else {
			fmt.Println(item)
		}
	}
	if command.name == "peek" {
		item, err := queue.Peek()
		if err != nil {
			fmt.Println("None")
		} else {
			fmt.Println(item)
		}
	}
	if command.name == "size" {
		fmt.Println(queue.size)
	}
}