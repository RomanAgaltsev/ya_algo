package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func movingAverage(array []int, windowSize int) []float64 {
	res := make([]float64, len(array) - windowSize + 1)
	currentSum := 0
	for i := 0; i < windowSize; i++ {
		currentSum += array[i]
	}
	res[0] = float64(currentSum) / float64(windowSize)
	for i := 0; i < len(array) - windowSize; i++ {
		currentSum -= array[i]
		currentSum += array[i+windowSize]
		currentAvg := float64(currentSum) / float64(windowSize)
		res[i+1] = currentAvg
	}
	return res
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	windowsSize := readInt(scanner)
	printArray(movingAverage(array, windowsSize))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []float64) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.FormatFloat(arr[i], 'f', 8, 64))
		writer.WriteString(" ")
	}
	writer.Flush()
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