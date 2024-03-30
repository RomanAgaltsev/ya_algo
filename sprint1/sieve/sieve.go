package main

import "fmt"

func eratosthenes(n int) []bool {
	numbers := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		numbers[i] = true
	}
	for num := 2; num < n; num++ {
		if numbers[num] {
			for j := 2 * num; j <= n; j += num {
				numbers[j] = false
			}
		}
	}
	numbers[0] = false
	numbers[1] = false
	return numbers
}

func eratosthenesEffective(n int) []bool {
	numbers := make ([]bool, n+1)
	for i := 2; i <= n; i++ {
		numbers[i] = true
	}
	for num := 2; num * num < n; num++ {
		if numbers[num] {
			for j := num * num; j <= n; j += num {
				numbers[j] = false
			}
		}
	}
	numbers[0] = false
	numbers[1] = false
	return numbers
}

func main() {
//	for i, b := range eratosthenes(20) {
//		fmt.Printf("int %v is prime %v\n", i, b)
//	}
	for i, b := range eratosthenesEffective(30) {
		fmt.Printf("int %v is prime %v\n", i, b)
	}
}