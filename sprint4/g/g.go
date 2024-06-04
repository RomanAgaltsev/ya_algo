/*
G. Сумма четвёрок

У Гоши есть любимое число S. Помогите ему найти все уникальные четвёрки чисел в массиве, которые в сумме дают заданное число S.

Формат ввода
В первой строке дано общее количество элементов массива n (0 ≤ n ≤ 1000).
Во второй строке дано целое число S.
В третьей строке задан сам массив. Каждое число является целым и не превосходит по модулю 1010.

Формат вывода
В первой строке выведите количество найденных четвёрок чисел.
В последующих строках выведите найденные четвёрки. Числа внутри одной четверки должны быть упорядочены по возрастанию.
Между собой четвёрки упорядочены лексикографически.

Пример 1
Ввод
8
10
2 3 2 4 1 10 3 0
Вывод
3
0 3 3 4
1 2 3 4
2 2 3 3

Пример 2
Ввод
6
0
1 0 -1 0 2 -2
Вывод
3
-2 -1 1 2
-2 0 0 2
-1 0 0 1

Пример 3
Ввод
5
4
1 1 1 1 1
Вывод
1
1 1 1 1
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "reflect"
    "slices"
    "strconv"
    "strings"
)

func in(result [][]int, four []int) bool {
    for _, resFour := range result {
        if reflect.DeepEqual(resFour, four) {
            return true
        }
    }
    return false
}

func getSumsOfFours(arr []int, s int) [][]int {
    result := make([][]int, 0)
    sumToPair := make(map[int][2]int)
    if len(arr) < 4 {
        return result
    }
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            sum := s - (arr[i] + arr[j])
            if pair, ok := sumToPair[sum]; ok {
                four := []int{arr[i], arr[j], pair[0], pair[1]}
                slices.Sort(four)
                if !in(result,four) {
                    result = append(result, four)
                }
            }
        }
        for j := 0; j < i; j++ {
            sumToPair[arr[i]+arr[j]] = [2]int{arr[i], arr[j]}
        }
    }
    slices.SortFunc(result, func(a,b []int) int {
        return slices.Compare(a, b)
    })
    return result
}

func main() {
    scanner := makeScanner()
    _, s := readInt(scanner), readInt(scanner)
    sumsOfFours := getSumsOfFours(readArray(scanner), s)
    fmt.Print(len(sumsOfFours))
    fmt.Print("\n")
    for _, sumOfFours := range sumsOfFours {
        printArray(sumOfFours)
    }
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