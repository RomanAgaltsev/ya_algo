/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 8
Задача - A. Packed Prefix

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/26133/run-report/116132293/

-- ПРИНЦИП РАБОТЫ --


n — однозначное натуральное число
Второй стек для множителей
Бинарный поиск для нахождения наибольшего общего префикса

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --


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
            if j >= len(prefix) || nextLine[j] != prefix[j] {
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
