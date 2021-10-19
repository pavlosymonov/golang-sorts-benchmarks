package main

import (
	"Go-Sorts/sorts"
	"fmt"
)

func main() {
	a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	b := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	c := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	sorts.BubbleSort(a)
	sorts.InsertionSort(b)
	sorts.SelectionSort(c)

	fmt.Printf("Bubble Sort - %v\nInsertion Sort - %v\nSelection Sort - %v\n", a, b, c)
}