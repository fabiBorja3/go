package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type Estructura struct {
	mapa map[int]Persona
	mutex sync.Mutex
}

type IntRango struct {
	min, max int
}


type Persona struct {
	documento string
	nombre string
	fondos [5] float64
}

var mapPersonas = Estructura{mapa: make(map[int]Persona)}

func main () {
	rand.Seed(time.Now().UnixNano())
	crearPersonas()
	fmt.Println(mapPersonas)
	wg := &sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < len(mapPersonas.mapa); i++ {
	  go fondos(mapPersonas,i, wg)
	}
	wg.Wait()
	fmt.Println("-----------------------------")
	fmt.Println("MAPPERSONAS FINAL ",mapPersonas)

	var espera string
	fmt.Scan(&espera)
}

func fondos(mapPersonas Estructura,id int, wg *sync.WaitGroup) {
	canal := make(chan float64,1)
	wgFondo := &sync.WaitGroup{}
	wgFondo.Add(5)
	persona := mapPersonas.mapa[id]
	for j := 0; j < len(persona.fondos); j++ {
		go inversion(canal, persona.fondos[j],wgFondo)
	}

	for j := 0; j < len(persona.fondos); j++ {
		inv := <- canal
		persona.fondos[j] = inv
	}
	wgFondo.Wait()
	mapPersonas.mutex.Lock()
	mapPersonas.mapa[id] = persona
	mapPersonas.mutex.Unlock()
	wg.Done()
}



func inversion(canal chan float64, fondo float64, wgFondo *sync.WaitGroup) {
		aleatorio := randFloats(-10,10)
		canal <- fondo * aleatorio
	    wgFondo.Done()
}


func randFloats(min, max float64) float64 {
        return (min + rand.Float64()*(max-min))
	
}


func crearPersonas(){

	persona1 := new(Persona)
	persona1.nombre = "Juan"
	persona1.documento = "12345"
	persona1.fondos = [5]float64{30000, 25000,18000,30000,20000}

	persona2 := new(Persona)
	persona2.nombre = "Tony"
	persona2.documento = "45678"
	persona2.fondos = [5]float64{40000, 28000,18000,20000,26000}

	persona3 := new(Persona)
	persona3.nombre = "Ana"
	persona3.documento = "34567"
	persona3.fondos = [5]float64{35000, 15000,18000,26000,30000}

	persona4 := new(Persona)
	persona4.nombre = "Lina"
	persona4.documento = "56789"
	persona4.fondos = [5]float64{21000, 35000,18000,18000,28000}

	persona5 := new(Persona)
	persona5.nombre = "Ricardo"
	persona5.documento = "23456"
	persona5.fondos = [5]float64{30000, 23000,18000,25000,22000}

	mapPersonas.mapa[0] = *persona1
	mapPersonas.mapa[1] = *persona2
	mapPersonas.mapa[2] = *persona3
	mapPersonas.mapa[3] = *persona4
	mapPersonas.mapa[4] = *persona5


}

