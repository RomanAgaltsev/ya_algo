/*
K. Списочная форма

Вася просил Аллу помочь решить задачу. На этот раз по информатике.
Для неотрицательного целого числа X списочная форма –— это массив его цифр слева направо.
К примеру, для 1231 списочная форма будет [1,2,3,1].
На вход подается количество цифр числа Х, списочная форма неотрицательного числа Х и неотрицательное число K.
Число К не превосходят 10000. Длина числа Х не превосходит 1000.

Нужно вернуть списочную форму числа X + K.
Не используйте встроенные средства языка для сложения длинных чисел.

Формат ввода
В первой строке — длина списочной формы числа X. На следующей строке — сама списочная форма с цифрами записанными через пробел.
В последней строке записано число K, 0 ≤ K ≤ 10000.

Формат вывода
Выведите списочную форму числа X+K.

Пример 1
Ввод
4
1 2 0 0
34
Вывод
1 2 3 4

Пример 2
Ввод
2
9 5
Вывод
17

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetSum(bigNumber []int, smallNumber int) []int {
	smallNumberSlice := make([]int, 0)
	smallString := strconv.Itoa(smallNumber)
	
	for i := 0; i < len(smallString); i++ {
		dig, _ := strconv.Atoi(string(smallString[i]))
		smallNumberSlice = append(smallNumberSlice, dig)
	}
	
	if len(smallNumberSlice) < len(bigNumber) {
		for i := len(smallNumberSlice); i < len(bigNumber); i++ {
			smallNumberSlice = append([]int{0}, smallNumberSlice...)
		}
	}
	if len(bigNumber) < len(smallNumberSlice) {
		for i := len(bigNumber); i < len(smallNumberSlice); i++ {
			bigNumber = append([]int{0}, bigNumber...)
		}
	}
	
	carry, sum := 0, 0
	for i := len(bigNumber) - 1; i >= 0; i-- {
		sum = carry + bigNumber[i] + smallNumberSlice[i]
		bigNumber[i] = sum % 10
		carry = sum / 10
	}

	if carry != 0 {
		bigNumber = append([]int{carry}, bigNumber...)
	}

	return bigNumber
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	bigNumer := readArray(scanner)
	smallNumber := readInt(scanner)
	printArray(GetSum(bigNumer, smallNumber))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
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
