package main

import "fmt"
import "math/rand"
func main() {
	arrayz := Random_Arr(10)
    fmt.Println("Initial array is as follows:", arrayz)
    fmt.Println("")

    for d := int(len(arrayz)/2); d > 0; d /= 2 {
    	for j := d; j < len(arrayz); j++ {
    		for k := j; k >= d && arrayz[k-d] > arrayz[k]; k -= d {
    			arrayz[k], arrayz[k-d] = arrayz[k-d], arrayz[k]
    		}
    	}
    }

    fmt.Println("Sorted array is as follows: ", arrayz) 	
}
func Random_Arr(n int) []int {
    arrayz := make([]int, n)
    for i := 0; i <= n - 1; i++ {
        arrayz[i] = rand.Intn(n)
    }
    return arrayz
}