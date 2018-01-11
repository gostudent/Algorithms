package main

import "fmt"

func main() {
	arrayz := [10]int{1, 2, 31, 43, 12,23,44,443,56,54}
    fmt.Println("Initial array is as follows:", arrayz)
    fmt.Println("")

    var mini int = 0
    var tmp int = 0

    for i := 0; i < len(arrayz); i++ {
        mini = i
        for j := i + 1; j < len(arrayz); j++ {
            if arrayz[j] < arrayz[mini] {
                mini = j
            }
        }
		
        tmp = arrayz[i]
        arrayz[i] = arrayz[mini]
        arrayz[mini] = tmp
    }

    fmt.Println("Sorted array is as follows:    ", arrayz)
}