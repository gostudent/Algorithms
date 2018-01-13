package main

import (
	"fmt"	
       )

func main() {
		A := []int{6, 1, 1, 3, 2, 9, 0, 5, 7}
	        sorted := countingSort(A, 10)
	        fmt.Println("A  after counting sort:", sorted)		
	    }

/**
countingSort sorts input array *A* using counting sort
Time complexity: O(len(A) + k)
*/
//parameter A: input array
//parameter k: values of *A* should be in *0..k*

func countingSort(A []int, k int) []int {
	C := make([]int, k+1)
	B := make([]int, len(A))

	for j := 0; j < len(A); j++ {
		C[A[j]]++
	}

	for i := 1; i <= k; i++ {
		C[i] += C[i-1]
	}

	for j := len(A) - 1; j >= 0; j-- {
		n := A[j]
		B[C[n]-1] = n
		C[n]--
	}

	return B
}

