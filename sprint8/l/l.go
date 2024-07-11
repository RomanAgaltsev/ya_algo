/*
L. Подсчёт префикс-функции

В этой задаче вам необходимо посчитать префикс-функцию для заданной строки.

Формат ввода
На вход подаётся строка, состоящая из строчных латинских букв. Длина строки не превосходит 106.

Формат вывода
Если длина входной строки L, то выведите через пробел L целых неотрицательных чисел — массив значений префикс-функции исходной строки.

Пример 1
Ввод
abracadabra
Вывод
0 0 0 1 0 1 0 1 2 3 4

Пример 2
Ввод
xxzzxxz
Вывод
0 1 0 0 1 2 3

Пример 3
Ввод
aaaaa
Вывод
0 1 2 3 4

*/

package main

import (
    "bufio"
    "os"
    "strconv"
)

func prefixFunction(s string) []int {
    n := len(s)
    pi := make([]int, n)
    pi[0] = 0
    for i := 1; i < n; i++ {
        k := pi[i-1]
        for k > 0 && s[k] != s[i] {
            k = pi[k-1]
        }
        if s[k] == s[i] {
            k++
        }
        pi[i] = k
    }
    return pi
}

func main() {
    scanner := makeScanner()
    s := readLine(scanner)
    printArray(prefixFunction(s))
}

func makeScanner() *bufio.Scanner {
    const maxCapacity = 3 * 1024 * 1024
    buf := make([]byte, maxCapacity)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(buf, maxCapacity)
    return scanner
}

func readLine(scanner *bufio.Scanner) string {
    scanner.Scan()
    return scanner.Text()
}

func printArray(arr []int) {
    writer := bufio.NewWriter(os.Stdout)
    for i := 0; i < len(arr); i++ {
        writer.WriteString(strconv.Itoa(arr[i]))
        writer.WriteString(" ")
    }
    writer.Flush()
}
