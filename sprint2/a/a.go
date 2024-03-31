/*
A. Мониторинг

Алла получила задание, связанное с мониторингом работы различных серверов.
Требуется понять, сколько времени обрабатываются определённые запросы на конкретных серверах.
Эту информацию нужно хранить в матрице, где номер столбца соответствуют идентификатору запроса,
а номер строки — идентификатору сервера. Алла перепутала строки и столбцы местами. С каждым бывает.
Помогите ей исправить баг.

Есть матрица размера m × n. Нужно написать функцию, которая её транспонирует.
Транспонированная матрица получается из исходной заменой строк на столбцы.

Формат ввода
В первой строке задано число n — количество строк матрицы.
Во второй строке задано m — число столбцов, m и n не превосходят 1000.
В следующих n строках задана матрица. Числа в ней не превосходят по модулю 1000.

Формат вывода
Напечатайте транспонированную матрицу в том же формате, который задан во входных данных.
Каждая строка матрицы выводится на отдельной строке, элементы разделяются пробелами.

Пример 1
Ввод
4
3
1 2 3
0 2 6
7 4 1
2 7 0
Вывод
1 0 7 2
2 2 4 7
3 6 1 0

Пример 2
Ввод
9
5
-7 -1 0 -4 -9
5 -1 2 2 9
3 1 -8 -1 -7
9 0 8 -8 -1
2 4 5 2 8
-7 10 0 -4 -8
-3 10 -7 10 3
1 6 -7 -5 9
-1 9 9 1 9
Вывод
-7 5 3 9 2 -7 -3 1 -1
-1 -1 1 0 4 10 10 6 9
0 2 -8 8 5 0 -7 -7 9
-4 2 -1 -8 2 -4 10 -5 1
-9 9 -7 -1 8 -8 3 9 9

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func newMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func transposeMatrix(matrix [][]int, rows, cols int) [][]int {
	transposedMatrix := newMatrix(cols, rows)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++{
			transposedMatrix[col][row] = matrix[row][col]
		}
	}
	return transposedMatrix
}

func main() {
	scanner := makeScanner()
	rows := readInt(scanner)
	cols := readInt(scanner)
	matrix := readMatrix(scanner, rows, cols)
	transposedMatrix := transposeMatrix(matrix, rows, cols)
	for _, line := range transposedMatrix {
		printArray(line)
	}
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.WriteString("\n")
	writer.Flush()
}

func readMatrix(scanner *bufio.Scanner, rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = readArray(scanner)
	}
	return matrix
}