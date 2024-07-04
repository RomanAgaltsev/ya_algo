/*
E. Вставка строк

У Риты была строка s, Гоша подарил ей на 8 марта ещё n других строк ti, 1≤ i≤ n. Теперь Рита думает, куда их лучше поставить.
Один из вариантов —– расположить подаренные строки внутри имеющейся строки s,
поставив строку ti сразу после символа строки s с номером ki (в частности, если ki=0, то строка вставляется в самое начало s).

Помогите Рите и определите, какая строка получится после вставки в s всех подаренных Гошей строк.

Формат ввода
В первой строке дана строка s. Строка состоит из строчных букв английского алфавита, не бывает пустой и её длина не превышает 105 символов.
Во второй строке записано количество подаренных строк — натуральное число n, 1 ≤ n ≤ 105.

В каждой из следующих n строк через пробел записаны пары ti и ki.
Строка ti состоит из маленьких латинских букв и не бывает пустой. ki — целое число, лежащее в диапазоне от 0 до |s|.
Все числа ki уникальны. Гарантируется, что суммарная длина всех строк ti не превосходит 105.

Формат вывода
Выведите получившуюся в результате вставок строку.

Пример 1
Ввод
abacaba
3
queue 2
deque 0
stack 7
Вывод
dequeabqueueacabastack

Пример 2
Ввод
kukareku
2
p 1
q 2
Вывод
kpuqkareku
*/

package main

import (
    "bufio"
    "cmp"
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
)

type Line struct {
    sub string
    pos int
}

func main() {
    scanner := makeScanner()
    s := readLine(scanner)
    n := readInt(scanner)
    lines := make([]Line, n)
    for i := 0; i < n; i++ {
        tk := strings.Split(readLine(scanner), " ")
        t := tk[0]
        k, _ := strconv.Atoi(tk[1])
        lines = append(lines, Line{t, k})
    }
    slices.SortFunc(lines, func(a, b Line) int {
        return cmp.Compare(a.pos, b.pos)
    })
    result := make([]byte, 0)
    pos := 0
    for i := 0; i < len(lines); i++ {
        result = append(result, s[pos:lines[i].pos]...)
        result = append(result, lines[i].sub...)
        pos = lines[i].pos
    }
    if pos < len(s) {
        result = append(result, s[pos:]...)
    }
    fmt.Print(string(result))
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

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    stringInt := scanner.Text()
    res, _ := strconv.Atoi(stringInt)
    return res
}
