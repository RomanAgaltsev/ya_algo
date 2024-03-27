/*
E. Самое длинное слово

Чтобы подготовиться к семинару, Гоше надо прочитать статью по эффективному менеджменту.
Так как Гоша хочет спланировать день заранее, ему необходимо оценить сложность статьи.

Он придумал такой метод оценки: берётся случайное предложение из текста и в нём ищется самое длинное слово.
Его длина и будет условной сложностью статьи.

Помогите Гоше справиться с этой задачей.

Формат ввода
В первой строке дана длина текста L (1 ≤ L ≤ 105).

В следующей строке записан текст, состоящий из строчных латинских букв и пробелов.
Слово — последовательность букв, не разделённых пробелами.
Пробелы могут стоять в самом начале строки и в самом её конце.
Текст заканчивается переносом строки, этот символ не включается в число остальных L символов.

Формат вывода
В первой строке выведите самое длинное слово. Во второй строке выведите его длину.
Если подходящих слов несколько, выведите то, которое встречается раньше.

Пример 1
Ввод
19
i love segment tree
Вывод
segment
7

Пример 2
Ввод
21
frog jumps from river
Вывод
jumps
5
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getLongestWord(line string) string {
	longestWord, currentWord := "", ""
	for _, letter := range line {
		if string(letter) == " " {
			if len(currentWord) > len(longestWord) {
				longestWord = currentWord
			}
			currentWord = ""
		} else {
			currentWord += string(letter)
		}
	}
	if len(currentWord) > len(longestWord) {
		longestWord = currentWord
	}
	return longestWord
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	line := readLine(scanner)
	longestWord := getLongestWord(line)
	fmt.Println(longestWord)
	fmt.Println(len(longestWord))
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}