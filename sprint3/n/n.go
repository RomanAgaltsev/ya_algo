/*
N. Клумбы

Алла захотела, чтобы у неё под окном были узкие клумбы с тюльпанам.
На схеме земельного участка клумбы обозначаются просто горизонтальными отрезками, лежащими на одной прямой.
Для ландшафтных работ было нанято n садовников. Каждый из них обрабатывал какой-то отрезок на схеме.
Процесс был организован не очень хорошо, иногда один и тот же отрезок или его часть могли быть обработаны сразу несколькими садовниками.
Таким образом, отрезки, обрабатываемые двумя разными садовниками, сливаются в один. Непрерывный обработанный отрезок затем станет клумбой.
Нужно определить границы будущих клумб.

Рассмотрим примеры.
Пример 1:
Даны 4 отрезка:
[7,8],[7,8],[2,3],[6,10]. Два одинаковых отрезка [7,8] и [7,8] сливаются в один, но потом их накрывает отрезок [6,10].
Таким образом, имеем две клумбы с координатами [2,3] и [6,10].

Пример 2
Во втором сэмпле опять 4 отрезка: [2,3],[3,4],[3,4],[5,6]. Отрезки [2,3],[3,4] и [3,4] сольются в один отрезок [2,4].
Отрезок [5,6] ни с кем не объединяется, добавляем его в ответ.

Формат ввода
В первой строке задано количество садовников n. Число садовников не превосходит 100000.
В следующих n строках через пробел записаны координаты клумб в формате:
start end, где start — координата начала, end — координата конца. Оба числа целые, неотрицательные и не превосходят 107. start строго меньше, чем end.

Формат вывода
Нужно вывести координаты каждой из получившихся клумб в отдельных строках.
Данные должны выводиться в отсортированном порядке – сначала клумбы с меньшими координатами, затем – с бОльшими.

Пример 1
Ввод
4
7 8
7 8
2 3
6 10
Вывод
2 3
6 10

Пример 2
Ввод
4
2 3
5 6
3 4
3 4
Вывод
2 4
5 6

Пример 3
Ввод
6
1 3
3 5
4 6
5 6
2 4
7 10
Вывод
1 6
7 10
*/
package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := makeScanner()
	gardnersNumber := readInt(scanner)
	sections := readSections(scanner, gardnersNumber)
	slices.SortFunc(sections, func(a, b []int) int {
		if a[0] == b[0] && a[1] == b[1] {
			return 0
		}
		if a[0] < b[0] || (a[0] == b[0] && a[1] < b[1]) {
			return -1
		}
		return 1
	})
	flowerbeds := make([][]int, 0)
	for i := 0; i < gardnersNumber-1; i++ {
		if sections[i][1] >= sections[i+1][0] {
			sections[i+1][0] = min(sections[i][0], sections[i+1][0])
			sections[i+1][1] = max(sections[i][1], sections[i+1][1])
		} else {
			flowerbeds = append(flowerbeds, sections[i])
		}
	}
	flowerbeds = append(flowerbeds, sections[gardnersNumber-1])
	printArray(flowerbeds)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readSections(scanner *bufio.Scanner, n int) [][]int {
	sections := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		secBeg, _ := strconv.Atoi(arr[0])
		secEnd, _ := strconv.Atoi(arr[1])
		sections[i] = []int{secBeg, secEnd}
	}
	return sections
}

func printArray(arr [][]int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i][0]))
		writer.WriteString(" ")
		writer.WriteString(strconv.Itoa(arr[i][1]))
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
