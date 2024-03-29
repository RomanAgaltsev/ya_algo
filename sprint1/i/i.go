package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPowerOfFour(number int) bool {
	pow := 1
	for pow < number {
		pow *= 4
	}
	return pow == number
}

func main() {
	scanner := makeScanner()
	number := readInt(scanner)
	if isPowerOfFour(number) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
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
