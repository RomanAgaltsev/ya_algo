/*
B. Расписание

Дано количество учебных занятий, проходящих в одной аудитории.
Для каждого из них указано время начала и конца. Нужно составить расписание, в соответствии с которым в классе можно будет провести как можно больше занятий.
Если возможно несколько оптимальных вариантов, то выведите любой. Возможно одновременное проведение более чем одного занятия нулевой длительности.

Формат ввода
В первой строке задано число занятий. Оно не превосходит 1000.
Далее для каждого занятия в отдельной строке записано время начала и конца, разделённые пробелом.
Время задаётся одним целым числом h, если урок начинается/заканчивается ровно в h часов.
Если же урок начинается/заканчивается в h часов m минут, то время записывается как h.m.
Гарантируется, что каждое занятие начинается не позже, чем заканчивается. Указываются только значащие цифры.

Формат вывода
Выведите в первой строке наибольшее число уроков, которое можно провести в аудитории.
Далее выведите время начала и конца каждого урока в отдельной строке в порядке их проведения.

Пример 1
Ввод
5
9 10
9.3 10.3
10 11
10.3 11.3
11 12

Вывод
3
9 10
10 11
11 12

Пример 2
Ввод
3
9 10
11 12.25
12.15 13.3

Вывод
2
9 10
11 12.25

Пример 3
Ввод
7
19 19
7 14
12 14
8 22
22 23
5 21
9 23

Вывод
3
7 14
19 19
22 23
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

type Lesson struct {
    begin, end float64
}

func getSchedule(lessons []Lesson) []Lesson {
    result := make([]Lesson, 0)
    slices.SortFunc(lessons, func(a, b Lesson) int {
        if a.end == b.end {
            return cmp.Compare(a.begin, b.begin)
        }
        return cmp.Compare(a.end, b.end)
    })
    prev := Lesson{0, 0}
    for _, lesson := range lessons {
        if lesson.begin >= prev.end {
            prev = lesson
            result = append(result, prev)
        }
    }
    return result
}

func main() {
    scanner := makeScanner()
    n := readInt(scanner)
    lessons := readLessons(scanner, n)
    schedule := getSchedule(lessons)
    fmt.Print(len(schedule), "\n")
    printLessons(schedule)
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

func readLessons(scanner *bufio.Scanner, n int) []Lesson {
    lessons := make([]Lesson, n)
    for i := 0; i < n; i++ {
        lessons[i] = readLesson(scanner)
    }
    return lessons
}

func readLesson(scanner *bufio.Scanner) Lesson {
    scanner.Scan()
    listLesson := strings.Split(scanner.Text(), " ")
    begin, _ := strconv.ParseFloat(listLesson[0], 64)
    end, _ := strconv.ParseFloat(listLesson[1], 64)
    return Lesson{
        begin,
        end,
    }
}

func printLessons(lessons []Lesson) {
    writer := bufio.NewWriter(os.Stdout)
    for _, lesson := range lessons {
        writer.WriteString(strconv.FormatFloat(lesson.begin, 'g', -1, 64))
        writer.WriteString(" ")
        writer.WriteString(strconv.FormatFloat(lesson.end, 'g', -1, 64))
        writer.WriteString("\n")
    }
    writer.Flush()
}