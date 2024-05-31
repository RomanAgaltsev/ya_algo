/*
M. Золотая середина

Задача повышенной сложности
На каждом острове в архипелаге Алгосы живёт какое-то количество людей или же остров необитаем (тогда на острове живёт 0 людей).
Пусть на i-м острове численность населения составляет ai. Тимофей захотел найти медиану среди всех значений численности населения.

Определение: Медиана (https://ru.wikipedia.org/wiki/Медиана_(статистика)) массива чисел a_i
— это такое число, что половина чисел из массива не больше него, а другая половина не меньше.
В общем случае медиану массива можно найти, отсортировав числа и взяв среднее из них.
Если количество чисел чётно, то возьмём в качестве медианы полусумму соседних средних чисел, (a[n/2] + a[n/2 + 1])/2.

У Тимофея уже есть отдельно данные по северной части архипелага и по южной, причём значения численности населения в каждой группе отсортированы по неубыванию.
Определите медианную численность населения по всем островам Алгосов.

Подсказка: Если n –— число островов в северной части архипелага, а m — в южной, то ваше решение должно работать за .

Формат ввода
В первой строке записано натуральное число n, во второй — натуральное число m. Они не превосходят 10 000.
Далее в строку через пробел записаны n целых неотрицательных чисел, каждое из которых не превосходит 10 000, — значения численности населения в северной части Алгосов.
В последней строке через пробел записаны m целых неотрицательных чисел, каждое из которых не превосходит 10 000 — значения численности населения в южной части Алгосов.

Значения в третьей и четвёртой строках упорядочены по неубыванию.

Формат вывода
Нужно вывести одной число — найденную медиану.

Пример 1
Ввод
2
1
1 3
2
Вывод
2

Пример 2
Ввод
2
2
1 2
3 4
Вывод
2.5

Пример 3
Ввод
8
10
0 0 0 1 3 3 5 10
4 4 5 7 7 7 8 9 9 10
Вывод
5
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func getKElement(a []int, n int, b []int, m int, k int) {
    if n > m {
        getKElement(b, m, a, n, k)
    }
    if n == 0 {
		fmt.Printf("%d", b[k-1])
    }
    if k == 1 {
		if (len(a)+len(b))%2 == 0 {
			fmt.Printf("%f", (a[len(a)-n] + b[len(b)-m])/2)
        } else {
			fmt.Printf("%d", min(a[len(a)-n], b[len(b)-m]))
        }
    }
    i := min(n, k/2)
    j := min(m, k/2)
    if a[i-1] > b[j-1] {
        getKElement(a, n, b, m-j, k-j)
    } else {
        getKElement(a, n-i, b, m, k-i)
    }
}

func main() {
	scanner := makeScanner()
	n, m := readInt(scanner), readInt(scanner)
	a, b := readArray(scanner), readArray(scanner)
	getKElement(a, n, b, m, (n+m+1)/2)
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}