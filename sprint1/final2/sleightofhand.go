/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Задача - B. Ловкость рук

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/22450/run-report/110483578/
- Ревью 2 -	https://contest.yandex.ru/contest/22450/run-report/110531558/

Алгоритмическая сложность:
Память - O(n)
Предполагаю, что за n принимаем количество ячеек на поле. Для 4х4 - 16 ячеек.
При увеличении поля будет увеличиваться память, требуемая для хранения поля.
Для решения используем:
- Слайс слайсов для поля - память для него зависит от количества ячеек. При увеличении поля память будет расти линейно;
- Несколько переменных  - k, points, count. Для них память всегда будет константой. При увеличении поля расти не будет;
- Мапа для результата сканирования поля - в худшем случае в мапе будет 9 элементов (цифры от 1 до 9 и их количества). Можно считать константой. При увеличении поля расти больше 9 не будет.
Таким образом, память, требуемая для решения задачи, будет линейно зависеть от размера поля (количества ячеек).

Вычислительная - O(n)
Опять же, за n принимаем количество ячеек.
В ходе решения:
- Чтение поля с ввода в слайс слайсов - n, чтение выполяется 1 раз. При увеличении поля будет расти линейно.
- Сканирование поля слайса слайсов - n, обход поля выполняется 1 раз. При увеличении поля будет расти линейно.
- Подсчет баллов по мапе - 9 (в худшем случае). При увеличении поля расти не будет.
Таким образом, можно сказать, что вычислительная сложность будет равна O(n + n + 9).
Но берем максимальную и отбрасываем коэффициенты и константы - O(n).

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getPoint - определяет для переданной комбинации k и количества цифр, заработан балл или нет
func getPoint(k int, count int) int {
	// Проверяем условие задачи
	if 2 * k >= count {
		// Условие выполнилось, возвращаем 1 балл
		return 1
	}
	// Условие не выполнилось, балл не возвращаем
	return 0
}

// scanField - сканирует переданное поле и возвращает мапу цифра:количество её появлений на поле
func scanField(field [][]int) map[int]int {
	// Мапа для хранения результата сканирования поля - цифра:количество её появлений на поле
	res := make(map[int]int)
	// Обходим поле в цикле
	for _, line := range field {
		for _, val := range line {
			// Проверяем цифру на 0
			if val != 0 {
				// В мапу складываем только ненулевые цифры
				res[val] += 1
			}
		}
	}
	return res
}

// GetPpoints - по переданным полю и k подсчитывает количество заработанных баллов
func GetPpoints(field [][]int, k int) int {
	// Счетчик заработанных баллов
	points := 0
	// Сканируем поле только 1 раз, результат - мапа
	scannedField := scanField(field)
	// Обходим в цикле уже не поле, а мапу
	// В цикле будет столько итераций, сколько различных цифр было найдено на поле - кроме 0
	for _, count := range scannedField {
		// Считаем баллы
		points += getPoint(k, count)
	}
	return points
}

func main() {
	scanner := makeScanner()
	k := readInt(scanner)
	field := readField(scanner)
	fmt.Print(GetPpoints(field, k))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readField(scanner *bufio.Scanner) [][]int {
	field := make([][]int, 0)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), "")
		arr := make([]int, len(listString))
		for j := 0; j < len(listString); j++ {
			arr[j], _ = strconv.Atoi(listString[j])
		}
		field = append(field, arr)
	}
	return field
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}