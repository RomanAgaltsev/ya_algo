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

// Ошибка для вывода, если максимального остовного дерева не существует
var errGraphIsDisconnected = errors.New("Oops! I did it again")

// Edge - структура ребра графа
type Edge struct {
    u, v, w int // Две вершины и вес
}

// Graph - структура графа
type Graph struct {
    n     int      // Количество вершин графа
    edges int      // Количество ребер графа
    adj   [][]Edge // Список смежности графа - слайс слайсов
}

// addEdge - добавляет ребро в список смежности графа
func (g *Graph) addEdge(u, v, w int) {
    // Так как граф неориентированный, добавляем два ребра в список смежности
    // Добавляем ребро (u, v)
    g.adj[u] = append(g.adj[u], Edge{u, v, w})
    // Добавляем ребро (v, u)
    g.adj[v] = append(g.adj[v], Edge{v, u, w})
    // Увеличиваем количество ребер
    g.edges++
}

// max - возвращает вес максимального остовного дерева в графе:
// - если максимальное остовное дерево (МОД) существует, возвращается его вес и nil
// - если не существует, возвращается 0 и ошибка
func (g *Graph) max() (int, error) {
    // Для графа из единственной вершины вес МОД равен 0, ошибки нет
    if g.n == 1 {
        return 0, nil
    }
    // Для графа без ребер вес МОД равен 0, ошибка есть
    if g.n > 1 && g.edges == 0 {
        return 0, errGraphIsDisconnected
    }
    // Слайс для хранения посещенных вершин
    visited := make([]bool, g.n+1)
    // Создаем приоритетную очередь (она же куча) ребер
    eq := &EdgeQueue{}
    // Кладем в очередь фиктивное ребро (-1, 1) с весом 0
    heap.Push(eq, Edge{-1, 1, 0})
    // Переменная для веса МОД
    maxWeight := 0
    // Берем ребра из очереди, пока есть что брать и накапливаем вес МОД
    for eq.Len() > 0 {
        // Берем ребро - очередь приоритетная, поэтому берется ребро с максимальным весом
        e := heap.Pop(eq).(Edge)
        // Если уже посещали смежную вершину v, пропускаем
        if visited[e.v] {
            continue
        }
        // Если еще не посещали смежную вершину v, отмечаем её как посещенную
        visited[e.v] = true
        // Накапливаем вес МОД
        maxWeight += e.w
        // Получаем смежные вершины и обрабатываем в цикле
        for _, edge := range g.adj[e.v] {
            // Проверяем, не посещали ли смежную вершину v ранее
            if !visited[edge.v] {
                // Не посещали, кладем ребро в приоритетную очередь
                heap.Push(eq, edge)
            }
        }
    }
    // Проверяем, все ли вершины графа посетили в ходе обработки
    if slices.Contains(visited[1:], false) {
        // Посетили не все вершины, вернем 0 и ошибку
        return 0, errGraphIsDisconnected
    }
    // Посетили все вершины, возвращаем вес и nil
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