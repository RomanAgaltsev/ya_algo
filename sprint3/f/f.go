/*
F. Периметр треугольника

Перед сном Рита решила поиграть в игру на телефоне.
Дан массив целых чисел, в котором каждый элемент обозначает длину стороны треугольника.
Нужно определить максимально возможный периметр треугольника, составленного из сторон с длинами из заданного массива.
Помогите Рите скорее закончить игру и пойти спать.

Напомним, что из трёх отрезков с длинами a ≤ b ≤ c можно составить треугольник, если выполнено неравенство треугольника: c < a + b

Разберём пример:
Даны длины сторон 6, 3, 3, 2. Попробуем в качестве наибольшей стороны выбрать 6.
Неравенство треугольника не может выполниться, так как остались 3, 3, 2 — максимальная сумма из них равна 6.
Без шестёрки оставшиеся три отрезка уже образуют треугольник со сторонами 3, 3, 2. Н
еравенство выполняется: 3 < 3 + 2. Периметр равен 3 + 3 + 2 = 8.

Формат ввода
В первой строке записано количество отрезков n, 3≤ n≤ 10000.

Во второй строке записано n неотрицательных чисел, не превосходящих 10 000, –— длины отрезков.

Формат вывода
Нужно вывести одно число — наибольший периметр треугольника.

Гарантируется, что тройка чисел, которая может образовать треугольник, всегда есть.

Пример 1
Ввод
4
6 3 3 2
Вывод
8

Пример 2
Ввод
6
5 3 7 2 8 3
Вывод
20
*/

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
	"strconv"
)

func getMaxPerimeter(sides []int) int {
	slices.SortFunc(sides, func(a, b int) int {
		if n := cmp.Compare(b, a); n != 0 {
			return n
		}
		return 0
	})
	for i := 0; i < len(sides)-2; i++ {
		for j := i+1; j <= len(sides)-2; j++{
			if sides[i] < sides[j] + sides[j+1] {
				return sides[i] + sides[j] + sides[j+1]
			}
		}
	}
	return 0
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	sides := readArray(scanner)
	fmt.Print(getMaxPerimeter(sides))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}