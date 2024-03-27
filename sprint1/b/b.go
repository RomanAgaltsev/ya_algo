/*
B. Чётные и нечётные числа

Представьте себе онлайн-игру для поездки в метро: игрок нажимает на кнопку,
и на экране появляются три случайных числа.
Если все три числа оказываются одной чётности, игрок выигрывает.

Напишите программу, которая по трём числам определяет, выиграл игрок или нет.

Формат ввода
В первой строке записаны три случайных целых числа a, b и c. Числа не превосходят 109 по модулю.

Формат вывода
Выведите «WIN», если игрок выиграл, и «FAIL» в противном случае.

Пример 1
Ввод
1 2 -3
Вывод
FAIL

Пример 2
Ввод
7 11 7
Вывод
WIN

Пример 3
Ввод
6 -2 0
Вывод
WIN
*/

package main

import (
	"fmt"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkParity(a int, b int, c int) bool {
	remaindersSum := abs(a)%2 + abs(b)%2 + abs(c)%2
	return remaindersSum == 0 || remaindersSum == 3
}

func main() {
	a := readInt()
	b := readInt()
	c := readInt()
	if checkParity(a, b, c) {
		fmt.Println("WIN")
	} else {
		fmt.Println("FAIL")
	}
}

func readInt() int {
	var aString string
	fmt.Scan(&aString)
	a, _ := strconv.Atoi(aString)
	return a
}
