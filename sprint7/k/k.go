/*
K. Гороскопы

В мире последовательностей нет гороскопов.
Поэтому когда две последовательности хотят понять, могут ли они счастливо жить вместе,
они оценивают свою совместимость как длину их наибольшей общей подпоследовательности.
Подпоследовательность получается из последовательности удалением некоторого (возможно, нулевого) числа элементов.
То есть элементы сохраняют свой относительный порядок, но не обязаны изначально идти подряд.

Найдите наибольшую общую подпоследовательность двух одиноких последовательностей и выведите её!

Формат ввода
В первой строке дано число n — количество элементов в первой последовательности (1 ≤ n ≤ 1000).
Во второй строке даны n чисел ai (0 ≤ |ai| ≤ 109) — элементы первой последовательности.
Аналогично в третьей строке дано m (1 ≤ m ≤ 1000) — число элементов второй последовательности.
В четвертой строке даны элементы второй последовательности через пробел bi (0 ≤ |bi| ≤ 109).

Формат вывода
Сначала выведите длину найденной наибольшей общей подпоследовательности, во второй строке выведите индексы элементов первой последовательности,
которые в ней участвуют, в третьей строке — индексы элементов второй последовательности.
Нумерация индексов с единицы, индексы должны идти в корректном порядке.

Если возможных НОП несколько, то выведите любую.

Пример 1
Ввод
5
4 9 2 4 6
7
9 4 0 0 2 8 4

Вывод
3
1 3 4
2 5 7

Пример 2
Ввод
4
1 1 1 1
2
2 2

Вывод
0
Пример 3
Ввод
8
1 2 1 9 1 2 1 9
5
9 9 1 9 9

Вывод
3
3 4 8
3 4 5
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

func getMaxSubarray(a, b []int) ([]int, []int) {
    dp := make([][]int, len(a)+1)
    for i := 0; i < len(a)+1; i++ {
        dp[i] = make([]int, len(b)+1)
    }
    for i := 1; i <= len(a); i++ {
        for j := 1; j <= len(b); j++ {
            if a[i-1] == b[j-1] {
                dp[i][j] = 1 + dp[i-1][j-1]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    resA := make([]int, 0)
    resB := make([]int, 0)
    posA := len(a)
    posB := len(b)
    for posA > 0 && posB > 0 {
        if a[posA-1] == b[posB-1] {
            posA--
            posB--
            resA = append(resA, posA+1)
            resB = append(resB, posB+1)
        } else {
            if dp[posA-1][posB] >= dp[posA][posB-1] {
                posA--
            } else {
                posB--
            }
        }
    }
    slices.Reverse(resA)
    slices.Reverse(resB)
    return resA, resB
}

func main() {
    scanner := makeScanner()
    _ = readInt(scanner)
    a := readArray(scanner)
    _ = readInt(scanner)
    b := readArray(scanner)
    resA, resB := getMaxSubarray(a, b)
    fmt.Print(len(resA), "\n")
    if len(resA) > 0 {
        printArray(resA)
        fmt.Print("\n")
        printArray(resB)
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
    writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    stringInt := scanner.Text()
    res, _ := strconv.Atoi(stringInt)
    return res
}