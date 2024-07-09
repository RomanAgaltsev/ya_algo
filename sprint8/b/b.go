/*
B. Пограничный контроль

Представьте, что вы работаете пограничником и постоянно проверяете документы людей по записи из базы.
При этом допустима ситуация, когда имя человека в базе отличается от имени в паспорте на одну замену, одно удаление или одну вставку символа.
Если один вариант имени может быть получен из другого удалением одного символа, то человека пропустят через границу.
А вот если есть какое-либо второе изменение, то человек грустно поедет домой или в посольство.
Например, если первый вариант —– это «Лена», а второй — «Лера», то девушку пропустят.
Также человека пропустят, если в базе записано «Коля», а в паспорте — «оля».
Однако вариант, когда в базе числится «Иннокентий», а в паспорте написано «ннакентий», уже не сработает.
Не пропустят также человека, у которого в паспорте записан «Иинннокентий», а вот «Инннокентий» спокойно пересечёт границу.
Напишите программу, которая сравнивает имя в базе с именем в паспорте и решает, пропускать человека или нет.
В случае равенства двух строк — путешественника, естественно, пропускают.

Формат ввода
В первой строке дано имя из паспорта.

Во второй строке — имя из базы.

Обе строки состоят из строчных букв английского алфавита. Размер каждой строки не превосходит 100 000 символов.

Формат вывода
Выведите «OK», если человека пропустят, или «FAIL» в противном случае.

Пример 1
Ввод
abcdefg
abdefg
Вывод
OK

Пример 2
Ввод
helo
hello
Вывод
OK

Пример 3
Ввод
dog
fog
Вывод
OK

Пример 4
Ввод
mama
papa
Вывод
FAIL

*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func canPass(passport, base string) string {
    if len(passport)-len(base) > 1 || len(passport)-len(base) < -1 {
        return "FAIL"
    }
    var deletion, insertion, substitution int
    switch {
    case len(passport) > len(base):
        deletion++
    case len(passport) < len(base):
        insertion++
    default:
        substitution++
    }
    for i := 0; i < len(passport); i++ {
        if passport[i] == base[i] {
            continue
        }
        if deletion > 0 {
            passport = passport[:i] + passport[i+1:]
            deletion--
            continue
        }
        if insertion > 0 {
            passport = passport[:i] + "#" + passport[i:]
            insertion--
            continue
        }
        if substitution > 0 {
            substitution--
            continue
        }
        return "FAIL"
    }
    return "OK"
}

func main() {
    scanner := makeScanner()
    passport, base := readLine(scanner), readLine(scanner)
    fmt.Print(canPass(passport, base))
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
