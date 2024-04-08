/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 2
Задача - B. Калькулятор

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
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem
}

func newStack() *Stack {
	return &Stack{items: []int{}}
}
type Calculator struct {
	exp string
	stack *Stack
}

func (c *Calculator) calculate() int {
	tokens := strings.Split(c.exp, " ")
	for _, token := range tokens {
		if isOperator(token) {
			b, a := c.stack.pop(), c.stack.pop()
			c.stack.push(doOperation(token, a, b))
		} else {
			intToken, _ := strconv.Atoi(token)
			c.stack.push(intToken)
		}
	}
	return c.stack.pop()
}

func newCalculator() *Calculator {
	return &Calculator{
		exp: "",
		stack: newStack(),
	}
}

func isOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func doOperation(op string, a, b int) int {
	var res int
	switch op {
	case "+":
		res = a + b
	case "-":
		res =  a - b
	case "*":
		res =  a * b
	case "/":
		res =  a / b
	}
	return res
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

func main() {
	scanner := makeScanner()
	calc := newCalculator()
	calc.exp = readLine(scanner)
	fmt.Print(calc.calculate())
}