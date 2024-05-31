/*
C. Странное сравнение

Жители Алгосского архипелага придумали новый способ сравнения строк.
Две строки считаются равными, если символы одной из них можно заменить на символы другой так,
что первая строка станет точной копией второй строки. При этом необходимо соблюдение двух условий:
- Порядок вхождения символов должен быть сохранён.
- Одинаковым символам первой строки должны соответствовать одинаковые символы второй строки. Разным символам — разные.
Например, если строка s = «abacaba», то ей будет равна строка t = «xhxixhx»,
так как все вхождения «a» заменены на «x», «b» — на «h», а «c» — на «i».
Если же первая строка s=«abc», а вторая t=«aaa», то строки уже не будут равны,
так как разные буквы первой строки соответствуют одинаковым буквам второй.

Формат ввода
В первой строке записана строка s, во второй –— строка t.
Длины обеих строк не превосходят 106.
Обе строки содержат хотя бы по одному символу и состоят только из маленьких латинских букв.

Строки могут быть разной длины.

Формат вывода
Выведите «YES», если строки равны (согласно вышеописанным правилам), и «NO» в ином случае.

Пример 1
Ввод
mxyskaoghi
qodfrgmslc
Вывод
YES

Пример 2
Ввод
agg
xdd
Вывод
YES

Пример 3
Ввод
agg
xda
Вывод
NO
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := makeScanner()
	s, t := readLine(scanner), readLine(scanner)
	if len(s) != len(t) {
		fmt.Print("NO")
		return
	}
	r := "YES"
	m := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		c, ok := m[s[i]]
		if !ok {
			m[s[i]] = t[i]
			continue
		}
		if c != t[i] {
			r = "NO"
			break
		}
	}
	fmt.Print(r)
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
