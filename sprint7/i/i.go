/*
I. Сложное поле с цветочками

Теперь черепашке Кондратине надо узнать не только, сколько цветочков она может собрать, но и как ей построить свой маршрут для этого. Помогите ей!
Напомним, что Кондратине надо дойти от левого нижнего до правого верхнего угла, а передвигаться она умеет только вверх и вправо.

Формат ввода
В первой строке даны размеры поля n и m (через пробел). Оба числа лежат в диапазоне от 1 до 1000.
В следующих n строках задано поле. Каждая строка состоит из m символов 0 или 1 и завершается переводом строки.
Если в клетке записана единица, то в ней растет цветочек.

Формат вывода
Выведите в первой строке максимальное количество цветочков, которое сможет собрать Кондратина.
Во второй строке выведите маршрут в виде последовательности символов «U» и «R», где «U» означает передвижение вверх, а «R» – передвижение вправо.

Если возможных оптимальных путей несколько, то выведите любой.

Пример 1
Ввод
2 3
101
110
Вывод
3
URR

Пример 2
Ввод
3 3
100
110
001
Вывод
2
UURR
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

func getMaxPoints(field [][]int, n, m int) {
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            y := n - i - 1
            x := j
            down := 0
            if y+1 != n {
                down = dp[y+1][x]
            }
            left := 0
            if x != 0 {
                left = dp[y][x-1]
            }
            dp[y][x] = max(left, down) + field[y][x]
        }
    }
    fmt.Print(dp[0][m-1], "\n")
    path := make([]string, 0)
    y := 0
    x := m - 1
    for y != n || x != 0 {
        if x == 0 {
            path = append(path, "U")
            y++
            continue
        }
        down := 0
        if y+1 < n {
            down = dp[y+1][x]
        }
        left := 0
        if x > 0 {
            left = dp[y][x-1]
        }
        if down > left {
            path = append(path, "U")
            y++
        } else {
            path = append(path, "R")
            x--
        }
    }
    slices.Reverse(path)
    fmt.Print(strings.Join(path[1:], ""))
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner, " ")
    n, m := nm[0], nm[1]
    field := readMatrix(scanner, n, m)
    getMaxPoints(field, n, m)
}

func makeScanner() *bufio.Scanner {
    const maxCapacity = 3 * 1024 * 1024
    buf := make([]byte, maxCapacity)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(buf, maxCapacity)
    return scanner
}

func readArray(scanner *bufio.Scanner, sep string) []int {
    scanner.Scan()
    listString := strings.Split(scanner.Text(), sep)
    arr := make([]int, len(listString))
    for i := 0; i < len(listString); i++ {
        arr[i], _ = strconv.Atoi(listString[i])
    }
    return arr
}

func readMatrix(scanner *bufio.Scanner, rows int, cols int) [][]int {
    matrix := make([][]int, rows)
    for i := 0; i < rows; i++ {
        matrix[i] = readArray(scanner, "")
    }
    return matrix
}