package main

import "fmt"

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

func binarySearchDescending(arr []int, x, left, right int) int {
	if right <= left {
		return -1
	}
	mid := (right + left) / 2
	if arr[mid] == x {
		return mid
	} else if arr[mid] < x {
		return binarySearchDescending(arr, x, left, mid)
	} else {
		return binarySearchDescending(arr, x, mid+1, right)
	}
}

func main() {
	s1 := []int{1, 2, 4, 7, 9, 21}
	fmt.Print(binarySearch(s1, 9, 0, len(s1)))
	
	s2 := []int{21, 9, 7, 4, 2, 1}
	fmt.Print(binarySearchDescending(s2, 9, 0, len(s2)))
}