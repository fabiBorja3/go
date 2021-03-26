package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	//var r int
	r := rand.Intn(10 - 1) + 1
	arreglo := make([][]int, r)
	for i := 0; i < r; i++ {
		go func (m int)  {
			for j := 0; j < 6; j++ {
			r := rand.Intn(10 - 1) + 1
		    arreglo[m][j] = r
			}
		}(i)
	}
	time.Sleep(2*time.Second)
	fmt.Println(arreglo)

	var espera string
	fmt.Scan(&espera)
}