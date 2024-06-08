/*
E. Компоненты связности

Вам дан неориентированный граф. Найдите его компоненты связности.

Формат ввода
В первой строке дано количество вершин n (1≤ n ≤ 105) и рёбер m (0 ≤ m ≤ 2 ⋅ 105).
В каждой из следующих m строк записано по ребру в виде пары вершин 1 ≤ u, v ≤ n.
Гарантируется, что в графе нет петель и кратных рёбер.

Формат вывода
Выведите все компоненты связности в следующем формате: в первой строке выведите общее количество компонент.
Затем на отдельных строках выведите вершины каждой компоненты, отсортированные по возрастанию номеров.
Компоненты между собой упорядочивайте по номеру первой вершины.

Пример 1
Ввод
6 3
1 2
6 5
2 3
Вывод
3
1 2 3
4
5 6

Пример 2
Ввод
2 0
Вывод
2
1
2
Пример 3
Ввод
4 3
2 3
2 1
4 3
Вывод
1
1 2 3 4
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

var (
    adjacencyList  [][]int
    components     [][]int
    color          []int
    componentCount = 0
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
    return adjacencyList
}

func DFS(v int) {
    color[v] = componentCount
    if len(components[componentCount]) == 0 {
        components[componentCount] = []int{v}
    } else {
        components[componentCount] = append(components[componentCount], v)
    }
    for _, w := range adjacencyList[v] {
        if color[w] == 0 {
            DFS(w)
        }
    }
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    adjacencyList = getAdjacencyList(scanner, n, m)
    components = make([][]int, n+1)
    color = make([]int, n+1)
    for i := 1; i < len(color); i++ {
        if color[i] == 0 {
            componentCount += 1
            DFS(i)
        }
    }
    fmt.Print(componentCount, "\n")
    for i := 1; i < len(components); i++ {
        if len(components[i]) > 0 {
            slices.Sort(components[i])
            printArray(components[i])
        }
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