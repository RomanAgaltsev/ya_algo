/*
J. Путешествие

Гоша решил отправиться в турне по островам Алгосского архипелага.
Туристическая программа состоит из последовательного посещения n достопримечательностей.
У i-й достопримечательности есть свой рейтинг ri.

Впечатление от i-й достопримечательности равно её рейтингу ri.
Гоша хочет, чтобы его впечатление от каждой новой посещённой достопримечательности было сильнее, чем от предыдущей.
Ради этого он даже готов пропустить некоторые места в маршруте –— в случае, если они нарушают этот порядок плавного возрастания.

Помогите Гоше и найдите наибольшую возрастающую подпоследовательность в массиве рейтингов ri.

Формат ввода
В первой строке дано натуральное число n (1 ≤ n ≤ 3 ⋅ 103) — сколько различных туристических мест есть в программе.
Во второй строке дано n натуральных чисел через пробел — рейтинги этих достопримечательностей ri (1 ≤ ri ≤ 109).

Формат вывода
Сначала в отдельной строке выведите длину найденной подпоследовательности.
В следующей строке выведите номера достопримечательностей, которые образуют эту подпоследовательность.

Пример 1
Ввод
5
4 2 9 1 13
Вывод
3
1 3 5

Пример 2
Ввод
6
1 2 4 8 16 32
Вывод
6
1 2 3 4 5 6
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

func getMaxArr(rates []int) []int {
    dp := make([]int, len(rates))
    for i := 0; i < len(dp); i++ {
        dp[i] = 1
    }
    //    ancestors := make([]int, len(rates))
    //    for i := range ancestors {
    //        ancestors[i] = -1
    //    }
    maxRate := 0
    for i := 1; i < len(rates); i++ {
        for j := 0; j < i; j++ {
            if rates[i] > rates[j] {
                dp[i] = max(dp[i], 1+dp[j])
                maxRate = max(maxRate, dp[i])
                //ancestors[i] = j
            }
        }

    }
    result := make([]int, 0)
    for i := len(dp) - 1; i >= 0; i-- {
        if dp[i] == maxRate {
            if i > 0 && dp[i] > dp[i-1] {
                result = append(result, i+1)
                maxRate--
            } else if i == 0 {
                result = append(result, i+1)
                maxRate--
            }
        }
    }
    slices.Reverse(result)
    return result
    //    maxRatePos := -1
    //    for i := len(dp)-1; i >= 0; i -- {
    //        if dp[i] == maxRate {
    //            maxRatePos = i
    //            break
    //        }
    //    }
    //    fmt.Print(maxRatePos)
    //result := []int{maxRatePos}
    // return ancestors
}

func main() {
    scanner := makeScanner()
    _, r := readInt(scanner), readArray(scanner)
    maxArr := getMaxArr(r)
    fmt.Print(len(maxArr), "\n")
    printArray(maxArr)
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
    writer.WriteString("\n")
    writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    stringInt := scanner.Text()
    res, _ := strconv.Atoi(stringInt)
    return res
}