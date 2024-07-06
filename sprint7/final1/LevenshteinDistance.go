/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 7
Задача - A. Расстояние по Левенштейну

Отчеты:
- Ревью 1 -

-- ПРИНЦИП РАБОТЫ --

https://en.wikipedia.org/wiki/Levenshtein_distance
https://en.wikipedia.org/wiki/Levenshtein_distance#Iterative_with_two_matrix_rows

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func LevenshteinDistance(s, t string) int {
    if len(s) < len(t) {
        s, t = t, s
    }
    lenS, lenT := len(s), len(t)
    prev := make([]int, lenS+1)
    curr := make([]int, lenS+1)
    for i := 0; i <= lenS; i++ {
        prev[i] = i
    }
    for i := 1; i < lenT+1; i++ {
        curr[0] = i + 1
        for j := 1; j < lenS+1; j++ {
            deletionCost := prev[j] + 1
            insertionCost := curr[j-1] + 1
            substitutionCost := 0
            if t[i-1] == s[j-1] {
                substitutionCost = prev[j-1]
            } else {
                substitutionCost = prev[j-1] + 1
            }
            curr[j] = min(deletionCost, insertionCost, substitutionCost)
        }
        prev, curr = curr, prev
    }
    return prev[lenS]
}

func main() {
    scanner := makeScanner()
    s, t := readLine(scanner), readLine(scanner)
    fmt.Print(LevenshteinDistance(s, t))
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
