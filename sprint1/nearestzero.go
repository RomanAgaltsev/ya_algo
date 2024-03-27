/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Задача - A. Ближайший ноль
Отчет - https://contest.yandex.ru/contest/22450/run-report/110431660/
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)
// GetNearestZeros - возвращает слайс с ближайшими нулями
// Название функции с большой буквы, потому что IDE ругалась на маленькую в файле тестов
func GetNearestZeros(houseNumbers *[]int) []int {
	// Размер входного слайса сохраним в переменной, чтобы не получать несколько раз
	houseNumbersLen := len(*houseNumbers)
	// Для экономии памяти используем тот же слайс
	// Новая переменная - для наглядности
	nearestZeros := *houseNumbers
	// Указатели для шагания по слайсу:
	// - prev - указывает на предыдущий 0, если он есть
	// - next - указывает на следующий 0, если он есть
	// Оба начинаются с (-1) - указывают на позицию до начала слайса
	prev, next := -1,-1
	// Если на улице первый же участок пустой, шагнем обоими указателями
	// Менять значение в слайсе не требуется - оно и так равно 0
	if nearestZeros[0] == 0 {
		prev, next = 0,0
	}
	// Флаг, определяющий, что next дошел до конца слайса
	// Нужен, чтобы отличить ситуации:
	// - next дошел до очередного 0 (пустого участка)
	// - next дошел до конца слайса
	nextIsStoped := false
	// next начинает шагать - шагает до конца слайса
	for next < houseNumbersLen - 1 {
		// next делает 1 шаг
		next += 1
		// Проверяем, не дошел ли next до конца слайса
		if next == houseNumbersLen - 1 {
			// next дошел до конца слайса
			nextIsStoped = true
		}
		// Проверяем, не нашел ли next 0 или не дошел ли до конца слайса
		if nearestZeros[next] == 0 || nextIsStoped {
			// next нашел 0 или дошел до конца слайса
			// prev начинает шагать - со следующей позиции и до next
			for i := prev + 1; i <= next; i++ {
				// prev делает 1 шаг
				// Проверяем, не пустой ли учаток
				if nearestZeros[i] == 0 {
					// Пустой, расстояние равно 0
					nearestZeros[i] = 0
					// Дальше выполнять код не надо
					continue
				}
				// Проверяем текущее состояние указателей
				if prev == -1 {
					// prev ни разу не шагнул
					// Расстояние берем от текущего участка до next
					nearestZeros[i] = next - i
				} else if nearestZeros[next] == 0 {
					// prev шагнул хотя бы 1 раз, next указывает на 0
					// Расстояние берем минимальное из "от prev до текущего" и "от текущего до next"
					prevDist := i - prev
					nextDist := next - i
					if prevDist < nextDist {
						nearestZeros[i] = prevDist
					} else {
						nearestZeros[i] = nextDist
					}
				} else {
					// next дошел до конца слайса и указывает на "не 0"
					nearestZeros[i] = i - prev
				}
			}
			// На самом деле, в предыдущем цикле шагал не prev, а i
			// Поэтому, надо передвинуть prev на текущий next
			prev = next
		}
	}
	return nearestZeros
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	houseNumbers := readArray(scanner)
	printArray(GetNearestZeros(&houseNumbers))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 10 * 1024 * 1024
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}