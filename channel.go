package main

import (
	"fmt"
	"sync"
	"time"
)

//Usa el patron Observer
//Principio de responsabilidad simple, que cada metodo se dedique a una cosa

func main(){
	canal := make(chan int,10)
	salida := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)//Adiciona dos rutinas a esperar
	go publicar(canal, salida, wg)
	go escuchar(canal, salida, wg)
	wg.Wait()//Espera al wg.Done() de todas las rutinas para finalizar
}

//Parametro de canal solo de entrada chan<-
func publicar(canal chan<- int, salida chan struct{}, wg *sync.WaitGroup){
	time.Sleep(2 * time.Second)
	fmt.Println("Publicando...")
	canal <- 10
	canal <- 10
	canal <- 10
	canal <- 10
	canal <- 10
	canal <- 10
	canal <- 10
	canal <- 10
	canal <- 10
	canal <- 10
	salida <- struct{}{}
	wg.Done()
}

//Parametro de canal solo de salida <-chan
func escuchar(canal <-chan int, salida chan struct{}, wg *sync.WaitGroup){
	for  {
		select{
		case dato:= <- canal:
		fmt.Println("Recibido: ", dato)
		case <-salida:
			wg.Done()
		break
		}
	//dato = <-canal	
	//fmt.Println("Recibido: ", dato)
	}
	
	wg.Done()
}