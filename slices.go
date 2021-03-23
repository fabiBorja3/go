package main

import "fmt"

func main6(){

	
	//var slice[]int
	
	arreglo := [10]int {1,2,3,4,5,6,7,8,9,10}
	slice := arreglo[:]
    slice2 := make([]int, 3, 10)

	monedas := make(map[string]string)

	slice := monedas

	slice2 = append(slice2,5)
	slice2[3] = 10
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(cap(slice))
	fmt.Println(slice2)

	//fuente := 

	//slice3 := arreglo:5
}