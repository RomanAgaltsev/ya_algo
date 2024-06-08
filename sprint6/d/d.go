/*
D. BFS

Задан неориентированный граф. Обойдите поиском в ширину все вершины, достижимые из заданной вершины s, и выведите их в порядке обхода, если начинать обход из s.

Формат ввода
В первой строке дано количество вершин n (1 ≤ n ≤ 105) и рёбер m (0 ≤ m ≤ 105).
Далее в m строках описаны рёбра графа. Каждое ребро описывается номерами двух вершин u и v (1 ≤ u, v ≤ n).
В последней строке дан номер стартовой вершины s (1 ≤ s ≤ n).
Гарантируется, что в графе нет петель и кратных рёбер.

Формат вывода
Выведите вершины в порядке обхода, считая что при запуске от каждой конкретной вершины её соседи
будут рассматриваться в порядке возрастания (то есть если вершина 2 соединена с 1 и 3, то сначала обход пойдёт в 1, а уже потом в 3).

Пример 1
Ввод
4 4
1 2
2 3
3 4
1 4
3

Вывод
3 2 4 1

Пример 2
Ввод
2 1
2 1
1

Вывод
1 2
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
)

var adjacencyList [][]int

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

func BFS(vertex int) {
    colors := make([]int, len(adjacencyList)+1)
    result := make([]int, 0)
    queue := []int{vertex}
    for len(queue) > 0 {
        v := queue[0]
        queue = queue[1:]
        if colors[v] == white {
            colors[v] = gray
            result = append(result, v)
            for i := 0; i < len(adjacencyList[v]); i++ {
                if colors[adjacencyList[v][i]] == white {
                    queue = append(queue, adjacencyList[v][i])
                }
            }
        }
    }
    printArray(result)
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    adjacencyList = getAdjacencyList(scanner, n, m)
    BFS(readInt(scanner))
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

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    stringInt := scanner.Text()
    res, _ := strconv.Atoi(stringInt)
    return res
}