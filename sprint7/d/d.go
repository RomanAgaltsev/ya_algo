/*
D. Числа Фибоначчи для взрослых

Гоша практикуется в динамическом программировании — он хочет быстро считать числа Фибоначчи.
Напомним, что числа Фибоначчи определены как последовательность . F0 = F1 = 1, Fn = Fn -1 + Fn-2, n ≥ 2.
Помогите Гоше решить эту задачу.

Формат ввода
В единственной строке дано целое число n (0 ≤ n ≤ 106).

Формат вывода
Вычислите значение Fn по модулю 109 + 7 и выведите его.

Пример 1
Ввод
5
Вывод
8

Пример 2
Ввод
2
Вывод
2

Пример 3
Ввод
10
Вывод
89
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func fib1(n int) int {
    fibs := []int{1, 1}
    for i := 1; len(fibs) <= n; i++ {
        fibs = append(fibs, (fibs[i]+fibs[i-1])%1_000_000_007)
    }
    return fibs[n]
}

func fib2(n int) int {
    lhs, rhs := 0, 1
    for i := 0; i < n; i++ {
        lhs, rhs = rhs, lhs
        rhs = (lhs + rhs) % 1_000_000_007
    }
    return rhs
}

func main() {
    scanner := makeScanner()
    fmt.Print(fib1(readInt(scanner)))
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