/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Задача - A. Ближайший ноль

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/22450/run-report/110431660/
- Ревью 2 - https://contest.yandex.ru/contest/22450/run-report/110567531/
- Ревью 3 - https://contest.yandex.ru/contest/22450/run-report/110624819/
+ Ревью 4 - https://contest.yandex.ru/contest/22450/run-report/110670435/
+ Ревью 5 - https://contest.yandex.ru/contest/22450/run-report/110705502/

Алгоритмическая сложность:
Память - O(n)
Предполагаю, что за n принимаем количество элементов в слайсе.
При увеличении слайса будет увеличиваться память, требуемая для его хранения.
Для решения используем:
- Слайс - память для него зависит от количества элементов. При увеличении входных данныъ память будет расти линейно;
- Несколько переменных - счетчики, расстояние. Для них память всегда будет константой. При увеличении поля расти не будет;
Таким образом, память, требуемая для решения задачи, будет линейно зависеть от размера входных данных.

Вычислительная - O(n)
Опять же, за n принимаем количество элементов в слайсе.
В ходе решения:
- Чтение ввода в слайс - n, чтение выполяется 1 раз. При увеличении ввода будет расти линейно.
- Обход слайса и вычисление ближайших нулей - 2 * n, так как обходим слайс 2 раза - слева-направо и справа-налево.
Таким образом, можно сказать, что вычислительная сложность будет равна O(3n).
Но берем максимальную и отбрасываем коэффициенты и константы - O(n).

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
	// Расстояние до 0 слева
	dist := 0
	// Если первый элемент не 0, то в первый проход слева-направо расстояние нельзя отсчитывать от 0
	if nearestZeros[0] != 0 {
		// Первый элемент не 0, берем длину слайса для начала
		dist = houseNumbersLen
	}
	// Обходим слайс слева-направо, записываем расстояния до 0 слева
	for i := 0; i <= houseNumbersLen-1; i++ {
		// Проверяем текущий элемент на равенство 0
		if nearestZeros[i] == 0 {
			// Если текущий элемент равен 0, расстояние обнуляем
			// Значение элемента в слайсе перезаписывать не надо - там и так 0
			dist = 0
		} else {
			// Если текущий элемент не равен 0, увеличиваем расстояние на 1
			dist += 1
			// И записываем расстояние до 0 слева в элемент слайса
			nearestZeros[i] = dist
		}
	}
	// В первый проход слева-направо в слайс записаны расстояния до 0 слева
	// Исключение - если в первом элементе слайса был не 0,
	//  то от первого элемента слайса до первого нуля слева значения в элементах слайса "кривые"
	
	// Второй обход справа-налево для корректировки записанных расстояний - уточняем расстояние до правых нулей
	// Обходим слайс в обратную сторону, начиная с позиции последнего правого 0:
	// - если последний правый 0 в последнем элементе слайса - dist = 0 с предыдущего обхода, этот обход начнется с последнего элемента
	// - если последний правый 0 не в последнем элементе слайса - dist != 0 с предыдущего обхода, этот обход начнется с последнего 0
	for i := houseNumbersLen-dist-1; i >= 0; i-- {
		// Проверяем текущий элемент на равенство 0
		if nearestZeros[i] == 0 {
			// Если текущий элемент равен 0, обнуляем расстояние
			dist = 0
			// Значение элемента в слайсе перезаписывать не надо - там и так 0
		} else {
			// Если текущий элемент не равен 0
			// Увеличиваем расстояние, пройденное влево от 0 справа
			dist += 1
			// Перезаписывать надо только те элементы, которые больше пройденного расстояния влево от текущего правого 0
			// Как только левый 0 станет ближе, элементы станут меньше расстояния - пропускаем, перемещаемся сразу на следующий 0 слева
			// Чтобы не вылететь за пределы слайса, проверяем, что позиция следующего 0 слева существует (положительная)
			if nextI := i - nearestZeros[i]; dist >= nearestZeros[i] && nextI >= 0 {
				i = nextI
				dist = 0
			}
			// Перезаписываем расстояние
			nearestZeros[i] = dist
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