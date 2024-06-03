/*
На вход подается строка. Нужно определить длину наибольшей подстроки, которая не содержит повторяющиеся символы.

Формат ввода
Одна строка, состоящая из строчных латинских букв. Длина строки не превосходит 10 000.

Формат вывода
Выведите натуральное число — ответ на задачу.

Пример 1
Ввод
abcabcbb
Вывод
3

Пример 2
Ввод
bbbbb
Вывод
1
*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func getMaxUniqueStringLen(line string) int {
    letterToPos := make(map[uint8]int)
    result, prev := 0, 0
    for i := 0; i < len(line); i++ {
        prev = max(prev, letterToPos[line[i]])
        letterToPos[line[i]] = i + 1
        result = max(result, i+1-prev)
    }
    return result
}

func main() {
    scanner := makeScanner()
    fmt.Print(getMaxUniqueStringLen(readLine(scanner)))
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