/*
J. Пузырёк

Чтобы выбрать самый лучший алгоритм для решения задачи, Гоша продолжил изучать разные сортировки.
На очереди сортировка пузырьком — https://ru.wikipedia.org/wiki/Сортировка_пузырьком

Её алгоритм следующий (сортируем по неубыванию):
1. На каждой итерации проходим по массиву, поочередно сравнивая пары соседних элементов.
Если элемент на позиции i больше элемента на позиции i + 1, меняем их местами.
После первой итерации самый большой элемент всплывёт в конце массива.
2. Проходим по массиву, выполняя указанные действия до тех пор, пока на очередной итерации не окажется,
что обмены больше не нужны, то есть массив уже отсортирован.
3. После не более чем n – 1 итераций выполнение алгоритма заканчивается,
так как на каждой итерации хотя бы один элемент оказывается на правильной позиции.

Помогите Гоше написать код алгоритма.

Формат ввода
В первой строке на вход подаётся натуральное число n — длина массива, 2 ≤ n ≤ 1000.
Во второй строке через пробел записано n целых чисел.
Каждое из чисел по модулю не превосходит 1000.

Обратите внимание, что считывать нужно только 2 строки: значение n и входной массив.

Формат вывода
После каждого прохода по массиву, на котором какие-то элементы меняются местами, выводите его промежуточное состояние.
Таким образом, если сортировка завершена за k меняющих массив итераций,
то надо вывести k строк по n чисел в каждой — элементы массива после каждой из итераций.
Если массив был изначально отсортирован, то просто выведите его.

Пример 1
Ввод
5
4 3 9 2 1
Вывод
3 4 2 1 9
3 2 1 4 9
2 1 3 4 9
1 2 3 4 9

Пример 2
Ввод
5
12 8 9 10 11
Вывод
8 9 10 11 12
*/

package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

func bubbleSort(arr []int) {
	isSorted := true
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
				isSorted = false
			}
		}
		if !swapped {
			break
		}
		printArray(arr)
	}
	if isSorted {
		printArray(arr)
	}
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	bubbleSort(arr)
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
