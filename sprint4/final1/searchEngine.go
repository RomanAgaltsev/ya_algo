package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	A = 997
	M = 1000000007
)

func getHash(s string) int {
	hash := 0
	for i := 0; i < len(s); i++ {
		hash = (hash*A%M + int(s[i])) % M
	}
	return hash
}

func getWordsCount(line string) map[int]int {
	wordsCount := make(map[int]int)
	for _, word := range strings.Split(line, " ") {
		wordsCount[getHash(word)]++
	}
	return wordsCount
}

func getUniqueHashes(query string) []int {
	hashes := make(map[int]int)
	for _, word := range strings.Split(query, " ") {
		hash := getHash(word)
		if _, ok := hashes[hash]; !ok {
			hashes[hash]++
		}
	}
	result := make([]int, 0)
	for hash := range hashes {
		result = append(result, hash)
	}
	return result
}

func getQueryRelevance(index []map[int]int, query string) map[int]int {
	result := make(map[int]int)
	queryHashes := getUniqueHashes(query)
	for _, queryHash := range queryHashes {
		for i, document := range index {
			result[i] += document[queryHash]
		}
	}
	return result
}

func main() {
	scanner := makeScanner()
	
	n := readInt(scanner)
	index := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		index[i] = getWordsCount(readLine(scanner))
	}
	
	//fmt.Print(index)
	
	m := readInt(scanner)
	for i := 0; i < m; i++ {
		//printArray(getQueryRelevance(index, readLine(scanner)))
		fmt.Print(getQueryRelevance(index, readLine(scanner)))
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

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.WriteString("\n")
	writer.Flush()
}