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
	errDequeIsEmpty = errors.New("deque is empty")
	errDequeIsFull = errors.New("deque is full")
)

type Deque struct {
	deque []int
	head int
	tail int
	max int
	size int
}

func newDeque(n int) *Deque {
	return &Deque{
		deque: make([]int, n),
		head: 0,
		tail: 0,
		max: n,
		size: 0,
	}
}

func (d *Deque) isEmpty() bool {
	return d.size == 0
}

func (d *Deque) isFull() bool {
	return d.size == d.max
}

func (d *Deque) moveHead(step int) {
	d.head = (d.head + step + d.max) % d.max
}

func (d *Deque) moveTail(step int) {
	d.tail = (d.tail + step + d.max) % d.max
}

func (d *Deque) pushBack(value int) error {
	if d.isFull() {
		return errDequeIsFull
	}
	d.deque[d.tail] = value
	d.moveTail(1)
	d.size += 1
	return nil
}

func (d *Deque) pushFront(value int) error {
	if d.isFull() {
		return errDequeIsFull
	}
	d.moveHead(-1)
	d.deque[d.head] = value
	d.size += 1
	return nil
}

func (d *Deque) popFront() (int, error) {
	if d.isEmpty() {
		return 0, errDequeIsEmpty
	}
	value := d.deque[d.head]
	d.moveHead(1)
	d.size -= 1
	return value, nil
}

func (d *Deque) popBack() (int, error) {
	if d.isEmpty() {
		return 0, errDequeIsEmpty
	}
	d.moveTail(-1)
	value := d.deque[d.tail]
	d.size -= 1
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
	deque := newDeque(dequeMax)
	commands := make([]Command, commandsNumber)
	for i := 0; i < commandsNumber; i++ {
		commands = append(commands, readCommand(scanner))
	}
	result := ""
	writer := bufio.NewWriter(os.Stdout)
	for _, command := range commands {
		result = executeCommand(deque, command)
		if result != "" {
			writer.WriteString(result)
			writer.WriteString("\n")
		}
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

func executeCommand(deque *Deque, command Command) string {
	if command.name == "push_back" {
		err := deque.pushBack(command.parameter)
		if err != nil {
			return "error"
		}
	}
	if command.name == "push_front" {
		err := deque.pushFront(command.parameter)
		if err != nil {
			return "error"
		}
	}
	if command.name == "pop_front" {
		value, err := deque.popFront()
		if err != nil {
			return "error"
		} else {
			return strconv.Itoa(value)
		}
	}
	if command.name == "pop_back" {
		value, err := deque.popBack()
		if err != nil {
			return "error"
		} else {
			return strconv.Itoa(value)
		}
	}
	return ""
}