/*
A. Кружки

В компании, где работает Тимофей, заботятся о досуге сотрудников и устраивают различные кружки по интересам.
Когда кто-то записывается на занятие, в лог вносится название кружка.
По записям в логе составьте список всех кружков, в которые ходит хотя бы один человек.

Формат ввода
В первой строке даётся натуральное число n, не превосходящее 10 000 — количество записей в логе.
В следующих n строках — названия кружков.

Формат вывода
Выведите уникальные названия кружков по одному на строке, в порядке появления во входных данных.

Пример
Ввод
8
вышивание крестиком
рисование мелками на парте
настольный керлинг
настольный керлинг
кухня африканского племени ужасмай
тяжелая атлетика
таракановедение
таракановедение

Вывод
вышивание крестиком
рисование мелками на парте
настольный керлинг
кухня африканского племени ужасмай
тяжелая атлетика
таракановедение
*/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	classes := make(map[string]bool)
	scanner := makeScanner()
	n := readInt(scanner)
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		class := readLine(scanner)
		if _, ok := classes[class]; !ok {
			classes[class] = true
			writer.WriteString(class)
			writer.WriteString("\n")
		}
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