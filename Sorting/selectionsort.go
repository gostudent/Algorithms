package main

import "fmt"

func Swap(array []int, i, j int) {
	tmp := array[j]
	array[j] = array[i]
	array[i] = tmp
}

func selectionSort(array []int) {
	var min_index int
	length_of_array := len(array)
	for i := 0; i < length_of_array - 1; i++ {
		//finding element with minimum value on right side of i
		min_index = i
		for j := i+1; j < length_of_array; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		// Swapping
		Swap(array, min_index, i)
	}
}

func main() {

	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", array)
	selectionSort(array)
	fmt.Println("Sorted array: ", array)
}