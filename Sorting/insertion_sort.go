package main


import "fmt"


func main() {
	arrayz := [9]int{2,1,4,3,5,9,7,6,8}

	for i:=1; i<len(arrayz); i++{
		tem:=arrayz[i]
		j := i

		for;j>0 && arrayz[j-1]>=tem;j--{
			arrayz[j] = arrayz[j-1]
		}
		arrayz[j] = tem
	}

	for Sorted_val:= range arrayz{
		fmt.Println(Sorted_val)
	}
}