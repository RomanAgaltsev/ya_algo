/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Задача - A. Дек

Отчеты:
- Ревью 1 -

-- ПРИНЦИП РАБОТЫ --

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

*/

package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	ErrDequeIsEmpty = errors.New("deque is empty")
	ErrDequeIsFull = errors.New("deque is full")
)

type Deque struct {
	deque []int
	head int
	tail int
	max int
	size int
}

func NewDeque(n int) *Deque {
	return &Deque{
		deque: make([]int, n),
		head: 0,
		tail: 0,
		max: n,
		size: 0,
	}
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) IsFull() bool {
	return d.size == d.max
}

func (d *Deque) MoveHead(step int) {
	d.head = (d.head + step + d.max) % d.max
}

func (d *Deque) MoveTail(step int) {
	d.tail = (d.tail + step + d.max) % d.max
}

func (d *Deque) PushBack(value int) error {
	if d.IsFull() {
		return ErrDequeIsFull
	}
	d.deque[d.tail] = value
	d.size += 1
	
	if !d.IsFull() {
		if d.head == d.tail {
			d.MoveHead(-1)
		}
		d.MoveTail(1)
	}
	return nil
}

func (d *Deque) PushFront(value int) error {
	if d.IsFull() {
		return ErrDequeIsFull
	}
	d.deque[d.head] = value
	d.size += 1
	if !d.IsFull() {
		if d.head == d.tail {
			d.MoveTail(1)
		}
		d.MoveHead(-1)
	}
	return nil
}

func (d *Deque) PopFront() (int, error) {
	if d.IsEmpty() {
		return 0, ErrDequeIsEmpty
	}
	if !d.IsFull() {
		d.MoveHead(1)
	}
	value := d.deque[d.head]
	d.deque[d.head] = 0
	d.size -= 1	
	if d.IsEmpty() {
		d.tail = d.head
	}
	return value, nil
}

func (d *Deque) PopBack() (int, error) {
	if d.IsEmpty() {
		return 0, ErrDequeIsEmpty
	}	
	if !d.IsFull() {
		d.MoveTail(-1)
	}	
	value := d.deque[d.tail]
	d.deque[d.tail] = 0
	d.size -= 1	
	if d.IsEmpty() {
		d.head = d.tail
	}
	return value, nil
}

type Command struct {
	name string
	parameter int
}

func main() {
	scanner := makeScanner()
	commandsNumber := readInt(scanner)
	dequeMax := readInt(scanner)
	deque := NewDeque(dequeMax)
	commands := make([]Command, commandsNumber)
	for i := 0; i < commandsNumber; i++ {
		commands = append(commands, readCommand(scanner))
	}
	writer := bufio.NewWriter(os.Stdout)
	for _, command := range commands {
		executeCommand(deque, command, writer)
	}
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 10 * 1024 * 1024
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

func executeCommand(deque *Deque, command Command, writer *bufio.Writer) {
	if command.name == "push_back" {
		err := deque.PushBack(command.parameter)
		if err != nil {
			writer.WriteString("error\n")
		}
	}
	if command.name == "push_front" {
		err := deque.PushFront(command.parameter)
		if err != nil {
			writer.WriteString("error\n")
		}
	}
	if command.name == "pop_front" {
		value, err := deque.PopFront()
		if err != nil {
			writer.WriteString("error\n")
		} else {
			writer.WriteString(strconv.Itoa(value))
			writer.WriteString("\n")
		}
	}
	if command.name == "pop_back" {
		value, err := deque.PopBack()
		if err != nil {
			writer.WriteString("error\n")
		} else {
			writer.WriteString(strconv.Itoa(value))
			writer.WriteString("\n")
		}
	}
}