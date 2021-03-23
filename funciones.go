package main

import "fmt"

func main3(){
	var nombre string
	fmt.Println("Ingresa tu nombre:")
	fmt.Scanf("%s",&nombre)
	saludar(nombre)
}

func saludar(nombre string){
	fmt.Println("hola " + nombre)
}