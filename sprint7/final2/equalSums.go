/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 7
Задача - B. Одинаковые суммы

Отчеты:
- Ревью 1 - https://contest.yandex.ru/contest/25597/run-report/115959800/
- Ревью 2 - https://contest.yandex.ru/contest/25597/run-report/115972377/

-- ПРИНЦИП РАБОТЫ --
Основная идея решения заключается в том, чтобы проверить - можно ли из некоторого множества элементов слайса размера n получить половину суммы элементов этого слайса.

По задаче необходимо проверить, можно ли разбить считанный слайс на две части, чтобы суммы элементов в них были одинаковыми.
При таком требовании получается, что сумма элементов одной части должна равняться половине суммы всех элементов и сумма элементов второй части также должна равняться половине суммы всех элементов.

Получается, чтобы разбиение в принципе было возможным, сумма всех элементов слайса должна быть четной.
Нечетную сумму всех элементов слайса нельзя разделить на два без остатка.
Поэтому, еще до всех рассчетов имеет смысл проверить четность суммы всех элементов и сразу вернуть False, если сумма нечетная.

Далее, используя динамику, для возможных количеств элементов слайса определяем, можно ли из каждого количества собрать половину суммы всех элементов слайса или нельзя.

В данном решении используется всего один слайс partition размера в половину суммы (+1).
Так как обход слайса выполняется справа (с конца), то слева от указателя будут расположены значения предыдущей строки динамики, а справа - текущей строки.
При определении, можно ли собрать половину суммы используются данные предыдущей итерации цикла, то есть данные предыдущей строки динамики.

Например, для значений:
n = 4
Слайс = 1 5 7 1

- Сумма всех элементов слайса равна 14.
- Половина суммы всех элементов равна 7.
- Будем проверять, можно ли собрать половину суммы для одного элемента, для двух, для трех. Если три элемента наберут половину суммы, то и оставшийся будет ей равен.

В итоге, прицнип работы решения следующий:
1. Рассчитываем сумму всех элементов слайса;
2. Проверяем полученную сумму на нечетность. Если сумма нечетная, прерываем выполнение;
3. Вычисляем половину суммы всех элементов слайса;
4. Инициируем слайс partition размера половины суммы (+1);
5. Перебираем количества элементов слайса и слайс partition (справа) для определения возможности разбивки;
6. После обхода искомое решение находится в partition[halfSum]

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Если сумма некоторого множества элементов слайса равна половине суммы всех элементов слайса,
то и сумма оставшихся элементов слайса, не входящих в это множество, будет равна половине суммы всех элементов слайса.
Значит, произвести разбиение слайса на две части с равными суммами элементов возможно.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
В решении присутствуют два цикла - один по количествам элементов до n, второй до половины суммы всех элементов массива halfSum.

Получается оценка O(n * halfSum).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Дополнительная память для работы решения требуется для хранения слайса partition размера halfSum+1.

Получается оценка O(halfSum).

*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// canBePartitioned - проверяет, возможно ли разбить переданный слайс на две части,
//  чтобы сумма значений в них была одинаковой:
// - возвращает "True", если разбиение возможно;
// - возвращает "False", если разбить нельзя.
func canBePartitioned(arr []int, n int) string {
    // Сначала расситываем сумму элементов слайса
    sum := 0
    for i := 0; i < n; i++ {
        sum += arr[i]
    }
    // Проверяем четность суммы элементов - нечетную сумму разбить на две равные части нельзя
    if sum%2 != 0 {
        // Сумма нечетная, возвращаем "False"
        return "False"
    }
    // Сумма элементов четная, идем дальше
    // Полуаем половину суммы элементов слайса
    // Именно половину суммы должно составлять некоторое множество элементов, чтобы множество других элементов также составляло половину
    halfSum := sum / 2
    // Инициируем слайс для текущей строки динамики с количеством элементов, равным половине суммы +1
    partition := make([]bool, halfSum+1)
    // Чтобы не проверять в циклах, устанавливаем нулевое значение в true - 0 элементов всегда можно разбить
    partition[0] = true
    // Так как вложенный цикл будет обходить слайс partition справа (с конца),
    // Второй слайс для хранения предыдущей строки динамики не требуется
    // В слайсе слева от указателя j будут располагаться данные предыдущей строки динамики,
    // А по и справа от указателя j будут располагаться данные текущей строки динамики
    //
    // Внешним циклом перебираем количества элементов слайса, 
    // а внутренним циклом определяем - можно ли из такого количества элементов собрать половину суммы или нельзя
    for i := 1; i < n; i++ {
        for j := halfSum; j > arr[i]-1; j-- {
            // Определяем, можно ли собрать половину суммы или нелья - соответственно, можно ли разбить слайс или нельзя
            if partition[j-arr[i]] == true {
                // Текущую сумму j с использованием текущего количества элементов i можно разбить,
                // если текущая сумма за вычетом текущего элемента слайса может быть разбита
                partition[j] = true
            }
        }
    }
    // В итоге, проверяем, что получилось в слайсе partition для половины суммы исходного слайса
    if partition[halfSum] == true {
        // Половину суммы элементов слайса можно собрать некоторым множеством элементов слайса
        // Значит, разбить слайс на два подмножества с одинаковыми суммами элементов можно
        return "True"
    }
    // В противном случае, нельзя
    return "False"
}

func main() {
    // Создаем новый сканер
    scanner := makeScanner()
    // Считываем количество элементов проверяемого слайса
    n := readInt(scanner)
    // Считываем сам слайс
    arr := readArray(scanner)
    // Проверяем, можно ли разбить и сразу выводим результат
    fmt.Print(canBePartitioned(arr, n))
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

func readArray(scanner *bufio.Scanner) []int {
    scanner.Scan()
    listString := strings.Split(scanner.Text(), " ")
    arr := make([]int, len(listString))
    for i := 0; i < len(listString); i++ {
        arr[i], _ = strconv.Atoi(listString[i])
    }
    return arr
}
