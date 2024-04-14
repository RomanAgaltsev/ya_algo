/*
H. Большое число

Вечером ребята решили поиграть в игру «Большое число».
Даны числа. Нужно определить, какое самое большое число можно из них составить.

Формат ввода
В первой строке записано n — количество чисел. Оно не превосходит 100.
Во второй строке через пробел записаны n неотрицательных чисел, каждое из которых не превосходит 1000.

Формат вывода
Нужно вывести самое большое число, которое можно составить из данных чисел.

Пример 1
Ввод
3
15 56 2
Вывод
56215

Пример 2
Ввод
3
1 783 2
Вывод
78321

Пример 3
Ввод
5
2 4 5 2 10
Вывод
542210
*/

package main

import (
	"bufio"
	"cmp"
	"os"
	"slices"
	"strings"
	"strconv"
)

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	slices.SortFunc(arr, func(a, b string) int {
		abSum, _ := strconv.Atoi(b + a)
		baSum, _ := strconv.Atoi(a + b)
		if n := cmp.Compare(abSum, baSum); n != 0 {
			return n
		}
		return 0
	})
	printArray(arr)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	return arr
}

func printArray(arr []string) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(arr[i])
	}
	writer.WriteString("\n")
	writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
