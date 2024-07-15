/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 8
Задача - A. Packed Prefix

Отчеты:
- Ревью 1 -

-- ПРИНЦИП РАБОТЫ --


n — однозначное натуральное число
Второй стек для множителей
Бинарный поиск для нахождения наибольшего общего префикса

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --


*/

package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const (
    openBracket  = byte('[')
    closeBracket = byte(']')
)

type Stack [][]byte

func (s *Stack) Push(b []byte) {
    *s = append(*s, b)
}

func (s *Stack) Pop() []byte {
    if s.Size() == 0 {
        return nil
    }
    lastIndex := len(*s) - 1
    lastBytes := (*s)[lastIndex]
    *s = (*s)[:lastIndex]
    return lastBytes
}

func (s *Stack) Size() int {
    return len(*s)
}

func (s *Stack) String() string {
    var builder strings.Builder
    for i := 0; i < s.Size(); i++ {
        builder.WriteString(string((*s)[i]))
    }
    return builder.String()
}

func unpack(packedLine string) string {
    result := Stack{}
    for i := 0; i < len(packedLine); i++ {
        if packedLine[i] == openBracket {
            continue
        }
        if packedLine[i] == closeBracket {
            fragment := result.Pop()
            multiplier, err := strconv.Atoi(string(fragment[0]))
            if err == nil {
                fragment = bytes.Repeat(fragment[1:], multiplier)
            }
            prevFragment := result.Pop()
            prevFragment = append(prevFragment, fragment...)
            result.Push(prevFragment)
            continue
        }
        _, err := strconv.Atoi(string(packedLine[i]))
        if err == nil {
            result.Push([]byte{packedLine[i]})
            continue
        }
        fragment := result.Pop()
        fragment = append(fragment, packedLine[i])
        result.Push(fragment)
    }
    return result.String()
}

func main() {
    scanner := makeScanner()
    n := readInt(scanner)
    prefix := unpack(readLine(scanner))
    for i := 1; i < n; i++ {
        nextLine := unpack(readLine(scanner))
        if len(prefix) > len(nextLine) {
            prefix = prefix[:len(nextLine)]
        }
        for j := 0; j < len(nextLine); j++ {
            if j >= len(prefix) || nextLine[j] != prefix[j] {
                prefix = prefix[:j]
                break
            }
        }
    }
    fmt.Print(prefix)
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

func readLine(scanner *bufio.Scanner) string {
    scanner.Scan()
    return scanner.Text()
}
