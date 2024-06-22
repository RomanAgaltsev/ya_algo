/*
I. Разные деревья поиска

Ребятам стало интересно, сколько может быть различных деревьев поиска, содержащих в своих узлах все уникальные числа от 1 до n.
Помогите им найти ответ на этот вопрос.

Формат ввода
В единственной строке задано число n. Оно не превосходит 15.

Формат вывода
Нужно вывести число, равное количеству различных деревьев поиска, в узлах которых могут быть размещены числа от 1 до n включительно.

Пример 1
Ввод
2
Вывод
2

Пример 2
Ввод
3
Вывод
5

Пример 3
Ввод
4
Вывод
14
*/

package main

import (
    "fmt"
    "strconv"
)

func getBSTNumber(n int) int {
    if n <= 1 {
        return 1
    }
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for i := 2; i <= n; i++ {
        for j := 0; j < i; j++ {
            dp[i] += dp[j] * dp[i-j-1]
        }
    }
    return dp[n]
}

func main() {
    fmt.Println(getBSTNumber(readInt()))
}

func readInt() int {
    var aString string
    fmt.Scan(&aString)
    a, _ := strconv.Atoi(aString)
    return a
}