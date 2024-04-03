/*
J. Списочная очередь

Любимый вариант очереди Тимофея — очередь, написанная с использованием связного списка.
Помогите ему с реализацией. Очередь должна поддерживать выполнение трёх команд:
- get() — вывести элемент, находящийся в голове очереди, и удалить его.
	Если очередь пуста, то вывести «error».
- put(x) — добавить число x в очередь
- size() — вывести текущий размер очереди

Формат ввода
В первой строке записано количество команд n — целое число, не превосходящее 1000.
В каждой из следующих n строк записаны команды по одной строке.

Формат вывода
Выведите ответ на каждый запрос по одному в строке.

Пример 1
Ввод
10
put -34
put -23
get
size
get
size
get
get
put 80
size

Вывод
-34
1
-23
0
error
error
1

Пример 2
Ввод
6
put -66
put 98
size
size
get
get

Вывод
2
2
-66
98

Пример 3
Ввод
9
get
size
put 74
get
size
put 90
size
size
size
error

Вывод
0
74
0
1
1
1

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

var ErrQueueIsEmpty = errors.New("queue is empty")

type ListNode struct {
	data int
	next *ListNode
}

type Queue struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue) Put(item int) {
	node := &ListNode{
		data: item,
	}
	q.size += 1
	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node
}

func (q *Queue) Get() (int, error) {
	if q.IsEmpty() {
		return 0, ErrQueueIsEmpty
	}
	node := q.head
	q.head = node.next
	if q.head == nil {
		q.tail = nil
	}
	q.size -= 1
	return node.data, nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

type Command struct {
	name string
	parameter int
}

func main() {
	scanner := makeScanner()
	commandsNumber := readInt(scanner)
	queue := NewQueue()
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
	if command.name == "put" {
		queue.Put(command.parameter)
	}
	if command.name == "get" {
		item, err := queue.Get()
		if err != nil {
			fmt.Println("error")
		} else {
			fmt.Println(item)
		}
	}
	if command.name == "size" {
		fmt.Println(queue.size)
	}
}