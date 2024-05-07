/*
E. Сломай меня

Гоша написал программу, которая сравнивает строки исключительно по их хешам.
Если хеш равен, то и строки равны. Тимофей увидел это безобразие и поручил вам сломать программу Гоши, чтобы остальным неповадно было.
В этой задаче вам надо будет лишь найти две различные строки, которые для заданной хеш-функции будут давать одинаковое значение.
Гоша использует следующую хеш-функцию:

для a = 1000 и m = 123 987 123.
В данной задаче необходимо использовать в качестве значений отдельных символов их коды в таблице ASCII.

*/

package main

import (
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func getHash(s string) int {
	a := 1_000
	m := 123_987_123
	hash := 0
	for i := 0; i < len(s); i++ {
		hash = (hash*a%m + int(s[i])) % m
	}
	return hash
}

func getRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func main() {
	n := 20
	//for i := 1; i <= 1000; i++ {
	for {
		s := getRandomString(n)
		t := getRandomString(n)
		if getHash(s) == getHash(t) {
			fmt.Print(s, "\n", t)
			break
		}
	}
}