package main

import (
	"fmt"
	"time"
)

//import {
//	"ftm"
//}

func main(){

	fmt.Println("Inicio")
	for i:= 1; i <= 20; i++ {
		fmt.Println("Lanzando Go routine ", i)
		go func(num int) {
			time.Sleep(2 * time.Second)
			fmt.Println("Go routine # %d\n", num)
		}(i)
	}

	fmt.Print("Fin")
	var espera string
	fmt.Scanln(&espera)


}