/*
K. Сравнить две строки

Алла придумала новый способ сравнивать две строки:
чтобы сравнить строки a и b, в них надо оставить только те буквы, которые в английском алфавите стоят на четных позициях.
Затем полученные строки сравниваются по обычным правилам.
Помогите Алле реализовать новое сравнение строк.

Формат ввода
На вход подаются строки a и b по одной в строке.
Обе строки состоят из маленьких латинских букв, не бывают пустыми и не превосходят 105 символов в длину.

Формат вывода
Выведите -1, если a < b, 0, если a = b, и 1, если a > b.

Пример 1
Ввод
gggggbbb
bbef
Вывод
-1

Пример 2
Ввод
z
aaaaaaa
Вывод
1

Пример 3
Ввод
ccccz
aaaaaz
Вывод
0
*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func compare(s, t string) int {
    sb := make([]byte, 0)
    tb := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        if s[i]%2 == 0 {
            sb = append(sb, s[i])
        }
    }
    for i := 0; i < len(t); i++ {
        if t[i]%2 == 0 {
            tb = append(tb, t[i])
        }
    }
    s = string(sb)
    t = string(tb)
    if s < t {
        return -1
    }
    if s > t {
        return 1
    }
    return 0
}

func main() {
    scanner := makeScanner()
    s, t := readLine(scanner), readLine(scanner)
    fmt.Print(compare(s, t))
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
