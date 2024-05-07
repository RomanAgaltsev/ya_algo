/*
F. Префиксные хеши

Алла не остановилась на достигнутом — теперь она хочет научиться быстро вычислять хеши произвольных подстрок данной строки. Помогите ей!
На вход поступают запросы на подсчёт хешей разных подстрок. Ответ на каждый запрос должен выполняться за O(1).
Допустимо в начале работы программы сделать предподсчёт для дальнейшей работы со строкой.

Напомним, что полиномиальный хеш считается по формуле
h(s)=(s1*qn−1 + s2*qn−2 + ⋯ +sn−1*q + sn) mod R
В данной задаче необходимо использовать в качестве значений отдельных символов их коды в таблице ASCII.

Формат ввода
В первой строке дано число a (1 ≤ a ≤ 1000) — основание, по которому считается хеш.
Во второй строке дано число m (1 ≤ m ≤ 107) — модуль.
В третьей строке дана строка s (0 ≤ |s| ≤ 106), состоящая из больших и маленьких латинских букв.

В четвертой строке дано число запросов t — натуральное число от 1 до 105.
В каждой из следующих t строк записаны через пробел два числа l и r — индексы начала и конца очередной подстроки.
(1 ≤ l ≤ r ≤ |s|).

Формат вывода
Для каждого запроса выведите на отдельной строке хеш заданной в запросе подстроки.

Пример 1
Ввод
1000
1000009
abcdefgh
7
1 1
1 5
2 3
3 4
4 4
1 8
5 8

Вывод
97
225076
98099
99100
100
436420
193195

Пример 2
Ввод
100
10
a
1
1 1

Вывод
7
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getPrefixHashes(s string, a, m int) []int {
	prefixHashes := make([]int, len(s)+1)
	for i := 1; i <= len(s); i++ {
		prefixHashes[i] = (prefixHashes[i-1] * a % m + int(s[i-1])) % m
	}
	return prefixHashes
}

func getPowers(s string, a, m int) []int {
	powers := make([]int, len(s)+1)
	powers[0] = 1
	for i := 1; i <= len(s); i++ {
		powers[i] = (powers[i-1] * a) % m
	}
	return powers
}

func getHash(prefixHashes, powers []int, l, r, m int) int {
	return (prefixHashes[r] + m - (prefixHashes[l] * powers[r - l]) % m) % m
}

func main() {
	scanner := makeScanner()
	
	a := readInt(scanner)
	m := readInt(scanner)
	s := readLine(scanner)
	t := readInt(scanner)
	
	prefixHashes := getPrefixHashes(s, a, m)
	powers := getPowers(s, a, m)
	
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < t; i++ {
		l, r := readInterval(scanner)
		writer.WriteString(strconv.Itoa(getHash(prefixHashes, powers, l-1, r, m)))
		writer.WriteString("\n")
	}
	writer.Flush()
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

func readInterval(scanner *bufio.Scanner) (int,int) {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr[0], arr[1]
}