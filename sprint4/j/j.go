/*
J. Общий подмассив

Гоша увлёкся хоккеем и часто смотрит трансляции матчей.
Чтобы более-менее разумно оценивать силы команд, он сравнивает очки, набранные во всех матчах каждой командой.
Гоша попросил вас написать программу, которая по результатам игр двух выбранных команд найдёт наибольший по длине отрезок матчей,
когда эти команды зарабатывали одинаковые очки.

Рассмотрим первый пример:
Результаты первой команды: [1 2 3 2 1].
Результаты второй команды: [3 2 1 5 6].
Наиболее продолжительный общий отрезок этих массивов имеет длину 3 — это [3 2 1].

Формат ввода
В первой строке находится число n (1 ≤ n ≤ 105) –— количество матчей, которые были сыграны первой командой.
Во второй строке записано n целых чисел — очки в этих играх.
В третьей строке дано число m (1 ≤ m ≤ 105) – количество матчей, которые сыграла вторая команда.
В четвертой строке заданы m целых чисел — результаты второй команды.

Число очков, заработанных в одной игре, лежит в диапазоне от 0 до 255.

Формат вывода
Выведите целое неотрицательное число — максимальное количество матчей подряд, в которых команды зарабатывали одинаковые очки.

Пример 1
Ввод
5
1 2 3 2 1
5
3 2 1 5 6
Вывод
3

Пример 2
Ввод
5
1 2 3 4 5
3
4 5 9
Вывод
2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getHash(s string, a, m int) int {
	hash := 0
	for i := 0; i < len(s); i++ {
		c, _ := strconv.Atoi(string(s[i]))
		hash = (hash*a%m + c) % m
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

func checkCollision(nArr, mArr []string, window int) bool {
	const a = 997
	const m = 1000000007
	hashes := make(map[int]string)
	s := strings.Join(nArr[:window], "")
	hash := getHash(s, a, m)
	hashes[hash] = s
	power := getPower(window, a, m)
	for i := 1; i+window-1 < len(nArr); i++ {
		c1, _ := strconv.Atoi(nArr[i-1])
		c2, _ := strconv.Atoi(nArr[i+window-1])
		hash = (hash + m - c1*power%m) % m
		hash = (hash*a%m + c2) % m
		hashes[hash] = strings.Join(nArr[i:i+window], "")
	}
	s = strings.Join(mArr[:window], "")
	hash = getHash(s, a, m)
	if str, ok := hashes[hash]; ok && s == str {
		return true
	}
	for i := 1; i+window-1 < len(mArr); i++ {
		c1, _ := strconv.Atoi(nArr[i-1])
		c2, _ := strconv.Atoi(nArr[i+window-1])
		hash = (hash + m - c1*power%m) % m
		hash = (hash*a%m + c2) % m
		s = strings.Join(mArr[i:i+window], "")
		if str, ok := hashes[hash]; ok && s == str {
			return true
		}
	}
	return false
}

func getGLC(nArr, mArr []string) int {
	maxLen := 0
	left := 0
	right := min(len(nArr), len(mArr))
	for left <= right {
		middle := (left + right) / 2
		if checkCollision(nArr, mArr, middle) {
			maxLen = middle
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return maxLen
}

func main() {
	scanner := makeScanner()
	_, nArr := readInt(scanner), readArray(scanner)
	_, mArr := readInt(scanner), readArray(scanner)
	fmt.Print(getGLC(nArr, mArr))
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}