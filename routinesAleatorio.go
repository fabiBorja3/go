package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){

	fmt.Println("Inicio")
	n := rand.Intn(30 - 10) + 10
	//fmt.Println("n: ",n)

	for i:= 1; i <= n; i++ {
		fmt.Println("Lanzando Go routine ", i)
		go func(num int) {
			m := rand.Intn(5 - 1) + 1
			//fmt.Println("m: ",m)
			time.Sleep(time.Duration(m) * time.Second)
			fmt.Println("Go routine # %d", num)
		}(i)
	}

	fmt.Print("Fin")
	var espera string
	fmt.Scanln(&espera)

}