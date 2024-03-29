package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func factorize(number int) []int {
	factors := make([]int, 0)
	for number > 0 {
		if number == 2 {
			factors = append(factors, 2)
			number = 0
			break
		}
		if number % 2 == 0 {
			factors = append(factors, 2)
			number /= 2
			continue
		}
		hasPrimeFactor := false
		for i := 3; i*i <= number; i += 2 {
			if number%i == 0 {
				factors = append(factors, i)
				number = number / i
				hasPrimeFactor = true
				break
			}
		}
		if !hasPrimeFactor {
			factors = append(factors, number)
			break
		}
	}
	sort.Ints(factors)
	return factors
}

func main() {
	scanner := makeScanner()
	number := readInt(scanner)
	factorization := factorize(number)
	printArray(factorization)
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
