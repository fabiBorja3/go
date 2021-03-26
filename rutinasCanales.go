package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){
	canal := make(chan string,1)
	canal2 := make(chan string,1)
	canal3 := make(chan string,1)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(3)


	go primeraFuncion(canal, waitGroup)
	fmt.Println(<-canal)
	go segundaFuncion(canal2, waitGroup)
	fmt.Println(<-canal2)
	go terceraFuncion(canal3, waitGroup)
	fmt.Println(<-canal3)

	waitGroup.Wait()



}

func primeraFuncion(chanel chan string, waitGroup *sync.WaitGroup){
	chanel <- "Primera Funcion"
	parcialPrimeraFuncion(waitGroup)
	waitGroup.Done()
}

func segundaFuncion(chanel chan string, waitGroup *sync.WaitGroup){
	chanel <- "Segunda Funcion"
	parcialSegundaFuncion(waitGroup)
	waitGroup.Done()
}

func terceraFuncion(chanel chan string, waitGroup *sync.WaitGroup){
	chanel <- "Tercera Funcion"
	parcialTerceraFuncion(waitGroup)
	waitGroup.Done()
}


func parcialPrimeraFuncion(waitGroup *sync.WaitGroup){
	canal := make(chan string,1)
	waitGroup.Done()
}

func parcialSegundaFuncion(waitGroup *sync.WaitGroup){
	canal := make(chan string,1)
	waitGroup.Done()
}

func parcialTerceraFuncion(waitGroup *sync.WaitGroup){
	time.Sleep(2* time.Second)
	canal := make(chan string,1)
	waitGroup.Done()
}