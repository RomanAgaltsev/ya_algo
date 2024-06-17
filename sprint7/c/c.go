/*
C. Золотая лихорадка

Гуляя по одному из островов Алгосского архипелага, Гоша набрёл на пещеру, в которой лежат кучи золотого песка.
К счастью, у Гоши есть с собой рюкзак грузоподъёмностью до M килограмм, поэтому он может унести с собой какое-то ограниченное количество золота.
Всего золотых куч n штук, и все они разные. В куче под номером i содержится mi килограммов золотого песка, а стоимость одного килограмма — ci алгосских франков.

Помогите Гоше наполнить рюкзак так, чтобы общая стоимость золотого песка в пересчёте на алгосские франки была максимальной.

Формат ввода
В первой строке задано целое число M — грузоподъёмность рюкзака Гоши (0 ≤ M ≤ 108).
Во второй строке дано количество куч с золотым песком — целое число n (1 ≤ n ≤ 105).
В каждой из следующих n строк описаны кучи: i-ая куча задаётся двумя целыми числами ci и mi, записанными через пробел (1 ≤ ci ≤ 107, 1 ≤ mi ≤ 108).

Формат вывода
Выведите единственное число —– максимальную сумму (в алгосских франках), которую Гоша сможет вынести из пещеры в своём рюкзаке.

Пример 1
Ввод
10
3
8 1
2 10
4 5

Вывод
36

Пример 2
Ввод
10000
1
4 20

Вывод
80
*/

package main

import (
    "bufio"
    "cmp"
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
)

type Pile struct {
    cost, weight int
}

func getMaxSum(piles []Pile, capacity int) int {
    slices.SortFunc(piles, func(a, b Pile) int {
        return cmp.Compare(b.cost, a.cost)
    })
    result := 0
    for _, pile := range piles {
        if capacity == 0 {
            break
        }
        maxWeight := min(capacity, pile.weight)
        result += maxWeight * pile.cost
        capacity -= maxWeight
    }
    return result
}

func main() {
    scanner := makeScanner()
    capacity, n := readInt(scanner), readInt(scanner)
    piles := readPiles(scanner, n)
    fmt.Print(getMaxSum(piles, capacity))
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

func readPiles(scanner *bufio.Scanner, n int) []Pile {
    piles := make([]Pile, n)
    for i := 0; i < n; i++ {
        piles[i] = readPile(scanner)
    }
    return piles
}

func readPile(scanner *bufio.Scanner) Pile {
    scanner.Scan()
    listPile := strings.Split(scanner.Text(), " ")
    cost, _ := strconv.Atoi(listPile[0])
    weight, _ := strconv.Atoi(listPile[1])
    return Pile{
        cost,
        weight,
    }
}