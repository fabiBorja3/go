package main

import "fmt"

type Animal interface{

	Mover() string 
}

type Perro struct {
	nombre string
	patas string
}

type Serpiente struct {
	nombre string
}

func (p Perro) Mover() string{
	return "Caminando... moviendo" + p.patas +" patas"
}

func (s Serpiente) Mover () string{
	return "Reptando..."
}

func iniciar(animal Animal) {
	fmt.Println(animal.Mover())
}

func main() {
	perro := Perro{"sdfasfa", "4"}
	serpiente := Serpiente{"serpiente"}

	iniciar(perro)
	iniciar(serpiente)

}