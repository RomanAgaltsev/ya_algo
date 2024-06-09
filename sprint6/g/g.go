/*
G. Максимальное расстояние

Под расстоянием между двумя вершинами в графе будем понимать длину кратчайшего пути между ними в рёбрах.
Для данной вершины s определите максимальное расстояние от неё до другой вершины неориентированного графа.

Формат ввода
В первой строке дано количество вершин n (1 ≤ n ≤ 105) и рёбер m (0 ≤ m ≤ 105).
Далее в m строках описаны рёбра графа. Каждое ребро описывается номерами двух вершин u и v (1 ≤ u, v ≤ n).
В последней строке дан номер вершины s (1 ≤ s ≤ n). Гарантируется, что граф связный и что в нём нет петель и кратных рёбер.

Формат вывода
Выведите длину наибольшего пути от s до одной из вершин графа.

Пример 1
Ввод
5 4
2 1
4 5
4 3
3 2
2
Вывод
3

Пример 2
Ввод
3 3
3 1
1 2
2 3
1
Вывод
1

Пример 3
Ввод
6 8
6 1
1 3
5 1
3 5
3 4
6 5
5 2
6 2
4
Вывод
3
*/

package main

import (
    "bufio"
    "fmt"
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
    previous      []int
    distance      []int
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
        if len(adjacencyList[v]) == 0 {
            adjacencyList[v] = []int{u}
        } else {
            adjacencyList[v] = append(adjacencyList[v], u)
        }
    }
    for i := 1; i <= n; i++ {
        slices.Sort(adjacencyList[i])
    }
    return adjacencyList
}

func BFS(s int) {
    planned := []int{s}
    color[s] = gray
    distance[s] = 0
    previous[s] = -1
    for len(planned) > 0 {
        u := planned[0]
        planned = planned[1:]
        for _, v := range adjacencyList[u] {
            if color[v] == white {
                distance[v] = distance[u] + 1
                previous[v] = u
                color[v] = gray
                planned = append(planned, v)
            }
        }
        color[u] = black
    }
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    adjacencyList = getAdjacencyList(scanner, n, m)
    color = make([]int, n+1)
    previous = make([]int, n+1)
    distance = make([]int, n+1)
    BFS(readInt(scanner))
    fmt.Print(slices.Max(distance))
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