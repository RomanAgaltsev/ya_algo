/*
G. Работа из дома

Вася реализовал функцию, которая переводит целое число из десятичной системы в двоичную.
Но, кажется, она получилась не очень оптимальной.
Попробуйте написать более эффективную программу.
Не используйте встроенные средства языка по переводу чисел в бинарное представление.

Формат ввода
На вход подаётся целое число в диапазоне от 0 до 10000.

Формат вывода
Выведите двоичное представление этого числа.

Пример 1
Ввод
5
Вывод
101

Пример 2
Ввод
14
Вывод
1110
*/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func GetBinaryNumber(n int) []int {
	res := make([]int, 0)
	if n == 0 {
		return append(res, 0)
	}
	for n > 0 {
		res = append(res, n%2)
		n = n / 2
	}
	for i, j := 0, len(res)-1; i < j; i,j = i+1,j-1 {
		res[i], res[j] = res[j],res[i]
	}
	return res
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	binaryNumber := GetBinaryNumber(n)
	printArray(binaryNumber)
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
	}
	writer.Flush()
}
