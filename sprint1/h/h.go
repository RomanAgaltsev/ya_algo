/*
H. Двоичная система

Тимофей записал два числа в двоичной системе счисления и попросил Гошу вывести их сумму, также в двоичной системе.
Встроенную в язык программирования возможность сложения двоичных чисел применять нельзя.
Помогите Гоше решить задачу.

Решение должно работать за O(N), где N –— количество разрядов максимального числа на входе.

Формат ввода
Два числа в двоичной системе счисления, каждое на отдельной строке.
Длина каждого числа не превосходит 10 000 символов.

Формат вывода
Одно число в двоичной системе счисления.

Пример 1
Ввод
1010
1011
Вывод
10101

Пример 2
Ввод
1
1
Вывод
10

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetSum(firstNumber string, secondNumber string) string {
	// Переменная для хранения результата - суммы
	res := ""
	
	// По условию длина чисел может быть разной - необходимо выравнять строки
	
	// В первую строку помещаем короткое число
	// Во вторую строку помещаем длинное число
	firstString, secondString := "", ""
	if len(firstNumber) < len(secondNumber) {
		firstString, secondString = firstNumber, secondNumber
	} else {
		firstString, secondString = secondNumber, firstNumber
	}
	
	// Цикл по разнице длинн - дополняем слева короткую строку
	for i := len(firstString); i < len(secondString); i++ {
		firstString = "0"+firstString
	}
	
	// Остаток для переноса на следующий разряд слева
	carry := 0
	// Обходим вторую строку и собираем результат
	for i := len(secondString)-1 ; i >= 0; i-- {
		// Бит первой строки - unt8->string->int
		bit1, _ := strconv.Atoi(string(firstString[i]))
		// Бит второй строки - unt8->string->int
		bit2, _ := strconv.Atoi(string(secondString[i]))
		// XOR-им биты и перенос с предыдущего разряда
		// 0 ^ 0 ^ 0 = 0
		// 0 ^ 1 ^ 0 = 1
		// 0 ^ 1 ^ 1 = 0
		// 1 ^ 0 ^ 0 = 1
		// 1 ^ 1 ^ 0 = 0
		// 1 ^ 1 ^ 1 = 1
		sum := bit1 ^ bit2 ^ carry
		// Полученное число преобразуем в строку и добавляем в результат слева
		res = strconv.Itoa(sum) + res
		// Вычисляем остаток для переноса
		// 0 & 0 = 0, 0 & 1 = 0, 1 & 0 = 0, 1 & 1 = 1
		// 0 | 0 = 0, 0 | 1 = 1, 1 | 0 = 1, 1 | 1 = 1
		carry = (bit1&carry)|(bit2&carry)|(bit1&bit2)
	}
	// После цикла может остаться перенос на следующий разряд - добавим слева
	if carry == 1 {
		res = strconv.Itoa(1) + res
	}
	return res
}

func main() {
	scanner := makeScanner()
	firstNumber := readLine(scanner)
	secondNumber := readLine(scanner)
	sum := GetSum(firstNumber, secondNumber)
	fmt.Println(sum)
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

