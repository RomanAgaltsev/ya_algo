package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetNearestZeros(houseNumbers []int) []int {
	// Слайс для результата
	nearetsZeros := make([]int, len(houseNumbers))
	// Если улица пустая - вернем пустой слайс
	if len(nearetsZeros) == 0 {
		return nearetsZeros
	}
	// Указатели для шагания по слайсу
	// Оба начинаются с -1 - указывают на позицию до начала слайса
	prev, next := -1, -1
	// Если на улице первый же участок пустой, шагнем обоими указателями
	if houseNumbers[0] == 0 {
		prev, next = 0, 0
		// Сразу запишем ближайший 0 для пустого участка
		nearetsZeros[0] = 0
	}
	// Флаг, определяющий, что next дошел до конца слайса
	// Нужен, чтобы отличить ситуации:
	// - next дошел до очередного 0 (пустого участка)
	// - next дошел до конца слайса, а 0 (пустой участок) не нашелся
	nextIsStoped := false

	fmt.Println("next начинает шагать:", next)
	// next начинает шагать
	for next < len(houseNumbers)-1 {
		// next делает 1 шаг
		next += 1
		fmt.Println("next шагает:", next)
		// Проверяем, не дошел ли next до конца слайса
		if next == len(houseNumbers)-1 {
			// next дошел до конца слайса
			nextIsStoped = true
			fmt.Println("next дошел до конца:", next)

		}
		// Проверяем, не нашел ли next 0 или не дошел ли до конца
		if houseNumbers[next] == 0 || nextIsStoped {
			// next нашел 0 или дошел до конца слайса
			fmt.Println("next нашел 0 или дошел до конца:", next)

			fmt.Println("prev начинает шагать:", prev)
			// prev начинает шагать
			for i := prev + 1; i <= next; i++ {
				fmt.Println("prev шагает:", i)
				// prev делает 1 шаг

				var prevDist = i - prev
				var nextDist int
				if nextIsStoped {
					nextDist = i - prev
				} else {
					nextDist = next - i
				}

				if houseNumbers[i] == 0 {
					nearetsZeros[i] = 0
				} else if prevDist < nextDist {
					nearetsZeros[i] = prevDist
				} else {
					nearetsZeros[i] = nextDist
				}
				switch {
				case houseNumbers[i] == 0:
					nearetsZeros[i] = 0
				case nextIsStoped && houseNumbers[next] == 0:
					prevDist := i - prev
					nextDist := next - i
					if prevDist < nextDist {
						nearetsZeros[i] = prevDist
					} else {
						nearetsZeros[i] = nextDist
					}
				default:
					nearetsZeros[i] = i - prev
				}

			}
			// prev прошел все шаги до next-1
			// prev шагает до next
			//prev = next
			//fmt.Println("prev приравняли к next:", prev, next)
		}
	}
	return nearetsZeros
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	houseNumbers := readArray(scanner)
	fmt.Println(GetNearestZeros(houseNumbers))
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
