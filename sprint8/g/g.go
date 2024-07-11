/*
G. Поиск со сдвигом

Гоша измерял температуру воздуха n дней подряд. В результате у него получился некоторый временной ряд.
Теперь он хочет посмотреть, как часто встречается некоторый шаблон в получившейся последовательности.
Однако температура — вещь относительная, поэтому Гоша решил, что при поиске шаблона длины m (a1, a2, ..., am) стоит также рассматривать сдвинутые на константу вхождения.
Это значит, что если для некоторого числа c в исходной последовательности нашёлся участок вида (a1 + c, a2 + c, ... , am + c),
то он тоже считается вхождением шаблона (a1, a2, ..., am).

По заданной последовательности измерений X и шаблону A=(a1, a2, ..., am) определите все вхождения A в X, допускающие сдвиг на константу.
Подсказка: если вы пишете на питоне и сталкиваетесь с TL, то попробуйте заменить какие-то из циклов операциями со срезами.

Формат ввода
В первой строке дано количество сделанных измерений n — натуральное число, не превышающее 104.
Во второй строке через пробел записаны n целых чисел xi, 0 ≤ xi ≤ 103 –— результаты измерений.
В третьей строке дано натуральное число m –— длина искомого шаблона, 1≤ m ≤ n.
В четвёртой строке даны m целых чисел ai — элементы шаблона, 0 ≤ ai ≤ 103.

Формат вывода
Выведите через пробел в порядке возрастания все позиции, на которых начинаются вхождения шаблона A в последовательность X.
Нумерация позиций начинается с единицы.

Пример 1
Ввод
9
3 9 1 2 5 10 9 1 7
2
4 10
Вывод
1 8

Пример 2
Ввод
5
1 2 3 4 5
3
10 11 12
Вывод
1 2 3

*/

package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func find(temp, patt []int, start int) int {
    if len(temp) < len(patt) {
        return -1
    }
    for pos := start; pos <= len(temp)-len(patt); pos++ {
        match := true
        shift := patt[0] - temp[pos]
        for offset := 0; offset < len(patt); offset++ {
            if (shift == 0 && temp[pos+offset] != patt[offset]) || (shift != 0 && temp[pos+offset]+shift != patt[offset]) {
                match = false
                break
            }
        }
        if match {
            return pos
        }
    }
    return -1
}

func findAll(temp, patt []int) []int {
    occur := make([]int, 0)
    start := 0
    for {
        pos := find(temp, patt, start)
        if pos == -1 {
            break
        }
        start = pos + 1
        occur = append(occur, start)
    }
    return occur
}

func main() {
    scanner := makeScanner()
    _ = readInt(scanner)
    temp := readArray(scanner)
    _ = readInt(scanner)
    patt := readArray(scanner)
    printArray(findAll(temp, patt))
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

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    stringInt := scanner.Text()
    res, _ := strconv.Atoi(stringInt)
    return res
}

func printArray(arr []int) {
    writer := bufio.NewWriter(os.Stdout)
    for i := 0; i < len(arr); i++ {
        writer.WriteString(strconv.Itoa(arr[i]))
        writer.WriteString(" ")
    }
    writer.Flush()
}
