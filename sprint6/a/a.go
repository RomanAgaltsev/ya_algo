/*
A. Построить список смежности

Алла пошла на стажировку в студию графического дизайна, где ей дали такое задание:
для очень большого числа ориентированных графов преобразовать их список рёбер в список смежности.
Чтобы побыстрее решить эту задачу, она решила автоматизировать процесс.

Помогите Алле написать программу, которая по списку рёбер графа будет строить его список смежности.

Формат ввода
В первой строке дано число вершин n (1 ≤ n ≤ 100) и число ребер m (1 ≤ m ≤ n(n-1)).
В следующих m строках заданы ребра в виде пар вершин (u,v), если ребро ведет от u к v.

Формат вывода
Выведите информацию о рёбрах, исходящих из каждой вершины.
В строке i надо написать число рёбер, исходящих из вершины i, а затем перечислить вершины, в которые ведут эти рёбра –— в порядке возрастания их номеров.

Пример
Ввод
5 3
1 3
2 3
5 2

Вывод
1 3
1 3
0
0
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

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    adjacencyList := make([][]int, n+1)
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
        adjacencyList[i] = append([]int{len(adjacencyList[i])}, adjacencyList[i]...)
        printArray(adjacencyList[i])
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