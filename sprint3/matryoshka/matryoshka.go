package main

import "fmt"

func buildMatryoshka(size, n int) {
	if n >= 1 {
		fmt.Printf("Создаем низ матрешки размера %d.\n", size)
		buildMatryoshka(size-1, n-1)
		fmt.Printf("Создаем верх матрешки размера %d.\n", size)
	} else {
		return
	}
}

func main() {
	buildMatryoshka(4, 3)
}