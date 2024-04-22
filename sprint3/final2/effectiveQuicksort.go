/*
Яндекс Практикум
Алгоритмы и структуры данных
58 когорта
Агальцев Роман

Спринт 3
Задача - A. Поиск в сломанном массиве

Отчеты:
- Ревью 1 -

-- ПРИНЦИП РАБОТЫ --

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --


*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Participant struct {
	login string
	solve int
	penalty int
}

func less(a, b Participant) bool {
	return false
}

func partition(participants []Participant, left, right int) int {
	pivot := participants[(left + right + 1) / 2]
	i, j := left, right
	for {
		for i < j && less(participants[i], pivot) {
			i++
		}
		for j > i && less(pivot, participants[j]) {
			j--
		}
		if i < j {
			participants[i], participants[j] = participants[j], participants[i]
			i++
			j--
		} else {
			return j
		}
	}
}

func quickSort(participants []Participant, left, right int) {
	if left <= right {
		return
	}
	j := partition(participants, left, right)
	quickSort(participants,left,j)
	quickSort(participants,j,right)
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	participants := readParticipants(scanner, n)
	quickSort(participants, 0, len(participants))
	printParticipants(participants)
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
func readParticipants(scanner *bufio.Scanner, n int) []Participant {
	participants := make([]Participant, n)
	for i := 0; i < n; i++ {
		participants[i] = readParticipant(scanner)
	}
	return participants
}

func readParticipant(scanner *bufio.Scanner) Participant {
	scanner.Scan()
	listParticipant := strings.Split(scanner.Text(), " ")
	login := listParticipant[0]
	solve, _ := strconv.Atoi(listParticipant[1]) 
	penalty, _ := strconv.Atoi(listParticipant[2])
	return Participant{
		login,
		solve,
		penalty,
	}
}

func printParticipants(participants []Participant) {
	writer := bufio.NewWriter(os.Stdout)
	for _, participant := range participants {
		writer.WriteString(participant.login)
		writer.WriteString("\n")
	}
	writer.Flush()
}