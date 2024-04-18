/*
E. Покупка домов

Тимофей решил купить несколько домов на знаменитом среди разработчиков Алгосском архипелаге.
Он нашёл n объявлений о продаже, где указана стоимость каждого дома в алгосских франках.
А у Тимофея есть k франков.
Помогите ему определить, какое наибольшее количество домов на Алгосах он сможет приобрести за эти деньги.

Формат ввода
В первой строке через пробел записаны натуральные числа n и k.
n — количество домов, которые рассматривает Тимофей, оно не превосходит 100000;
k — общий бюджет, не превосходит 100000;
В следующей строке через пробел записано n стоимостей домов.
Каждое из чисел не превосходит 100000. Все стоимости — натуральные числа.

Формат вывода
Выведите одно число — наибольшее количество домов, которое может купить Тимофей.

Пример 1
Ввод
3 300
999 999 999
Вывод
0

Пример 2
Ввод
3 1000
350 999 200
Вывод
2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"strconv"
)

func countHouses(budget int, prices []int) int {
	count := 0
	slices.Sort(prices)
	for _, price := range prices {
		if price <= budget {
			count++
			budget -= price
		}
	}
	return count
}

func main() {
	scanner := makeScanner()
	firstArray := readArray(scanner)
	secondArray := readArray(scanner)
	fmt.Print(countHouses(firstArray[1], secondArray))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}