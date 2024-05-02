package main

import (
	"bufio"
	"cmp"
	"os"
	"slices"
	"strconv"
	"strings"
)

const docsLimit = 5

func getSearchIndex(scanner *bufio.Scanner, n int) map[string]map[int]int {
	index := make(map[string]map[int]int)
	for i := 0; i < n; i++ {
		for _, word := range strings.Split(readLine(scanner), " ") {
			if wordsCount, ok := index[word]; ok {
				wordsCount[i+1]++
			} else {
				index[word] = map[int]int{i + 1: 1}
			}
		}
	}
	return index
}

func getUniqueWords(query string) []string {
	wordsCount := make(map[string]int)
	for _, word := range strings.Split(query, " ") {
		wordsCount[word]++
	}
	result := make([]string, 0)
	for word := range wordsCount {
		result = append(result, word)
	}
	return result
}

func getQueryRelevance(index map[string]map[int]int, query string) []int {
	relevance := make(map[int]int)
	words := getUniqueWords(query)
	for _, word := range words {
		if docs, ok := index[word]; ok {
			for doc, count := range docs {
				relevance[doc] += count
			}
		}
	}
	relevanceSorted := make([][2]int, 0)
	for doc, rel := range relevance {
		if rel != 0 {
			relevanceSorted = append(relevanceSorted, [2]int{doc, rel})
		}
	}
	slices.SortFunc(relevanceSorted, func(a, b [2]int) int {
		if res := cmp.Compare(b[1], a[1]); res != 0 {
			return res
		}
		if res := cmp.Compare(a[0], b[0]); res != 0 {
			return res
		}
		return 0
	})
	result := make([]int, 0)
	for i := 0; i < docsLimit && i < len(relevanceSorted); i++ {
		result = append(result, relevanceSorted[i][0])
	}
	return result
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	index := getSearchIndex(scanner, n)
	m := readInt(scanner)
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = getQueryRelevance(index, readLine(scanner))
	}
	for i := 0; i < m; i++ {
		printArray(result[i])
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 10 * 1024 * 1024
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