/*
I. Любители конференций

На IT-конференции присутствовали студенты из разных вузов со всей страны.
Для каждого студента известен ID университета, в котором он учится.
Тимофей предложил Рите выяснить, из каких k вузов на конференцию пришло больше всего учащихся.

Формат ввода
В первой строке дано количество студентов в списке —– n (1 ≤ n ≤ 15 000).
Во второй строке через пробел записаны n целых чисел —– ID вуза каждого студента. Каждое из чисел находится в диапазоне от 0 до 10 000.
В третьей строке записано одно число k.

Формат вывода
Выведите через пробел k ID вузов с максимальным числом участников.
Они должны быть отсортированы по убыванию популярности (по количеству гостей от конкретного вуза).
Если более одного вуза имеет одно и то же количество учащихся, то выводить их ID нужно в порядке возрастания.

Пример 1
Ввод
7
1 2 3 1 2 3 4
3
Вывод
1 2 3

Пример 2
Ввод
6
1 1 1 2 2 3
1
Вывод
1
*/

package main

import (
	"bufio"
	"cmp"
	"os"
	"slices"
	"strings"
	"strconv"
)

func countGuests(guests []int, k int) []int {
	counts := make(map[int]int)
	for _, guest := range guests {
		counts[guest]++ 
	}
	sums := make([][]int, 0)
	for id, count := range counts {
		sums = append(sums, []int{id, count})
	}
	slices.SortFunc(sums, func(a, b []int) int {
		if n := cmp.Compare(b[1], a[1]); n != 0 {
			return n
		}
		if n := cmp.Compare(a[0], b[0]); n != 0 {
			return n
		}
		return 0
	})
	if k > len(sums) {
		k = len(sums)
	}
	res := make([]int, k)
	for i := 0; i <= k-1; i++ {
		res[i] = sums[i][0]
	}
	return res
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	guests := readArray(scanner)
	k := readInt(scanner)
	printArray(countGuests(guests, k))
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.WriteString("\n")
	writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}