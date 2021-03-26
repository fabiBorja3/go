package main

import (
	"fmt"
	"math/rand"
	"sync"
)


type Estructura struct {
	mapa map[string]string
	mutex sync.Mutex
}

var mapaPersonas = Estructura{mapa: make(map[string]string)}


func main() {

	wg := &sync.WaitGroup{}
	wg.Add(200)

	crear(mapaPersonas, wg)

	for i := 0; i < 200; i++ {
	n := rand.Intn(4-1) + 1

	switch n {
	case 1:
		go crear(mapaPersonas, wg)

	case 2:
		indice :=len(mapaPersonas.mapa)
		fmt.Println("INDICE r ",indice)
		r := rand.Intn(int(indice) - 1) +1
		fmt.Println("INDICE r RAMDOM",r)
		//go actualizar(mapaPersonas, string(r), wg)
		
	case 3:
		//go eliminar(mapaPersonas,string(j), wg)
	case 4:
		//go imprimir(mapaPersonas, wg)
	}
	}
	fmt.Println(mapaPersonas)
	wg.Wait()
	fmt.Print("Fin")
	var espera string
	fmt.Scanln(&espera)
}



func crear(mapaPersonas Estructura, wg *sync.WaitGroup){

	mapaPersonas.mutex.Lock()
	fmt.Println("CREAR")
    longitud := len(mapaPersonas.mapa) +1
	fmt.Println("LONGITUD ",longitud)
	mapaPersonas.mapa[string(longitud)] = "persona"+string(longitud)
	mapaPersonas.mutex.Unlock()
    wg.Done()

}

func actualizar(mapaPersonas Estructura, indice string, wg *sync.WaitGroup){
  
	mapaPersonas.mutex.Lock()
	fmt.Println("ACTUALIZAR")
    mapaPersonas.mapa[indice] = "persona"+string(indice)
	mapaPersonas.mutex.Unlock()
    wg.Done()

}


func eliminar(mapaPersonas Estructura, wg *sync.WaitGroup){

   mapaPersonas.mutex.Lock()
   fmt.Println("ACTUALIZAR")

   	indice :=len(mapaPersonas.mapa)
	fmt.Println("INDICE r ",indice)
	j := rand.Intn(int(indice) - 1) +1
	fmt.Println("INDICE r RAMDOM",j)
   if mapaPersonas.mapa[string(indice)] != "" {
	      mapaPersonas.mapa[string(indice)] = ""
   }
   mapaPersonas.mutex.Unlock()
   wg.Done()
}

func imprimir(mapaPersonas Estructura, wg *sync.WaitGroup){
   
   mapaPersonas.mutex.Lock()
   fmt.Println("IMPRIMIR")
   longitud := len(mapaPersonas.mapa)
   mapaPersonas.mapa[string(longitud)] = "persona"+string(longitud)
   mapaPersonas.mutex.Unlock()
   fmt.Println(mapaPersonas)
   wg.Done()

}

