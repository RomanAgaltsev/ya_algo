package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func zip(a []int, b []int) []int {
	res := make([]int, 2 * len(a))
	for i := 0; i < len(a); i++ {
		res[i*2] = a[i]
		res[i*2+1] = b[i]
	}
	return res
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	a := readArray(scanner)
	b := readArray(scanner)
	printArray(zip(a, b))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
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