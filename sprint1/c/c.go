/*
C. Соседи

Дана матрица. Нужно написать функцию, которая для элемента возвращает всех его соседей.
Соседним считается элемент, находящийся от текущего на одну ячейку влево, вправо, вверх или вниз.
Диагональные элементы соседними не считаются.

Например, в матрице A соседними элементами для (0, 0) будут 2 и 0. А для (2, 1) — 1, 2, 7, 7.

Формат ввода
В первой строке задано n — количество строк матрицы.
Во второй — количество столбцов m. Числа m и n не превосходят 1000.
В следующих n строках задана матрица.
Элементы матрицы — целые числа, по модулю не превосходящие 1000.
В последних двух строках записаны координаты элемента, соседей которого нужно найти.
Индексация начинается с нуля.

Формат вывода
Напечатайте нужные числа в возрастающем порядке через пробел.

Пример 1
Ввод
4
3
1 2 3
0 2 6
7 4 1
2 7 0
3
0
Вывод
7 7

Пример 2
Ввод
4
3
1 2 3
0 2 6
7 4 1
2 7 0
0
0
Вывод
0 2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getNeighbour(matrix [][]int, row int, col int) (int, bool) {
	if row > (len(matrix) - 1) || row < 0 {
		return 0, false
	}
	if col > (len(matrix[row]) - 1) || col < 0 {
		return 0, false
	}
	return matrix[row][col], true
}

func getNeighbours(matrix [][]int, row int, col int) []int {
	neighbours := make([]int, 0)
	if neighbour, found := getNeighbour(matrix, row, col+1); found {
		neighbours = append(neighbours, neighbour)
	}
	if neighbour, found := getNeighbour(matrix, row-1, col); found {
		neighbours = append(neighbours, neighbour)
	}
	if neighbour, found := getNeighbour(matrix, row+1, col); found {
		neighbours = append(neighbours, neighbour)
	}
	if neighbour, found := getNeighbour(matrix, row, col-1); found {
		neighbours = append(neighbours, neighbour)
	}
	sort.Ints(neighbours)
	return neighbours
}

func main() {
	scanner := makeScanner()
	rows := readInt(scanner)
	cols := readInt(scanner)
	matrix := readMatrix(scanner, rows, cols)
	rowId := readInt(scanner)
	colId := readInt(scanner)
	neighbours := getNeighbours(matrix, rowId, colId)
	for _, elem := range neighbours {
		fmt.Print(elem)
		fmt.Print(" ")
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
	writer.Flush()
}

func readMatrix(scanner *bufio.Scanner, rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = readArray(scanner)
	}
	return matrix
}
