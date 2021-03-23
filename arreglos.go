package main

import "fmt"

func main1(){

	//arreglo3 := [3]string{"A", "b","C"}
	var arreglo [5] int
	//matris var arreglo[4][5]
	arreglo[0] = 12

	fmt.Println(arreglo)

	for i := 0; i < len(arreglo); i++ {
		fmt.Print(arreglo)		
	}

}


