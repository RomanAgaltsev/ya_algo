/*
H. Поле с цветочками

Черепаха Кондратина путешествует по клетчатому полю из n строк и m столбцов.
В каждой клетке либо растёт цветочек, либо не растёт.
Кондратине надо добраться из левого нижнего в правый верхний угол и собрать как можно больше цветочков.

Помогите ей с этой сложной задачей и определите, какое наибольшее число цветочков она сможет собрать при условии,
что Кондратина умеет передвигаться только на одну клетку вверх или на одну клетку вправо за ход.

Формат ввода
В первой строке даны размеры поля n и m (через пробел). Оба числа лежат в диапазоне от 1 до 1000.
В следующих n строках задано поле. Каждая строка состоит из m символов 0 или 1, записанных подряд без пробелов, и завершается переводом строки.
Если в клетке записана единица, то в ней растёт цветочек.

Формат вывода
Выведите единственное число — максимальное количество цветочков, которое сможет собрать Кондратина.

Пример 1
Ввод
2 3
101
110
Вывод
3

Пример 2
Ввод
3 3
100
110
001
Вывод
2
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func getMaxPoints(field [][]int, n, m int) int {
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
    return dp[0][m-1]
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner, " ")
    n, m := nm[0], nm[1]
    field := readMatrix(scanner, n, m)
    fmt.Print(getMaxPoints(field, n, m))
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