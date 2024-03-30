/*
L. Лишняя буква

Васе очень нравятся задачи про строки, поэтому он придумал свою.
Есть 2 строки s и t, состоящие только из строчных букв.
Строка t получена перемешиванием букв строки s и добавлением 1 буквы в случайную позицию.
Нужно найти добавленную букву.

Формат ввода
На вход подаются строки s и t, разделённые переносом строки.
Длины строк не превосходят 1000 символов. Строки не бывают пустыми.

Формат вывода
Выведите лишнюю букву.

Пример 1
Ввод
abcd
abcde
Вывод
e
Пример 2

Ввод
go
ogg
Вывод
g

Пример 3
Ввод
xtkpx
xkctpx
Вывод
c

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func getExcessiveLetter(s1 string, t1 string) string {
	countsS := make(map[string]int)
	countsT := make(map[string]int)
	
	for _, bt := range s1 {
		countsS[string(bt)] += 1
	}
	for _, bt := range t1 {
		countsT[string(bt)] += 1
	}
	
	for kT, vT := range countsT {
		if vS, ok := countsS[kT]; !ok || vT > vS {
			return kT
		}
	}
	return ""
}

func main() {
	scanner := makeScanner()
	s := readLine(scanner)
	t := readLine(scanner)
	fmt.Printf(getExcessiveLetter(s, t))
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