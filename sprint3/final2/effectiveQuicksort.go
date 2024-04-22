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
	"cmp"
	"os"
	"strconv"
	"strings"
)

type Participant struct {
	login string
	solve int
	penalty int
}

func less(a, b Participant) int {
//	if a.solve > b.solve {
//		return -1
//	} else if a.solve < b.solve {
//		return 1
//	}
//	if a.penalty < b.penalty {
//		return -1
//	} else if a.penalty > b.penalty {
//		return 1
//	}
	if res := cmp.Compare(b.solve, a.solve); res != 0 {
		return res
	}
	if res := cmp.Compare(a.penalty, b.penalty); res != 0 {
		return res
	}
	return strings.Compare(a.login, b.login)
}

func partition(participants []Participant, left, right int) (int,int) {
	pivotIndex := left + (right - left) / 2
	pivot := participants[pivotIndex]
	for left < right {
		for less(participants[left], pivot) == -1 {
			left++
		}
		for less(pivot, participants[right]) == -1 {
			right--
		}
		if left < right {
			participants[left], participants[right] = participants[right], participants[left]
		}
	}
	return left, right
}

func quickSort(participants []Participant, left, right int) {
	if left >= right {
		return
	}
	lb, rb := partition(participants, left, right)
	if lb == rb {
		quickSort(participants, left, lb-1)
		quickSort(participants, rb+1, right)
	} else {
		quickSort(participants, left, lb)
		quickSort(participants, rb, right)
	}

}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	participants := readParticipants(scanner, n)
	quickSort(participants, 0, len(participants)-1)
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