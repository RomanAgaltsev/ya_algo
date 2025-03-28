/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 8
Задача - A. Packed Prefix

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/26133/run-report/116132293/
- Ревью 2 - https://contest.yandex.ru/contest/26133/run-report/116193610/

-- ПРИНЦИП РАБОТЫ --
Задачу можно разделить на две подзадачи:
1. Распаковка считываемых с ввода строк. Для решения используется стек слайсов байт для хранения фрагментов распаковываемых строк;
2. Определение наибольшего общего префикса распакованных строк. Для решения используется простой посимвольный перебор строк.

Подзадача распаковки считываемой строки решается следуюшим образом:
1. Создается новый стек для хранения фрагментов распаковываемой строки в виде слайсов байт;
2. Распаковываемая строка обходится в цикле посимвольно:
 - Если текущий символ равен открывающей скобке, ничего не делается. Символ просто пропускается;
 - Если текущий символ является целым числом, на стек добавляется новый фрагмент с этим числом;
 - Если текущий символ является буквой, он добавляется в фрагмент, который находится на вершине стека;
 - Если текущий символ равен закрывающей скобке, выполняется несколько действий:
    * Берется фрагмент строки с вершины стека;
    * Если первый символ фрагмента является числом, оставшиеся символы фрагмента (кроме первого) повторяются количество раз, равное этому числу;
    * Берется фрагмент строки с вершины стека - предыдущий фрагмент;
    * К этому предыдущему фрагменту добавляется справа текущий фрагмент;
    * Получившийся новый фрагмент строки помещается на вершину стека.
3. Все фрагменты, хранящиеся в стеке, объединяются в одну строку - это и есть распакованная строка.

Стек реализован при помощи типа Stack - слайс слайсов байт - с методами:
- Push - добавляет слайс байт (фрагмент строки) на вершину стека;
- Pop - извлекает с вершины стека слайс байт (фрагмент строки);
- Size - возвращает размер стека;
- String - формирует строку из фрагментов, хранящихся в стеке.

Подзадача поиска наибольшего общего префикса распакованных строк решается простым перебором символов сравниваемых строк:
1. До цикла обхода всех строк считывается и распаковывается первая строка - она берется за образец;
2. Далее в цикле считываются и распаковываются остальные строки;
3. Каждая считанная и распакованная строка сравнивается посимвольно с предыдущей;
4. Сравнение выполняется до окончания одной из строк или до первой неравной пары символов.

При решении подзадачи в памяти хранятся не все строки с ввода, а только две - предыдущая и текущая.
Результат всех сравнений является наибольшим общим префиксом.
Можно сказать, что наибольший общий префикс - это максимальное пересечение слева всех строк, и оно не может быть длиннее самой короткой строки.

Что можно поменять/доработать в решении:
1. По условию задачи, в запакованной строке "n[A]", n — однозначное натуральное число.
На этом строится решение в части определения множителя фрагаметов строки.
Если множитель будет не однозначным, а многозначным, такое решение не подойдет - надо будет поменять логику определения множителя;
2. Для множителей можно было бы использовать второй стек.
В этом случае немного упростилась бы обработка очередного фрагмента строки с умножением на множитель.
3. Подозреваю, что для поиска наибольшего общего префикса можно было бы использовать бинарный поиск.
Это немного усложнило бы реализацию, но сократило бы время и улучшило бы асимптотику.
Текущее решение прошло в Контесте по лимитам, поэтому не стал так делать.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Функция распаковки строк реализована в соответствии с постановкой задачи.
- Если запакованная строка D имеет вид D=AB, где A и B тоже запакованные строки, то результатом распаковки строки D является конкатенация строк A и B;
- Если запакованная строка D имеет вид D=n[A], где A - запакованная строка и n - целое однозначное число, то результатом распаковки строки D является повторенная n раз строка A.

При определении наибольшего общего префикса всех считанных и распакованных строк определяется наибольшее пересечение слева всех строк.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Примем, что:
- M - длина самой длинной считываемой строки в распакованном виде;
- n - количество строк.

Функция распаковки работает за O(M) для каждой строки - обрабатывается каждый символ строки, распаковка выполняется в цикле считывания строк;
Поиск общего префикса выполняется в том же цикле считвания строки и для пары строк работает за O(M).

*НЕКОРРЕКТНО!* Получается, что общая временная сложность - O(n * M * M) или O(n * M^2)
Получается, что общая временная сложность - O(n * M + n * M) или O(2 * n * M) или O(n * M).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Примем, что:
- M - длина самой длинной считываемой строки в распакованном виде;
- n - количество строк.

При распаковке строки для стека требуется O(M) дополнительной памяти.
При определении наибольшего общего префикса одновременно обрабатываются две строки - это O(2 * M).

Получается, решение требует O(M) дополнительной памяти.

*/

