/*
D. Хаотичность погоды

Метеорологическая служба вашего города решила исследовать погоду новым способом.

Под температурой воздуха в конкретный день будем понимать максимальную температуру в этот день.
Под хаотичностью погоды за n дней служба понимает количество дней, в которые температура строго больше,
чем в день до (если такой существует) и в день после текущего (если такой существует).
Например, если за 5 дней максимальная температура воздуха составляла [1, 2, 5, 4, 8] градусов,
то хаотичность за этот период равна 2: в 3-й и 5-й дни выполнялись описанные условия.
Определите по ежедневным показаниям температуры хаотичность погоды за этот период.

Заметим, что если число показаний n=1, то единственный день будет хаотичным.

Формат ввода
В первой строке дано число n –— длина периода измерений в днях, 1 ≤ n≤ 105. Во второй строке даны n целых чисел –— значения температуры в каждый из n дней. Значения температуры не превосходят 273 по модулю.

Формат вывода
Выведите единственное число — хаотичность за данный период.

Пример 1
Ввод
7
-1 -10 -8 0 2 0 5
Вывод
3

Пример 2
Ввод
5
1 2 5 4 8
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

func getWeatherRandomness(temperatures []int) int {
	if len(temperatures) == 1 {
		return 1
	}
	randomness := 0
	for currI, currTemp := range temperatures {
		prevI, nextI := currI - 1, currI + 1
		isPrevLess, isNextLess := false, false
		if prevI < 0 || temperatures[prevI] < currTemp {
			isPrevLess = true
		}
		if nextI > len(temperatures) - 1 || temperatures[nextI] < currTemp {
			isNextLess = true
		}
		if isPrevLess && isNextLess {
			randomness += 1
		}
	}
	return randomness
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	temperatures := readArray(scanner)
	fmt.Println(getWeatherRandomness(temperatures))
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
	for i := 0; i< len(listString); i++ {
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