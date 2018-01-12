//Program in Go to implement Linear Search
package main

import "fmt"

func linearSearch(array []int, query int) int {
	for i, item := range array {
		if item == query {
			return i
		}
	}
	return -1
}

func main() {
	
	array := []int{44, 23, 424, 633, 82, 102, 124, 148, 186, 168}
	index := linearSearch(array, 102)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("array[", index, "] = ", array[index])
	}
}
