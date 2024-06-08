/*
J. Топологическая сортировка

Дан ациклический ориентированный граф (так называемый DAG, directed acyclic graph).
Найдите его топологическую сортировку, то есть выведите его вершины в таком порядке, что все рёбра графа идут слева направо.
У графа может быть несколько подходящих перестановок вершин. Вам надо найти любую топологическую сортировку.

Формат ввода
В первой строке даны два числа – количество вершин n (1 ≤ n ≤ 105) и количество рёбер m (0 ≤ m ≤ 105).
В каждой из следующих m строк описаны рёбра по одному на строке.
Каждое ребро представлено парой вершин (from, to), 1≤ from, to ≤ n, соответственно номерами вершин начала и конца.

Формат вывода
Выведите номера вершин в требуемом порядке.

Пример 1
Ввод	Вывод
5 3
3 2
3 4
2 5
Вывод
1 3 2 4 5

Пример 2
Ввод
6 3
6 4
4 1
5 1
Вывод
2 3 5 6 4 1

Пример 3
Ввод
4 0
Вывод
1 2 3 4

*/

package main

import (
    "bufio"
    "os"
    "slices"
    "strconv"
    "strings"
)

const (
    white = iota
    gray
    black
)

var (
    adjacencyList [][]int
    color         []int
    order         []int
)

func getAdjacencyList(scanner *bufio.Scanner, n, m int) [][]int {
    adjacencyList = make([][]int, n+1)
    for i := 0; i < m; i++ {
        uv := readArray(scanner)
        u, v := uv[0], uv[1]
        if len(adjacencyList[u]) == 0 {
            adjacencyList[u] = []int{v}
        } else {
            adjacencyList[u] = append(adjacencyList[u], v)
        }
    }
    for i := 1; i <= n; i++ {
        slices.Sort(adjacencyList[i])
    }
    return adjacencyList
}

func topSort(v int) {
    color[v] = gray
    for _, w := range adjacencyList[v] {
        if color[w] == white {
            topSort(w)
        }
    }
    color[v] = black
    order = append(order, v)
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    adjacencyList = getAdjacencyList(scanner, n, m)
    color = make([]int, n+1)
    for i := 1; i < len(color); i++ {
        if color[i] == white {
            topSort(i)
        }
    }
    printArray(order)
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

func printArray(arr []int) {
    writer := bufio.NewWriter(os.Stdout)
    for i := len(arr)-1; i >= 0 ; i-- {
        writer.WriteString(strconv.Itoa(arr[i]))
        writer.WriteString(" ")
    }
    writer.WriteString("\n")
    writer.Flush()
}