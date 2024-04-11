/*
L. Два велосипеда

Вася решил накопить денег на два одинаковых велосипеда — себе и сестре.
У Васи есть копилка, в которую каждый день он может добавлять деньги (если, конечно, у него есть такая финансовая возможность).
В процессе накопления Вася не вынимает деньги из копилки.

У вас есть информация о росте Васиных накоплений — сколько у Васи в копилке было денег в каждый из дней.

Ваша задача — по заданной стоимости велосипеда определить
- первый день, в которой Вася смог бы купить один велосипед,
- и первый день, в который Вася смог бы купить два велосипеда.
Подсказка: решение должно работать за O(log n).

Формат ввода
В первой строке дано число дней n, по которым велись наблюдения за Васиными накоплениями. 1 ≤ n ≤ 106.

В следующей строке записаны n целых неотрицательных чисел. Числа идут в порядке неубывания.
Каждое из чисел не превосходит 106.

В третьей строке записано целое положительное число s — стоимость велосипеда. Это число не превосходит 106.

Формат вывода
Нужно вывести два числа — номера дней по условию задачи.

Если необходимой суммы в копилке не нашлось, нужно вернуть -1 вместо номера дня.

Пример 1
Ввод
6
1 2 4 4 6 8
3
Вывод
3 5

Пример 2
Ввод
6
1 2 4 4 4 4
3
Вывод
3 -1

Пример 3
Ввод
6
1 2 4 4 4 4
10
Вывод
-1 -1
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func daySearch(arr []int, x, left, right int) int {
	if right <= left+1 {
		if arr[left] >= x {
			return left
		}
		if arr[right] >= x {
			return right
		}
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] >= x {
		return daySearch(arr, x, left, mid)
	} else {
		return daySearch(arr, x, mid, right)
	}
}

func getDays(savings []int, price int) []int {
	day1 := daySearch(savings, price, 0, len(savings)-1)
	day2 := daySearch(savings, 2*price, day1+1, len(savings)-1)
	if day1 > -1 {
		day1 += 1
	}
	if day2 > -1 {
		day2 += 1
	}
	return []int{day1, day2}
}

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	savings := readArray(scanner)
	price := readInt(scanner)
	printArray(getDays(savings, price))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 10 * 1024 * 1024
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
