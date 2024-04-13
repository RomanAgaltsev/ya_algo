/*
B. Комбинации

На клавиатуре старых мобильных телефонов каждой цифре соответствовало несколько букв. Примерно так:

2:'abc',
3:'def',
4:'ghi',
5:'jkl',
6:'mno',
7:'pqrs',
8:'tuv',
9:'wxyz'

Вам известно в каком порядке были нажаты кнопки телефона, без учета повторов.
Напечатайте все комбинации букв, которые можно набрать такой последовательностью нажатий.

Формат ввода
На вход подается строка, состоящая из цифр 2-9 включительно. Длина строки не превосходит 10 символов.

Формат вывода
Выведите все возможные комбинации букв через пробел в лексикографическом (алфавитном) порядке по возрастанию.

Пример 1
Ввод
23
Вывод
ad ae af bd be bf cd ce cf

Пример 2
Ввод
92
Вывод
wa wb wc xa xb xc ya yb yc za zb zc

*/

package main

import (
	"fmt"
	"strings"
)

var keys = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func readLine() []string {
	var aString string
	fmt.Scan(&aString)
	line := strings.Split(aString, "")
	return line
}

func main() {
	line := readLine()
	printCombination(line, 0, "")
}

func printCombination(line []string, n int, result string) {
	if n == len(line) {
		fmt.Print(result)
		fmt.Print(" ")
		return
	}
	letters := keys[line[n]]
	for _, letter := range letters {
		printCombination(line, n+1, result+letter)
	}
}
