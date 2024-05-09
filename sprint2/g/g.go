/*
G. Стек - MaxEffective

Реализуйте класс StackMaxEffective, поддерживающий операцию определения максимума среди элементов в стеке.
Сложность операции должна быть O(1).
Для пустого стека операция должна возвращать None.
При этом push(x) и pop() также должны выполняться за константное время.

Формат ввода
В первой строке записано одно число — количество команд, оно не превосходит 100000.
Далее идут команды по одной в строке.
Команды могут быть следующих видов:
push(x) — добавить число x в стек. Число x не превышает 105;
pop() — удалить число с вершины стека;
get_max() — напечатать максимальное число в стеке;
top() — напечатать число с вершины стека;

Если стек пуст, при вызове команды get_max нужно напечатать «None», для команды pop и top — «error».

Формат вывода
Для каждой команды get_max() напечатайте результат её выполнения.
Если стек пустой, для команды get_max() напечатайте «None».
Если происходит удаление из пустого стека — напечатайте «error».

Пример 1
Ввод
13
pop
pop
top
push 4
push -5
top
push 7
pop
pop
get_max
top
pop
get_max

Вывод
error
error
error
-5
4
4
None

Пример 2
Ввод
10
get_max
push -6
pop
pop
get_max
push 2
get_max
pop
push -2
push -6

Вывод
None
error
None
2

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

var ErrStackIsEmpty = errors.New("stack is empty")

type MaxOfStack struct {
	maximums []int
}

func NewMaxOfStack() *MaxOfStack {
	return &MaxOfStack{maximums: []int{-100001}}
}

func (m *MaxOfStack) Push(item int) {
	m.maximums = append(m.maximums, item)
}

func (m *MaxOfStack) Pop() int {
	lastIndex := len(m.maximums) - 1
	lastMaximum := m.maximums[lastIndex]
	m.maximums = m.maximums[:lastIndex]
	return lastMaximum
}

func (m *MaxOfStack) Peek() int {
	lastIndex := len(m.maximums) - 1
	return m.maximums[lastIndex]
}

type StackMaxEffective struct {
	items []int
	maxItems *MaxOfStack
}

func NewStack() *StackMaxEffective {
	return &StackMaxEffective{
		items: []int{},
		maxItems: NewMaxOfStack(),
	}
}

func (s *StackMaxEffective) Push(item int) {
	s.items = append(s.items, item)
	maxItem := s.maxItems.Peek()
	if item > maxItem {
		s.maxItems.Push(item)
	} else {
		s.maxItems.Push(maxItem)
	}
}

func (s *StackMaxEffective) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrStackIsEmpty
	}
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	s.maxItems.Pop()
	return lastItem, nil
}

func (s *StackMaxEffective) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrStackIsEmpty
	}
	lastIndex := len(s.items) - 1
	return s.items[lastIndex], nil
}

func (s *StackMaxEffective) GetMax() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrStackIsEmpty
	}
	return s.maxItems.Peek(), nil 
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

func executeCommand(command Command, stack *StackMaxEffective) {
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
	if command.name == "top" {
		item, err := stack.Peek()
		if err != nil {
			fmt.Println("error")
		} else {
			fmt.Println(item)
		}
	}
}