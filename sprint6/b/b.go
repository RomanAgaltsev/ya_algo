/*
B. Перевести список ребер в матрицу смежности

Алла успешно справилась с предыдущим заданием, и теперь ей дали новое.
На этот раз список рёбер ориентированного графа надо переводить в матрицу смежности.
Конечно же, Алла попросила вас помочь написать программу для этого.

Формат ввода
В первой строке дано число вершин n (1 ≤ n ≤ 100) и число рёбер m (1 ≤ m ≤ n(n-1)).
В следующих m строках заданы ребра в виде пар вершин (u,v), если ребро ведет от u к v.

Формат вывода
Выведите матрицу смежности n на n.
На пересечении i-й строки и j-го столбца стоит единица, если есть ребро, ведущее из i в j.

Пример
Ввод
5 3
1 3
2 3
5 2

Вывод
0 0 1 0 0 
0 0 1 0 0 
0 0 0 0 0 
0 0 0 0 0 
0 1 0 0 0
*/

package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    adjacencyMatrix := make([][]int, n+1)
    for i := 1; i <= n; i++ {
        adjacencyMatrix[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        uv := readArray(scanner)
        u, v := uv[0], uv[1]
        adjacencyMatrix[u][v-1] = 1
    }
    for i := 1; i <= n; i++ {
        printArray(adjacencyMatrix[i])
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

func printArray(arr []int) {
    writer := bufio.NewWriter(os.Stdout)
    for i := 0; i < len(arr); i++ {
        writer.WriteString(strconv.Itoa(arr[i]))
        writer.WriteString(" ")
    }
    writer.WriteString("\n")
    writer.Flush()
}