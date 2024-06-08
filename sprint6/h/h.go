/*
H. Время выходить

Вам дан ориентированный граф. Известно, что все его вершины достижимы из вершины s=1.
Найдите время входа и выхода при обходе в глубину, производя первый запуск из вершины s.
Считайте, что время входа в стартовую вершину равно 0. Соседей каждой вершины обходите в порядке увеличения номеров.

Формат ввода
В первой строке дано число вершин n (1 ≤ n ≤ 2⋅ 105) и рёбер (0 ≤ m ≤ 2 ⋅ 105).
В каждой из следующих m строк записаны рёбра графа в виде пар (from, to), 1 ≤ from ≤ n — начало ребра, 1 ≤ to ≤ n — его конец.
Гарантируется, что в графе нет петель и кратных рёбер.

Формат вывода
Выведите n строк, в каждой из которых записана пара чисел tini, touti — время входа и выхода для вершины i.

Пример 1
Ввод
6 8
2 6
1 6
3 1
2 5
4 3
3 2
1 2
1 4

Вывод
0 11
1 6
8 9
7 10
2 3
4 5

Пример 2
Ввод
3 2
1 2
2 3

Вывод
0 5
1 4
2 3
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
    entry         []int
    leave         []int
    time          = -1
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

func DFS(v int) {
    time += 1
    entry[v] = time
    color[v] = gray
    for _, w := range adjacencyList[v] {
        if color[w] == white {
            DFS(w)
        }
    }
    time += 1
    leave[v] = time
    color[v] = black
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    adjacencyList = getAdjacencyList(scanner, n, m)
    color = make([]int, n+1)
    entry = make([]int, n+1)
    leave = make([]int, n+1)
    DFS(1)
    printTimes(n)
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

func printTimes(n int) {
    writer := bufio.NewWriter(os.Stdout)
    for i := 1; i <= n; i++ {
        writer.WriteString(strconv.Itoa(entry[i]))
        writer.WriteString(" ")
        writer.WriteString(strconv.Itoa(leave[i]))
        writer.WriteString("\n")
    }
    writer.Flush()
}