/*
F. Палиндром

Помогите Васе понять, будет ли фраза палиндромом‎.
Учитываются только буквы и цифры, заглавные и строчные буквы считаются одинаковыми.

Решение должно работать за O(N), где N — длина строки на входе.

Формат ввода
В единственной строке записана фраза или слово.
Буквы могут быть только латинские.
Длина текста не превосходит 20000 символов.

Фраза может состоять из строчных и прописных латинских букв, цифр, знаков препинания.

Формат вывода
Выведите «True», если фраза является палиндромом, и «False», если не является.

Пример 1
Ввод
A man, a plan, a canal: Panama
Вывод
True

Пример 2
Ввод
zo
Вывод
False
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isPalindrome(line string) bool {
	left, right := 0, len(line)-1
	for left < right {
		for left < right {
			leftRune := rune(line[left])
			if unicode.IsDigit(leftRune) || unicode.IsLetter(leftRune) {
				break
			}
			left += 1
		}
		for left < right {
			rightRune := rune(line[right])
			if unicode.IsDigit(rightRune) || unicode.IsLetter(rightRune) {
				break
			}
			right -= 1
		}
		if !strings.EqualFold(string(line[left]), string(line[right])) {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	scanner := makeScanner()
	line := readLine(scanner)
	if isPalindrome(line) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
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

