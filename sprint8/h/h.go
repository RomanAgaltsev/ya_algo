/*
H. Глобальная замена

Напишите программу, которая будет заменять в тексте все вхождения строки s на строку t.
Гарантируется, что никакие два вхождения шаблона s не пересекаются друг с другом.

Формат ввода
В первой строке дан текст — это строка из строчных букв английского алфавита, длина которой не превышает 106.
Во второй строке записан шаблон s, вхождения которого будут заменены.
В третьей строке дана строка t, которая будет заменять вхождения.
Обе строки s и t состоят из строчных букв английского алфавита, длина каждой строки не превосходит 105.
Размер итоговой строки не превосходит 2⋅ 106.

Формат вывода
В единственной строке выведите результат всех замен — текст, в котором все вхождения s заменены на t.

Пример 1
Ввод
pingpong
ng
mpi
Вывод
pimpipompi

Пример 2
Ввод
aaa
a
ab
Вывод
ababab
*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func search(p, text string) []int {
    var result []int
    s := p + "#" + text
    pi := make([]int, len(p))
    var piPrev int
    for i := 1; i < len(s); i++ {
        k := piPrev
        for k > 0 && s[k] != s[i] {
            k = pi[k-1]
        }
        if s[k] == s[i] {
            k++
        }
        if i < len(p) {
            pi[i] = k
        }
        piPrev = k
        if k == len(p) {
            result = append(result, i-2*len(p))
        }
    }
    return result
}

func main() {
    scanner := makeScanner()
    text, s, t := readLine(scanner), readLine(scanner), readLine(scanner)
    occur := search(s, text)
    result := text
    for i := len(occur) - 1; i >= 0; i-- {
        result = result[:occur[i]] + t + result[occur[i]+len(s):]
    }
    fmt.Print(result)
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
