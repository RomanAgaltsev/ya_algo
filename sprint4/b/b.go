/*
B. Соревнование

Жители Алгосов любят устраивать турниры по спортивному программированию.
Все участники разбиваются на пары и соревнуются друг с другом.
А потом два самых сильных программиста встречаются в финальной схватке, которая состоит из нескольких раундов.
Если в очередном раунде выигрывает первый участник, в таблицу с результатами записывается 0, если второй, то 1.
Ничьей в раунде быть не может.

Нужно определить наибольший по длине непрерывный отрезок раундов, по результатам которого суммарно получается ничья.
Например, если дана последовательность 0 0 1 0 1 1 1 0 0 0, то раунды с 2-го по 9-й (нумерация начинается с единицы) дают ничью.

Формат ввода
В первой строке задаётся n (0 ≤ n ≤ 105) — количество раундов.
Во второй строке через пробел записано n чисел –— результаты раундов.
Каждое число равно либо 0, либо 1.

Формат вывода
Выведите длину найденного отрезка.

Пример 1
Ввод
2
0 1
Вывод
2

Пример 2
Ввод
3
0 1 0
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

func getMaxDraw(arr []string) int {
    count, maxDraw := 0, 0
    draws := map[int]int{0: -1}
    for i, num := range arr {
        if num == "1" {
            count++
        } else {
            count--
        }
        if draw, ok := draws[count]; ok {
            maxDraw = max(maxDraw, i-draw)
        } else {
            draws[count] = i
        }
    }
    return maxDraw
}

func main() {
    scanner := makeScanner()
    readInt(scanner)
    arr := readArray(scanner)
    fmt.Print(getMaxDraw(arr))
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
    return strings.Split(scanner.Text(), " ")
}