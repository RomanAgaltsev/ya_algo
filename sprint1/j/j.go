/*
J. Факторизация

Основная теорема арифметики говорит:
любое число раскладывается на произведение простых множителей единственным образом, с точностью до их перестановки.
Например:
Число 8 можно представить как 2 × 2 × 2.
Число 50 –— как 2 × 5 × 5 (или 5 × 5 × 2, или 5 × 2 × 5). Три варианта отличаются лишь порядком следования множителей.
Разложение числа на простые множители называется факторизацией числа.

Напишите программу, которая производит факторизацию переданного числа.

Формат ввода
В единственной строке дано число n (2 ≤ n ≤ 109), которое нужно факторизовать.

Формат вывода
Выведите в порядке неубывания простые множители, на которые раскладывается число n.

Пример 1
Ввод
8
Вывод
2 2 2

Пример 2
Ввод
13
Вывод
13

Пример 3
Ввод
100
Вывод
2 2 5 5

*/

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func factorize(number int) []int {
	factors := make([]int, 0)
	for number > 0 {
		if number == 2 {
			factors = append(factors, 2)
			number = 0
			break
		}
		if number % 2 == 0 {
			factors = append(factors, 2)
			number /= 2
			continue
		}
		hasPrimeFactor := false
		for i := 3; i*i <= number; i += 2 {
			if number%i == 0 {
				factors = append(factors, i)
				number = number / i
				hasPrimeFactor = true
				break
			}
		}
		if !hasPrimeFactor {
			factors = append(factors, number)
			break
		}
	}
	sort.Ints(factors)
	return factors
}

func main() {
	scanner := makeScanner()
	number := readInt(scanner)
	factorization := factorize(number)
	printArray(factorization)
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
