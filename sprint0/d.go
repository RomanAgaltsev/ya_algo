package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(array []int, targetSum int) []int {
	mp := make(map[int]int, len(array))
	for i, v := range array {
		if _, ok := mp[targetSum - v]; ok {
			res := make([]int, 2)
			res[0] = targetSum - v
			res[1] = v
			return res
		}
		mp[v] = i
	}
	return make([]int, 0)
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	targetSum := readInt(scanner)
	result := twoSum(array, targetSum)
	if len(result) == 0 {
		fmt.Println("None")
	} else {
		fmt.Print(result[0])
		fmt.Print(" ")
		fmt.Print(result[1])
	}
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
	for i :=0; i < len(listString); i++ {
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