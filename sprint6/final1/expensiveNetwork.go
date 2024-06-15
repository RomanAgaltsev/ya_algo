/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 6
Задача - A. Дорогая сеть

Отчеты:
- Ревью 1 -

-- ПРИНЦИП РАБОТЫ --

https://pkg.go.dev/container/heap#example-package-PriorityQueue

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --


*/

package main

import (
    "bufio"
    "container/heap"
    "errors"
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
)

var errGraphIsDisconnected = errors.New("Oops! I did it again")

type Edge struct {
    u, v, w int
}

type Graph struct {
    n     int
    edges []Edge
    adj   [][]Edge
}

func (g *Graph) addEdge(u, v, w int) {
    g.adj[u] = append(g.adj[u], Edge{u, v, w})
    g.adj[v] = append(g.adj[v], Edge{v, u, w})
    g.edges = append(g.edges, Edge{u, v, w})
}

func (g *Graph) max() (int, error) {
    if g.n == 1 {
        return 0, nil
    }
    if g.n > 1 && len(g.edges) == 0 {
        return 0, errGraphIsDisconnected
    }
    visited := make([]bool, g.n+1)
    eq := &EdgeQueue{}
    heap.Push(eq, Edge{-1, 1, 0})
    maxWeight := 0
    for eq.Len() > 0 {
        e := heap.Pop(eq).(Edge)
        if visited[e.v] {
            continue
        }
        visited[e.v] = true
        maxWeight += e.w
        for _, edge := range g.adj[e.v] {
            if !visited[edge.v] {
                heap.Push(eq, edge)
            }
        }
    }
    if slices.Contains(visited[1:], false) {
        return 0, errGraphIsDisconnected
    }
    return maxWeight, nil
}

type EdgeQueue []Edge

func (eq EdgeQueue) Len() int { return len(eq) }

func (eq EdgeQueue) Less(i, j int) bool { return eq[i].w > eq[j].w }

func (eq EdgeQueue) Swap(i, j int) { eq[i], eq[j] = eq[j], eq[i] }

func (eq *EdgeQueue) Push(x any) {
    *eq = append(*eq, x.(Edge))
}

func (eq *EdgeQueue) Pop() any {
    old := *eq
    n := len(old)
    edge := old[n-1]
    *eq = old[0 : n-1]
    return edge
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    g := &Graph{
        n:   n,
        adj: make([][]Edge, 2*m),
    }
    for i := 0; i < m; i++ {
        uvw := readArray(scanner)
        g.addEdge(uvw[0], uvw[1], uvw[2])
    }
    maxWeight, err := g.max()
    if err != nil {
        fmt.Print(err)
        return
    }
    fmt.Print(maxWeight)
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