package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const (
    // Константы для скобок
    openBracket  = byte('[') // Открывающая скобка в виде байта
    closeBracket = byte(']') // Закрывающая скобка в виде байта
)

// Stack - стек для хранения фрагментов распаковываемой строки
type Stack [][]byte

// Push - добавляет переданный слайс байт в стек
func (s *Stack) Push(b []byte) {
    *s = append(*s, b)
}

// Pop - извлекает из стека и возвращает слайс байт
func (s *Stack) Pop() []byte {
    if s.Size() == 0 {
        return nil
    }
    lastIndex := len(*s) - 1
    lastBytes := (*s)[lastIndex]
    *s = (*s)[:lastIndex]
    return lastBytes
}

// Size - возвращает размер стека
func (s *Stack) Size() int {
    return len(*s)
}

// String - формирует строку из всех фрагментов стека и возвращает её
func (s *Stack) String() string {
    var builder strings.Builder
    for i := 0; i < s.Size(); i++ {
        builder.WriteString(string((*s)[i]))
    }
    return builder.String()
}

// unpack - распаковывает и возвращает переданную запакованную строку
func unpack(packedLine string) string {
    // Создаем стек для хранения фрагментов распаковываемой строки
    result := Stack{}
    // Обходим строку в цикле посимвольно и распаковываем её содержимое
    for i := 0; i < len(packedLine); i++ {
        // Если текущий символ равен открывающей скобке, ничего не делаем
        if packedLine[i] == openBracket {
            // Просто пропускаем символ
            continue
        }
        // Если текущий символ равен закрывающей скобке, распаковываем фрагмент строки с вершины стека
        if packedLine[i] == closeBracket {
            // Берем фрагмент с вершины стека
            fragment := result.Pop()
            // Пытаемся получить множитель из первого символа фрагмента - по условию множители однозначные
            multiplier, err := strconv.Atoi(string(fragment[0]))
            // Проверяем, случилась ли ошибка при получении множителя
            if err == nil {
                // Ошибки не случилось, умножаем фрагмент без первого символа (без множителя) на множитель
                fragment = bytes.Repeat(fragment[1:], multiplier)
            }
            // Опять берем фрагмент с вершины стека - предыдущий фрагмент относительно текущего
            prevFragment := result.Pop()
            // Приклеиваем текущий распакованный фрагмент к предыдущему справа
            prevFragment = append(prevFragment, fragment...)
            // Кладем на стек получившийся новый фрагмент
            result.Push(prevFragment)
            // Больше ничего не делаем
            continue
        }
        // Если текущий символ - это целое число, добавляем новый фрагмент на стек
        _, err := strconv.Atoi(string(packedLine[i]))
        if err == nil {
            // Ошибки не было, это целое число - в текущий фрагмент на стеке добавлять нельзя, надо добавлять новый
            result.Push([]byte{packedLine[i]})
            // Больше ничего не делаем
            continue
        }
        // Если текущий символ - это буква, просто добавляем в фрагмент на вершине стека
        // Берем фрагмент с вершины стека
        fragment := result.Pop()
        // Добавляем текущий символ-букву справа
        fragment = append(fragment, packedLine[i])
        // Возвращаем фрагмент обратно на стек
        result.Push(fragment)
    }
    // В итоге возвращаем распакованную строку
    return result.String()
}

func main() {
    // Создаем новый сканер
    scanner := makeScanner()
    // Считываем количество строк
    n := readInt(scanner)
    // Считываем и распаковываем первую строку до цикла, чтобы потом в цикле определить общий префикс со следующей
    prefix := unpack(readLine(scanner))
    // Считываем и распаковываем остальные строки, определяя общий префикс
    for i := 1; i < n; i++ {
        // Считываем и распаковываем очередную строку
        nextLine := unpack(readLine(scanner))
        // Проверяем длину очередной строки - общий префикс не может быть длинее самой короткой из всех строк
        if len(prefix) > len(nextLine) {
            // Очередная строка короче предыдущей, обрезаем текущую до длины очередной
            prefix = prefix[:len(nextLine)]
        }
        // Выполняем посимвольное сравнение строк для определения общего префикса
        for j := 0; j < len(nextLine); j++ {
            // Как только добрались до длины предыдущей строки или нашли неравные буквы
            if j == len(prefix) || nextLine[j] != prefix[j] {
                // Берем этот префикс
                prefix = prefix[:j]
                // С очередной строкой больше ничего не делаем
                break
            }
        }
    }
    // Выводим получившийся результат - наибольший общий префикс
    fmt.Print(prefix)
}

func makeScanner() *bufio.Scanner {
    const maxCapacity = 3 * 1024 * 1024
    buf := make([]byte, maxCapacity)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(buf, maxCapacity)
    return scanner
}

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    stringInt := scanner.Text()
    res, _ := strconv.Atoi(stringInt)
    return res
}

func readLine(scanner *bufio.Scanner) string {
    scanner.Scan()
    return scanner.Text()
}
