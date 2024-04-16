/*
C. Подпоследовательность

Гоша любит играть в игру «Подпоследовательность»:
даны 2 строки, и нужно понять, является ли первая из них подпоследовательностью второй.
Когда строки достаточно длинные, очень трудно получить ответ на этот вопрос, просто посмотрев на них.
Помогите Гоше написать функцию, которая решает эту задачу.

Формат ввода
В первой строке записана строка s.
Во второй — строка t.

Обе строки состоят из маленьких латинских букв, длины строк не превосходят 150000. Строки не могут быть пустыми.

Формат вывода
Выведите True, если s является подпоследовательностью t, иначе —– False.

Пример 1
Ввод
abc
ahbgdcu
Вывод
True

Пример 2
Ввод
abcp
ahpc
Вывод
False
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := makeScanner()
	s := readLine(scanner)
	t := readLine(scanner)
	for i := 0; i < len(t); i++ {
		if len(s) == 0 {
			break
		}
		if s[0] == t[i] {
			s = s[1:]
		}
	}
	if len(s) == 0 {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
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
