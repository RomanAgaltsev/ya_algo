/*
I. Анаграммная группировка

Вася решил избавиться от проблем с произношением и стать певцом. Он обратился за помощью к логопеду.
Тот посоветовал Васе выполнять упражнение, которое называется анаграммная группировка. В качестве подготовительного этапа нужно выбрать из множества строк анаграммы.

Анаграммы — это строки, которые получаются друг из друга перестановкой символов. Например, строки «SILENT» и «LISTEN» являются анаграммами.

Помогите Васе найти анаграммы.

Формат ввода
В первой строке записано число n — количество строк.
Далее в строку через пробел записаны n строк.
n не превосходит 6000. Длина каждой строки не более 100 символов.

Формат вывода
Нужно вывести в отсортированном порядке индексы строк, которые являются анаграммами.
Каждая группа индексов должна быть выведена в отдельной строке. Индексы внутри одной группы должны быть отсортированы по возрастанию.
Группы между собой должны быть отсортированы по возрастанию первого индекса.

Обратите внимание, что группа анаграмм может состоять и из одной строки.
Например, если в исходном наборе нет анаграмм, то надо вывести n групп, каждая из которых состоит из одного индекса.

Пример
Ввод
6
tan eat tea ate nat bat
Вывод
0 4
1 2 3
5
*/

package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

var letters = map[byte]int{
    byte('a'): 2,
    byte('b'): 3,
    byte('c'): 5,
    byte('d'): 7,
    byte('e'): 11,
    byte('f'): 13,
    byte('g'): 17,
    byte('h'): 19,
    byte('i'): 23,
    byte('j'): 29,
    byte('k'): 31,
    byte('l'): 37,
    byte('m'): 41,
    byte('n'): 43,
    byte('o'): 47,
    byte('p'): 53,
    byte('q'): 59,
    byte('r'): 61,
    byte('s'): 67,
    byte('t'): 71,
    byte('u'): 73,
    byte('v'): 79,
    byte('w'): 83,
    byte('x'): 89,
    byte('y'): 91,
    byte('z'): 101,
}

func getHash(word string) int {
    var hash int
    var bytes int
    for i := 0; i < len(word); i++ {
        hash += (int(word[i]) * letters[word[i]] * letters[word[i]]) % 1000000001
        bytes += (letters[word[i]] * len(word)) % 1000000001
    }
    return (hash * bytes * len(word)) % 1000000001
}

func getAnagramPositions(words []string) {
    data := make(map[int][]int)
    order := make([]int, 0)
    for i := 0; i < len(words); i++ {
        hash := getHash(words[i])
        if positions, ok := data[hash]; !ok {
            order = append(order, hash)
            data[hash] = []int{i}
        } else {
            data[hash] = append(positions, i)
        }
    }
    for i := 0; i < len(order); i++ {
        printArray(data[order[i]])
    }
}

func main() {
    scanner := makeScanner()
    _ = readInt(scanner)
    words := readArray(scanner)
    getAnagramPositions(words)
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

func readArray(scanner *bufio.Scanner) []string {
    scanner.Scan()
    return strings.Split(scanner.Text(), " ")
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