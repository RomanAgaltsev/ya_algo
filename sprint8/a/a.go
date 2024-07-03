/*
A. Разворот строки

В некоторых языках предложения пишутся и читаются не слева направо, а справа налево.
Вам под руку попался странный текст — в нём обычный (слева направо) порядок букв в словах.
А вот сами слова идут в противоположном направлении.
Вам надо преобразовать текст так, чтобы слова в нём были написаны слева направо.

Формат ввода
На ввод подаётся строка, состоящая из слов, разделённых пробелами (один пробел между соседними словами).
Всего слов не более 1000, длина каждого из них —– от 1 до 100 символов.
Слова состоят из строчных букв английского алфавита.

Формат вывода
Выведите строку с обратным порядком слов в ней.

Пример 1
Ввод
one two three
Вывод
three two one

Пример 2
Ввод
hello
Вывод
hello

Пример 3
Ввод
may the force be with you
Вывод
you with be force the may
*/

package main

import (
    "bufio"
    "os"
    "strings"
)

func main() {
    scanner := makeScanner()
    words := readArray(scanner)
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    printArray(words)
}

func makeScanner() *bufio.Scanner {
    const maxCapacity = 3 * 1024 * 1024
    buf := make([]byte, maxCapacity)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(buf, maxCapacity)
    return scanner
}

func readArray(scanner *bufio.Scanner) []string {
    scanner.Scan()
    return strings.Split(scanner.Text(), " ")
}

func printArray(arr []string) {
    writer := bufio.NewWriter(os.Stdout)
    for i := 0; i < len(arr); i++ {
        writer.WriteString(arr[i])
        writer.WriteString(" ")
    }
    writer.Flush()
}
