/*
L. Золото лепреконов

Лепреконы в данной задаче появились по соображениям общей морали, так как грабить банки — нехорошо.
Вам удалось заключить неплохую сделку с лепреконами, поэтому они пустили вас в своё хранилище золотых слитков.
Все слитки имеют единую пробу, то есть стоимость 1 грамма золота в двух разных слитках одинакова.
В хранилище есть n слитков, вес i-го слитка равен wi кг. У вас есть рюкзак, вместимость которого M килограмм.
Выясните максимальную суммарную массу золотых слитков, которую вы сможете унести.

Формат ввода
В первой строке дано число слитков —– натуральное число n (1 ≤ n ≤ 1000) и вместимость рюкзака –— целое число M (0 ≤ M ≤ 104).
Во второй строке записано n натуральных чисел wi (1 ≤ wi ≤ 104) -— массы слитков.

Формат вывода
Выведите единственное число — максимальную массу, которую можно забрать с собой.

Пример 1
Ввод
5 15
3 8 1 2 5
Вывод
15

Пример 2
Ввод
5 19
10 10 7 7 4
Вывод
18
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func getMaxWeight(weights []int, capacity int) int {
    dp := make([]int, capacity+1)
    dp[0] = 1
    maxWeight := 0
    for _, weight := range weights {
        for i := capacity; i >= weight; i-- {
            if dp[i-weight] == 1 {
                dp[i] = 1
                maxWeight = max(maxWeight, i)
            }
        }
    }
    return maxWeight
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    _, m := nm[0], nm[1]
    weights := readArray(scanner)
    fmt.Print(getMaxWeight(weights, m))
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