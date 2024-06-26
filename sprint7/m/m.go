/*
M. Рюкзак

Тимофей решил отправиться в поход. Ему надо собрать рюкзак.
Так как поход долгий и трудный, необходимо подбирать вещи вдумчиво.
Каждому предмету Тимофей присвоил условную значимость: она равна ci для предмета с номером i.
Также каждый предмет весит mi килограммов. А грузоподъёмность рюкзака  равна M килограмм.
Найдите максимальную суммарную значимость предметов, которые Тимофей может взять с собой, не порвав рюкзак, и укажите, как набрать эту значимость.

Формат ввода
В первой строке вводится число предметов n, не превышающее 100 и грузоподъемность M, не превышающая 104.
Далее следуют описания предметов по одному в строке.
Каждый предмет описывается парой mi, ci, оба числа не превосходят 100 по модулю.

Формат вывода
Выведите в первой строке единственное число — сколько предметов надо взять.
Во второй строке перечислите их номера (нумерация с единицы).
Если ответов несколько, то выведите любой.

Пример
Ввод
4 6
2 7
4 2
1 5
2 1
Вывод
3
4 3 1
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Item struct {
    weight, cost int
}

func getMaxCost(items []Item, capacity int) []int {
    dp := make([][]int, len(items)+1)
    for i := 0; i < len(items)+1; i++ {
        dp[i] = make([]int, capacity+1)
    }
    for i := 1; i <= len(items); i++ {
        for j := 0; j <= capacity; j++ {
            if j >= items[i-1].weight {
                dp[i][j] = max(dp[i-1][j], dp[i-1][j-items[i-1].weight]+items[i-1].cost)
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    result := make([]int, 0)
    i := len(items)
    j := capacity
    for dp[i][j] != 0 {
        if dp[i-1][j] == dp[i][j] {
            i--
            continue
        }
        if dp[i][j-1] == dp[i][j] {
            j--
            continue
        }
        result = append(result, i)
        j -= items[i-1].weight
        i--
    }
    return result
}

func main() {
    scanner := makeScanner()
    nm := readArray(scanner)
    n, m := nm[0], nm[1]
    items := readItems(scanner, n)
    res := getMaxCost(items, m)
    fmt.Print(len(res), "\n")
    printArray(res)
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

func readItems(scanner *bufio.Scanner, n int) []Item {
    items := make([]Item, n)
    for i := 0; i < n; i++ {
        items[i] = readItem(scanner)
    }
    return items
}

func readItem(scanner *bufio.Scanner) Item {
    scanner.Scan()
    listItem := strings.Split(scanner.Text(), " ")
    cost, _ := strconv.Atoi(listItem[0])
    weight, _ := strconv.Atoi(listItem[1])
    return Item{
        cost,
        weight,
    }
}