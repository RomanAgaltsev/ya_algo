package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func getBoundaries(arr []int, begin, end int) (int, int) {
	if begin + 1 == end {
		return end, begin
	}
	middle := begin + (end - begin) / 2
	if arr[middle] > arr[begin] {
		return getBoundaries(arr, middle, end)
	} else {
		return getBoundaries(arr, begin, middle)
	}
}

func binarySearch(arr []int, x, left, right int) int {
	if right <= left {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == x {
		return mid
	} else if x < arr[mid] {
		return binarySearch(arr, x, left, mid)
	} else {
		return binarySearch(arr, x, mid+1, right)
	}
}

func brokenSearch(arr []int, k int) int {
	var begin, end int
	if len(arr) == 1 || arr[0] < arr[len(arr)-1] {
		begin, end = 0, len(arr)-1
	} else {
		begin, end = getBoundaries(arr, 0, len(arr)-1)
	}
	if k >= arr[0] && k <= arr[end] {
		return binarySearch(arr, k, 0, end+1)
	} else {
		return binarySearch(arr, k, begin, len(arr))
	}
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	k := readInt(scanner)
	arr := readArray(scanner)
	fmt.Print(brokenSearch(arr, k))
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