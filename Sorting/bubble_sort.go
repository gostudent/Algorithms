package main

import (
	"fmt"
)

func Bubblesort(arr []int) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(arr) - 1; i++ {
			if arr[i + 1] < arr[i] {
				Swap(arr, i, i + 1)
				swapped = true
			}
		}
	}	
}

func Swap(arr []int, i, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}

func main() {

	arr := []int{199,66, 52,44, 29, 40, 45, 33, 74, 34}
	fmt.Println("Unsorted array is as follows: ", arr)
	Bubblesort(arr)
	fmt.Println("Sorted array is as follows: ", arr)
}