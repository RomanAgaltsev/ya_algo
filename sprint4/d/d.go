/*
D. Полиномиальный хеш

Алле очень понравился алгоритм вычисления полиномиального хеша.
Помогите ей написать функцию, вычисляющую хеш строки s.
В данной задаче необходимо использовать в качестве значений отдельных символов их коды в таблице ASCII.

Полиномиальный хеш считается по формуле:
h(s)=(s1*qn−1 + s2*qn−2 + ⋯ +sn−1*q + sn) mod R

Формат ввода
В первой строке дано число a (1 ≤ a ≤ 1000) — основание, по которому считается хеш.
Во второй строке дано число m (1 ≤ m ≤ 109) — модуль.
В третьей строке дана строка s (0 ≤ |s| ≤ 106), состоящая из больших и маленьких латинских букв.

Формат вывода
Выведите целое неотрицательное число — хеш заданной строки.

Пример 1
Ввод
123
100003
a
Вывод
97

Пример 2
Ввод
123
100003
hash
Вывод
6080

Пример 3
Ввод
123
100003
HaSH
Вывод
56156
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getPolyHash(s string, a, m int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return int(s[0]) % m
	}
	sum := ((int(s[0]) * a) + int(s[1])) % m
	for i := 2; i < len(s); i++ {
		sum = ((sum * a) + int(s[i])) % m
	}
	return sum
}

func main() {
	scanner := makeScanner()
	a, m := readInt(scanner), readInt(scanner)
	s := readLine(scanner)
	fmt.Print(getPolyHash(s, a, m))
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