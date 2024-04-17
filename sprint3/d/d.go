/*
D. Печеньки

К Васе в гости пришли одноклассники. Его мама решила угостить ребят печеньем.
Но не всё так просто. Печенья могут быть разного размера.
А у каждого ребёнка есть фактор жадности —– минимальный размер печенья, которое он возьмёт.
Нужно выяснить, сколько ребят останутся довольными в лучшем случае, когда они действуют оптимально.

Каждый ребёнок может взять не больше одного печенья.

Формат ввода
В первой строке записано n — количество детей.

Во второй — n чисел, разделённых пробелом, каждое из которых — фактор жадности ребёнка.
Это натуральные числа, не превосходящие 1000.

В следующей строке записано число m –— количество печенек.

Далее — m натуральных чисел, разделённых пробелом — размеры печенек. Размеры печенек не превосходят 1000.

Оба числа n и m не превосходят 10000.

Формат вывода
Нужно вывести одно число — количество детей, которые останутся довольными

Пример 1
Ввод
2
1 2
3
2 1 3
Вывод
2

Пример 2
Ввод
3
2 1 3
2
1 1
Вывод
1
*/

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
	"strconv"
)

func FeedKids(kids []int, cookies []int) int {
	feeded := 0
	slices.SortFunc(kids, func(a, b int) int {
		if n := cmp.Compare(b, a); n != 0 {
			return n
		}
		return 0
	})
	slices.SortFunc(cookies, func(a, b int) int {
		if n := cmp.Compare(b, a); n != 0 {
			return n
		}
		return 0
	})
	for _, kid := range kids {
		if len(cookies) == 0 {
			break
		}
		if kid <= cookies[0] {
			cookies = cookies[1:]
			feeded++
		}
	}
	return feeded
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	kids := readArray(scanner)
	readInt(scanner)
	cookies := readArray(scanner)
	fmt.Print(FeedKids(kids, cookies))
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