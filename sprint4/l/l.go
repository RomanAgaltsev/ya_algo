/*
L. МногоГоша

Дана длинная строка, состоящая из маленьких латинских букв.
Нужно найти все её подстроки длины n, которые встречаются хотя бы k раз.

Формат ввода
В первой строчке через пробел записаны два натуральных числа n и k.
Во второй строчке записана строка, состоящая из маленьких латинских букв. Длина строки 1 ≤ L ≤ 106.

n ≤ L, k ≤ L.

Формат вывода
Для каждой найденной подстроки выведите индекс начала её первого вхождения (нумерация в строке начинается с нуля).

Выводите индексы в любом порядке, в одну строку, через пробел.

Пример 1
Ввод
10 2
gggggooooogggggoooooogggggssshaa
Вывод
0 5

Пример 2
Ввод
3 4
allallallallalla
Вывод
0 1 2

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getHash(s string, a, m int) int {
	hash := 0
	for i := 0; i < len(s); i++ {
		hash = (hash*a%m + int(s[i])) % m
	}
	return hash
}

func getPower(n int, a int, m int) int {
	power := 1
	for i := 1; i < n; i++ {
		power = (power * a) % m
	}
	return power
}

func getPositions(s string, n, k, a, m int) []int {
	counter := make(map[int]int)
	hashToPos := make(map[int]int)
	hashes := make(map[int]string)
	hash := getHash(s[:n], a, m)
	hashes[hash] = s[:n]
	power := getPower(n, a, m)
	counter[hash]++
	hashToPos[hash] = 0
	result := make([]int, 0)
	for i := 1; i+n-1 < len(s); i++ {
		hash = (hash + m - int(s[i-1])*power%m) % m
		hash = (hash*a%m + int(s[i+n-1])) % m
		if counter[hash] == 0 {
			counter[hash]++
			hashToPos[hash] = i
			hashes[hash] = s[i : i+n]
		} else if hashes[hash] == s[i:i+n] {
			counter[hash]++
		}
		if counter[hash] == k {
			result = append(result, hashToPos[hash])
		}
	}
	return result
}

func main() {
	scanner := makeScanner()
	n, k := readNK(scanner)
	s := readLine(scanner)
	if n <= len(s) {
		printArray(getPositions(s, n, k, 997, 1000000007))
	}
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

func readNK(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr[0], arr[1]
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