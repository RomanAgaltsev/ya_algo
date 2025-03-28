package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Если ответ существует, верните список из двух элементов
// Если нет - то верните пустой список
func twoSum(array []int, targetSum int) []int {
	sort.Ints(array)
	left := 0
	right := len(array) - 1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == targetSum {
			res := make([]int, 2)
			res[0] = array[left]
			res[1] = array[right]
			return res
		} else if currentSum < targetSum {
			left += 1
		} else {
			right -= 1
		}
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