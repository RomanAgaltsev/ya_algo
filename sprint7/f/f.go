/*
F. Прыжки по лестнице

Алла хочет доказать, что она умеет прыгать вверх по лестнице быстрее всех.
На этот раз соревнования будут проходить на специальной прыгательной лестнице.
С каждой её ступеньки можно прыгнуть вверх на любое расстояние от 1 до k.
Число k придумывает Алла.

Гоша не хочет проиграть, поэтому просит вас посчитать количество способов допрыгать от первой ступеньки до n-й.
Изначально все стоят на первой ступеньке.

Формат ввода
В единственной строке даны два числа — n и k (1 ≤ n ≤ 1000, 1 ≤ k ≤ n).

Формат вывода
Выведите количество способов по модулю 109 + 7.

Пример 1
Ввод
6 3
Вывод
13

Пример 2
Ввод
7 7
Вывод
32

Пример 3
Ввод
2 2
Вывод
1
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := makeScanner()
    nk := readArray(scanner)
    n, k := nk[0], nk[1]
    dp := make([]int, n+1)
    dp[0] = 0
    dp[1] = 1
    for i := 2; i < len(dp); i++ {
        if k >= i {
            for j := 1; j < i; j++ {
                dp[i] += dp[j]
            }
        } else {
            for j := i - k; j < i; j++ {
                dp[i] += dp[j]
            }
        }
        dp[i] = dp[i] % 1000000007
    }
    fmt.Print(dp[n])
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