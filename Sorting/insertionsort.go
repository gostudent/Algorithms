package main

import (
	"fmt"
	"algoutils"
)

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j - 1]; j-- {
			algoutils.Swap(array, j, j - 1)
		}
	}
}

func main() {
  
  array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
  fmt.Println("Unsorted array: ", array)
  insertionSort(array)
  fmt.Println("Sorted array: ", array)
}
