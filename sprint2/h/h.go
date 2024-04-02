/*
H. Скобочная последовательность

Вот какую задачу Тимофей предложил на собеседовании одному из кандидатов.
Если вы с ней ещё не сталкивались, то наверняка столкнётесь –— она довольно популярная.

Дана скобочная последовательность. Нужно определить, правильная ли она.

Будем придерживаться такого определения:
- пустая строка —– правильная скобочная последовательность;
- правильная скобочная последовательность, взятая в скобки одного типа, –— правильная скобочная последовательность;
- правильная скобочная последовательность с приписанной слева или справа правильной скобочной последовательностью — тоже правильная.

На вход подаётся последовательность из скобок трёх видов: [], (), {}.
Напишите функцию is_correct_bracket_seq, которая принимает на вход скобочную последовательность и возвращает True,
если последовательность правильная, а иначе False.

Формат ввода
На вход подаётся одна строка, содержащая скобочную последовательность. Скобки записаны подряд, без пробелов.

Формат вывода
Выведите «True» или «False».

Пример 1
Ввод
{[()]}
Вывод
True

Пример 2
Ввод
()
Вывод
True

*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrStackIsEmpty = errors.New("stack is empty")

type Stack struct {
	items []string
}

func NewStack() *Stack {
	return &Stack{items: []string{}}
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", ErrStackIsEmpty
	} 
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem, nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func IsCorrectBracketSeq(seq string) string {
	if seq == "" {
		return "True"
	}
	stack := NewStack()
	for _, b := range seq {
		bracket := string(b)
		if bracket == "(" || bracket == "[" || bracket == "{" {
			stack.Push(bracket)
		} else {
			stBracket, err := stack.Pop()
			if err != nil {
				return "False"
			}
			if bracket == ")" && stBracket != "(" {
				return "False"
			}
			if bracket == "]" && stBracket != "[" {
				return "False"
			}
			if bracket == "}" && stBracket != "{" {
				return "False"
			}
		}
	}
	if stack.Size() != 0 {
		return "False"
	}
	return "True"
}

func main() {
	scanner := makeScanner()
	fmt.Print(IsCorrectBracketSeq(readLine(scanner)))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}