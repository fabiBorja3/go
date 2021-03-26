package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//Se crea un tipo Estructura
type Estructura struct {
	mapa map[string]int
	mutex sync.Mutex
}

//Se creo una variable definida con la estructura creada
var estructura = Estructura{mapa: make(map[string]int)}


func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go guardar(fmt.Sprintf("obj_%d", rand.Intn(100)), rand.Intn(100),wg)
	}
	go leerDato(fmt.Sprintf("obj_%d",rand.Intn(100)))
	wg.Wait()
	fmt.Println(estructura.mapa)

}

func guardar(clave string, valor int, wg *sync.WaitGroup) {
	estructura.mutex.Lock()
	estructura.mapa[clave] = valor
	estructura.mutex.Unlock()
	wg.Done()
}

func leerDato(clave string){
	estructura.mutex.Lock()
	fmt.Printf("Leyendo Mapa en [%s] -> %d\n",clave, estructura.mapa[clave])
	estructura.mutex.Unlock()
}

