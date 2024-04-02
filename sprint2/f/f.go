/*
F. Стек - Max

Нужно реализовать класс StackMax, который поддерживает операцию определения максимума среди всех элементов в стеке.
Класс должен поддерживать операции push(x), где x – целое число, pop() и get_max().

Формат ввода
В первой строке записано одно число n — количество команд, которое не превосходит 10000.
В следующих n строках идут команды. Команды могут быть следующих видов:

push(x) — добавить число x в стек. Число x не превышает 105;
pop() — удалить число с вершины стека;
get_max() — напечатать максимальное число в стеке;
Если стек пуст, при вызове команды get_max() нужно напечатать «None», для команды pop() — «error».

Формат вывода
Для каждой команды get_max() напечатайте результат её выполнения.
Если стек пустой, для команды get_max() напечатайте «None».
Если происходит удаление из пустого стека — напечатайте «error».

Пример 1
Ввод
8
get_max
push 7
pop
push -2
push -1
pop
get_max
get_max

Вывод
None
-2
-2

Пример 2
Ввод
7
get_max
pop
pop
pop
push 10
get_max
push -9

Вывод
None
error
error
error
10

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

const minInt = -100001

var ErrStackIsEmpty = errors.New("stack is empty")

type StackMax struct {
	items []int
	maxItem int
}

func NewStack() *StackMax {
	return &StackMax{items: []int{}, maxItem: minInt}
}

func (s *StackMax) Push(item int) {
	s.items = append(s.items, item)
	if item > s.maxItem {
		s.maxItem = item
	}
}

func (s *StackMax) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrStackIsEmpty
	}
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	maxItem := minInt
	for _, item := range s.items {
		if item > maxItem {
			maxItem = item
		}
	}
	s.maxItem = maxItem
	return lastItem, nil
}

func (s *StackMax) GetMax() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrStackIsEmpty
	}
	return s.maxItem, nil 
}

type Command struct {
	name string
	parameter int
}

func main() {
	stack := NewStack()
	scanner := makeScanner()
	commandsNumber := readInt(scanner)
	commands := make([]Command, commandsNumber)
	for i := 0; i < commandsNumber; i++ {
		commands = append(commands, readCommand(scanner))
	}
	for _, command := range commands {
		executeCommand(command, stack)
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

func executeCommand(command Command, stack *StackMax) {
	if command.name == "push" {
		stack.Push(command.parameter)
	}
	if command.name == "pop" {
		_, err := stack.Pop()
		if err != nil {
			fmt.Println("error")
		}
	}
	if command.name == "get_max" {
		item, err := stack.GetMax()
		if err != nil {
			fmt.Println("None")
		} else {
			fmt.Println(item)
		}
	}
}