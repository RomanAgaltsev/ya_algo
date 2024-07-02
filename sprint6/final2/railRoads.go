/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 6
Задача - B. Железные дороги

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/25070/run-report/115762081/

-- ПРИНЦИП РАБОТЫ --



-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --


*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const (
    // Константы цветов вершин графа
    white = iota // Белый = 0
    gray         // Серый = 1
    black        // Черный = 2
    // Контанты для типов дорог
    typeB = "B" // Тип дороги "B"
)

// Слайс цветов вершин
var colors []int

// Graph - структура графа
type Graph struct {
    n   int     // Количество вершин
    adj [][]int // Список смежности графа - слайс слайсов
}

// init - инициирует данные графа:
// - считывает с ввода и заполняет количество вершин графа
// - считывает с ввода карту дорог и заполняет список смежности графа
func (g *Graph) init() {
    // Создаем новый сканер
    scanner := makeScanner()
    // Считываем количество вершин графа
    n := readInt(scanner)
    // Заполняем количество вершин графа
    g.n = n
    // Создаем слайс для списка смежности и сразу присваиваем
    g.adj = make([][]int, n+1)
    // Запускаем цикл для считывания дорог и заполнения списка смежности
    for i := 1; i <= n-1; i++ {
        // Читаем строку дорог для i-ой вершины - слайс строк
        roads := readArray(scanner)
        // Каждая дорога - это ребро графа
        for j := 0; j < n-i; j++ {
            // Добавляем каждую дорогу-ребро в список смежности
            g.addRoad(i, i+j+1, roads[j])
        }
    }
}

// addRoad - добавляет дорогу-ребро в список смежности графа
func (g *Graph) addRoad(u, v int, t string) {
    // Если тип дороги = "B", меняем направление ребра
    if t == typeB {
        u, v = v, u
    }
    // Если для вершины еще нет ребер, инициируем слайс
    if len(g.adj[u]) == 0 {
        // Указываем для вершины новый слайс
        g.adj[u] = []int{v}
    } else {
        // Добавляем ребро в существующий слайс
        g.adj[u] = append(g.adj[u], v)
    }
}

// mapIsOptimal - проверяет, оптимальна ли карта дорог графа
// - true - оптимальная карта дорог
// - false - карта дорог не оптимальная
func (g *Graph) mapIsOptimal() bool {
    // Переменная для результата
    mapIsOptimal := true
    // Инициируем слайс цветов вершин графа
    // При этом, все его элементы инициируются значением по умолчанию
    // Для типа int это 0 - равно константе white
    // Получается, что весь слайс заполняется цветом white
    colors = make([]int, g.n+1)
    // Запускаем поиск циклов в графе от его вершин
    for i := 1; i < g.n; i++ {
        // Проверяем, какого цвета вершина
        if colors[i] == white {
            // Вершина белого цвета, запускаем поиск циклов - фактически, DFS
            // Карта оптимальна, если в графе нет циклов - условие должно выполняться при запуске от всех вершин
            mapIsOptimal = mapIsOptimal && !g.pathHasCycles(i)
        }
    }
    // Возвращаем результат
    return mapIsOptimal
}

// pathHasCycles - проверяет, есть ли в графе циклы от переданной вершины:
// - true - от переданной вершины цикл есть
// - false - от переданной вершины циклов нет
// Фактически, это DFS
func (g *Graph) pathHasCycles(v int) bool {
    // Переменная для результата
    pathHasCycles := false
    // Красим текущую вершину в серый цвет
    colors[v] = gray
    // Обходим смежные вершины и проверяем наличие цикла от них
    for _, w := range g.adj[v] {
        // Берем цвет смежнйо вершины
        color := colors[w]
        // Проверяем, не серого ли цвета смежная вершина
        if color == gray {
            // Смежная вершина серая - значит, ранее её уже посещали, это цикл
            return true
        }
        // Проверяем, не белого ли цвета смежная вершина
        if color == white {
            // Смежная вершина белого цвета - идем дальше в поиске цикла
            pathHasCycles = pathHasCycles || g.pathHasCycles(w)
        }
    }
    // Красим текущую вершину в черный цвет - отработали
    colors[v] = black
    // Возвращаем результат - есть ли цикл
    return pathHasCycles
}

// newGraph - конструктор нового графа
func newGraph() *Graph {
    // Просто возвращаем пустой объект
    return &Graph{}
}

func main() {
    // Создаем новый граф
    g := newGraph()
    // Инициируем граф - считываем данные с ввода, заполняем список смежности
    g.init()
    // Переменная для вывода результата - по умолчанию карта не оптимальная
    mapIsOptimal := "NO"
    // Проверяем, оптимальная карта ли карта в соответствии с задачей или нет
    if g.mapIsOptimal() {
        // Карта оптимальная
        mapIsOptimal = "YES"
    }
    // Выводим результат
    fmt.Print(mapIsOptimal)
}

func makeScanner() *bufio.Scanner {
    const maxCapacity = 3 * 1024 * 1024
    buf := make([]byte, maxCapacity)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(buf, maxCapacity)
    return scanner
}

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    stringInt := scanner.Text()
    res, _ := strconv.Atoi(stringInt)
    return res
}

func readArray(scanner *bufio.Scanner) []string {
    scanner.Scan()
    return strings.Split(scanner.Text(), "")
}