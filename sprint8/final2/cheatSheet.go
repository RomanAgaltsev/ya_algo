/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 8
Задача - B. Шпаргалка

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/26133/run-report/116162498/

-- ПРИНЦИП РАБОТЫ --


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --


*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

const (
    // Константы для ответов
    yes = "YES" // Ответ, когда строку разбить возможно
    no  = "NO"  // Ответ, когда строку разбить нельзя
)

// Node - структура узла префиксного дерева
type Node struct {
    children   map[byte]*Node // Мапа для хранения переходов
    isTerminal bool           // Признак терминального узла
}

// insert - добавляет переданное слово в префиксное дерево
func (n *Node) insert(word string) {
    // Начинаем с корневого узла
    node := n
    // Обходим переданное слово посимвольно
    for i := 0; i < len(word); i++ {
        // Пытаемся получить переходы для текущего символа
        next, ok := node.children[word[i]]
        // Проверяем, есть ли переходы для текущего символа в принципе
        if !ok {
            // Переходов нет, добавляем новый узел для текущего символа
            next = newNode()
            // Для текущего узла добавляем переход в следующий по текущему символу
            node.children[word[i]] = next
        }
        // Делаем следующий узел текущим - переходим в него
        node = next
    }
    // Слово закончилось, для последнего узла устанавливаем признак терминального
    node.isTerminal = true
}

// newNode - создает новый узел префиксного дерева и возвращает ссылку на него
func newNode() *Node {
    // Создаем новый узел
    node := &Node{
        children:   make(map[byte]*Node), // Инициируем мапу переходов
        isTerminal: false,                // Признак терминального узла по умолчанию = false
    }
    // Возвращаем узел
    return node
}

// getTrie - создает новое префиксное дерево и возвращает ссылку на его корневой узел
func getTrie(scanner *bufio.Scanner) *Node {
    // Считываем количество допустимых к использованию слов
    n := readInt(scanner)
    // Создаем новый узел - корень префиксного дерева
    root := newNode()
    // Считываем слова
    for i := 0; i < n; i++ {
        // Каждое считанное слово добавляем в префиксное дерево
        root.insert(readLine(scanner))
    }
    // Возвращаем ссылку на корневой узел
    return root
}

// isLineSplitable - проверяет, можно ли разбить переданную строку на слова из словаря префиксного дерева:
// - false - разбить строку нельзя
// - true - разить строку возможно
func isLineSplitable(line string, root *Node) bool {
    // Сохраняем длину строки в переменную
    lenLine := len(line)
    // Создаем слайс для динамики по длине строки+1
    dp := make([]bool, lenLine+1)
    // Считаем, что строку из нуля символов можно разбить
    dp[0] = true
    // Обходим строку посимвольно для заполнения слайса динамики
    for i := 0; i <= lenLine; i++ {
        // Для каждого символа поиск будем выполнять, начиная с корневого узла
        node := root
        // Проверяем, можно ли разбить строку для предыдущей позиции
        // i-ый символ строки соответствует (i-1)-ому значению динамики
        if dp[i] {
            // Значение динамики true - имеет смысл идти дальше
            // Обходим строку посимвольно с текущей позиции и до конца строки
            for j := i; j <= lenLine; j++ {
                // Проверяем, текущий узел префиксного дерева не терминальный ли
                if node.isTerminal {
                    // Текущий узел терминальный - символами строки мы прошли в дереве слово полностью
                    // Значит, до текущей позиции разбить можно, записываем true в слайс динамики
                    dp[j] = true
                }
                // Если дошли до конца строки или для текущего символа нет переходов в префиксном дереве
                if j == lenLine || node.children[line[j]] == nil {
                    // Прерываем вложенный цикл
                    break
                }
                // Переходим из узла текущего символа в следующий
                node = node.children[line[j]]
            }
        }
    }
    // Искомый ответ хранится в последнем элементе слайса динамики
    return dp[lenLine]
}

func main() {
    // Создаем новый сканер
    scanner := makeScanner()
    // Проверяем, можно ли разбить считанную строку
    if isLineSplitable(readLine(scanner), getTrie(scanner)) {
        // Строку разбить можно - выводим "YES"
        fmt.Print(yes)
        return
    }
    // Строку разить нельзя - выводим "NO"
    fmt.Print(no)
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
