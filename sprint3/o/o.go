/*
O. Разность треш-индексов

Гоша долго путешествовал и измерил площадь каждого из n островов Алгосов, но ему этого мало!
Теперь он захотел оценить, насколько разнообразными являются острова в составе архипелага.
Для этого Гоша рассмотрел все пары островов (таких пар, напомним, n * (n-1) / 2) и посчитал попарно разницу площадей между всеми островами.
Теперь он собирается упорядочить полученные разницы, чтобы взять k-ую по порядку из них.

Помоги Гоше найти k-ю минимальную разницу между площадями эффективно.

Пояснения к примерам
Пример 1
Выпишем все пары площадей и найдём соответствующие разницы
|2 - 3| = 1
|3 - 4| = 1
|2 - 4| = 2
Так как нам нужна 2-я по величине разница, то ответ будет 1.

Пример 2
У нас есть два одинаковых элемента в массиве — две единицы, поэтому минимальная (первая) разница равна нулю.

Формат ввода
В первой строке записано натуральное число n — количество островов в архипелаге (2 ≤ n ≤ 100 000).
В следующей строке через пробел записаны n площадей островов — n натуральных чисел, каждое из которых не превосходит 1 000 000.
В последней строке задано число k. Оно находится в диапазоне от 1 до n(n - 1) / 2.

Формат вывода
Выведите одно число — k-ую минимальную разницу.

Пример 1
Ввод
3
2 3 4
2
Вывод
1

Пример 2
Ввод
3
1 3 1
1
Вывод
0

Пример 3
Ввод
3
1 3 5
3
Вывод
4
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

func countDiffs(areas []int, m int) int {
	count, l, r, p := 0, 0, 0, 0
	for r < len(areas) {
		for areas[r] - areas[l] > m && l < r {
			l += 1
		}
		for r < len(areas) && areas[r] - areas[l] <= m {
			r += 1
		}
		count += (r - l) * (r - l - 1) / 2
		if p - l > 1 {
			count -= (p - l) * (p - l - 1) / 2
		}
		p = r
	}
	return count
}

func getKDiff(areas []int, k int) int {
	slices.Sort(areas)
	l, m, r := 0, 0, areas[len(areas)-1]
	for r > l {
		m = (r + l) / 2
		if countDiffs(areas, m) < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	areas := readArray(scanner)
	k := readInt(scanner)
	fmt.Print(getKDiff(areas, k))
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