package main

import "fmt"
import "math/rand"


func Quick_Sorting(arrayz []int) []int {
    
    if len(arrayz) <= 1 {
        return arrayz
    }
    median := arrayz[rand.Intn(len(arrayz))]
    low := make([]int, 0, len(arrayz))
    high := make([]int, 0, len(arrayz))
    middle := make([]int, 0, len(arrayz))
    for _, item := range arrayz {
        switch {
        case item < median:
            low = append(low, item)
        case item == median:
            middle = append(middle, item)
        case item > median:
            high = append(high, item)
        }
    }
 
    low = Quick_Sorting(low)
    high = Quick_Sorting(high)

    low = append(low, middle...)
    low = append(low, high...)
 
    return low
}

func main() {
    arrayz := RandomArr(10)

    fmt.Println("Initial array :", arrayz)
    fmt.Println("")
    fmt.Println("Sorted array : ", Quick_Sorting(arrayz))
}

func RandomArr(n int) []int {
    arrayz := make([]int, n)
    for i := 0; i <= n - 1; i++ {
        arrayz[i] = rand.Intn(n)
    }
    return arrayz
}