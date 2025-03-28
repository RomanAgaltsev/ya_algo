/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 2
Задача - B. Калькулятор

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/22781/run-report/111645051/
- Ревью 2 - https://contest.yandex.ru/contest/22781/run-report/111713420/

-- ПРИНЦИП РАБОТЫ --
Принцип работы решения довольно простой:
- Создается структура-калькулятор;
- При создании калькулятора ему сразу передается выражение в ОПН, которое необходимо вычислить;
- Также при создании калькулятора инициируется новый стек для выполнения вычислений.
- Для выполнения вычислений вызывается метод калькулятора, который возвращает результат вычислений.

Для получения результата вычислений:
- Полученная строка-выражение в ОПН разбивается на слайс токенов по пробелам строки;
- Полученный слайс токенов обходится в цикле;
- Если токен является оператором, из стека достается два операнда, выполняется вычисление, результат складывается обратно в стек;
- Если токен является операндом (целым числом), он складывается в стек;
- После завершения цикла на вершине стека находится результат всех вычислений - он возвращается и выводится.

Прочитав описание задачи, предположил, что создание калькулятора-объекта упростит работу с ним, будет более наглядным.
Сейчас логика калькулятора не очень сложная. Но если потребуется её усложнить, работать с структурой будет удобнее.
Опять же, если подумать, что пакет с калькулятором потребуется использовать в других модулях, удобно скрыть его реализацию от потребителя.
И для создания тестов, показалось, так удобнее.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
В постановке задачи говорится, что для решения необходимо использовать стек для вычислений.
Стек используется.

Выражение в ОПН вычислется корректно - приложенные к решению тесты и тесты в Я Контест проходят.

В качестве деления в решении используется метод Floor пакета math.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Для оценки временой сложности, за n принимаем количество токенов во входной строке-выражении в ОПН.

Временные затраты складываются из (чтение строки с ввода исключаем):
- Разбивка строки-выражения на токены по пробелам - O(m), где m - количество символов в строке-выражении
- Обход слайса с токенами в цикле - O(n)
- На каждый токен-операнд выполняется помещение в стек - O(1)
- На каждый токен оператор выполняется
	- извлечение двух операндов из стека - O(1)
	- вычисление оператора над операндами - O(1)
	- помещение результата в стек - O(1)

Общее время - O(n+n) - O(2n) - O(n)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
При решении задачи память потребуется для:
- Мапа для хранения операторов и функций - O(1)
- Стек калькулятора - O(k), где k - размер слайса стека
- Строка-выражение в ОПН - O(m), где m - количество символов в строке-выражении
- Слайс для хранения токенов - разбитая по пробелам строка складывается в этот слайс - O(n)
- Переменные, используемые для вычислений - token, intToken, a, b - O(1)

Получается, что по памяти требуется O(k+m+n).

*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// operations - мапа для хранения операций-функций по ключам-операторам
var operations = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return int(math.Floor(float64(a) / float64(b))) },
}

// Stack - структура стека для хранения операндов
type Stack struct {
	items []int // Слайс для хранения операндов
}

// push - добавляет операнд на вершину стека
func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

// pop - забирает операнд с вершины стека и возвращает его
func (s *Stack) pop() int {
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem
}

// newStack - конструктор стека
func newStack() *Stack {
	return &Stack{items: []int{}}
}

// Calculator - структура калькулятора
type Calculator struct {
	exp string // Выражение в обратной польской нотации, которое необходимо вычислить
	stack *Stack // Стек для операндов, используемый при вычислении
}

// calculate - выполняет вычисление выражения в ОПН и возвращает результат
func (c *Calculator) calculate() int {
	// Строку выражения в ОПН разбиваем по пробелам в слайс строк
	tokens := strings.Split(c.exp, " ")
	// Полученный слайс строк обходим в цикле и выполняем вычисления
	for _, token := range tokens {
		// Проверяем, не оператор ли это
		if operation, ok := operations[token]; ok {
			// Это оператор, выполняем вычисление
			// Забираем из стека два операнда - т.к. хранятся в стеке, то в обратном порядке
			b, a := c.stack.pop(), c.stack.pop()
			// Выполняем операцию - по ключу оператора в мапе лежит функция - можно сразу её выполнить
			// Результат отправляем обратно в стек
			c.stack.push(operation(a, b))
		} else {
			// Это не оператор
			// Преобразовываем строку в целое число
			intToken, _ := strconv.Atoi(token)
			// И отправляем в стек
			c.stack.push(intToken)
		}
	}
	// Результат вычислений на вершине стека - его и возвращаем
	return c.stack.pop()
}

// newCalculator - конструктор калькулятора
// Инициирует новый калькулятор с использованием переданного выражения в ОПН
func newCalculator(expression string) *Calculator {
	return &Calculator{
		exp: expression, // Выражение в ОПН
		stack: newStack(), // Новый стек для вычислений
	}
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
	// Создаем сканер
	scanner := makeScanner()
	// Создаем новый калькулятор с полученным с ввода выражением
	calc := newCalculator(readLine(scanner))
	// Выводим результат вычисления
	fmt.Print(calc.calculate())
}