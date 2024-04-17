/*
G. Гардероб

Рита решила оставить у себя одежду только трёх цветов: розового, жёлтого и малинового.
После того как вещи других расцветок были убраны, Рита захотела отсортировать свой новый гардероб по цветам.
Сначала должны идти вещи розового цвета, потом — жёлтого, и в конце — малинового. Помогите Рите справиться с этой задачей.

Примечание: попробуйте решить задачу за один проход по массиву!

Формат ввода
В первой строке задано количество предметов в гардеробе: n — оно не превосходит 1000000.
Во второй строке даётся массив, в котором указан цвет для каждого предмета.
Розовый цвет обозначен 0, жёлтый — 1, малиновый — 2.

Формат вывода
Нужно вывести в строку через пробел цвета предметов в правильном порядке.

Пример 1
Ввод
7
0 2 1 2 0 0 1
Вывод
0 0 0 1 1 2 2

Пример 2
Ввод
5
2 1 2 0 1
Вывод
0 1 1 2 2

Пример 3
Ввод
6
2 1 1 2 0 2
Вывод
0 1 1 2 2 2
*/

package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

func WardrobeSort(wardrobe []int) []int {
	l, r, i := 0, len(wardrobe)-1, 0
	for i <= r {
		if wardrobe[i] == 0 {
			wardrobe[l], wardrobe[i] = wardrobe[i], wardrobe[l]
			l++
			if wardrobe[i] != 2 {
				i++
			}
		} else if wardrobe[i] == 2 {
			wardrobe[r], wardrobe[i] = wardrobe[i], wardrobe[r]
			r--
			if wardrobe[i] == 1 {
				i++
			}
		} else {
			i++
		}
	}
	return wardrobe
}

func main() {
	scanner := makeScanner()
	wardrobeSize := readInt(scanner)
	if wardrobeSize > 0 {
		wardrobe := readArray(scanner)
		printArray(WardrobeSort(wardrobe))
	}
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