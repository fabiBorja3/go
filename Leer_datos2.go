package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)


func main() {
	fmt.Println("Inicio del main")
	ejecutarFuncion()
	fmt.Println("Fin del main")

}

func verDefer(){
	fmt.Println("Inicio del metodo")
	defer fmt.Println("Esto es ejecutando el DEFER")
	fmt.Println("Este es el fin del metodo")
}

func leerArchivo(){
		archivo, err := os.Open("./holaa.go")
	defer func() {
		fmt.Print("Se cierra con defer funcion que se lanza cuando termina el metodo donde fue definido")
		archivo.Close()
		if r := recover(); r != nil{
			fmt.Println(r)
		}
	}()
	if err != nil {
		fmt.Println("Algo salio mal")
		panic(err)
	}
		scanner := bufio.NewScanner(archivo)
		for i := 1; scanner.Scan();{
			linea := scanner.Text()
			fmt.Println("Linea %d -> %s",i, linea)
			i++
			panic(errors.New("Simulando Error"))
		}
}

func ejecutarFuncion(){
	fmt.Println("Inicio de la tercera funcion")
	leerArchivo()
	fmt.Println("Fin de la tercera funcion")
}