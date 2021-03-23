package main

import (
	"fmt"
	"strings"
)

func main5(){

	frase := "ESTA ES UNA FRASE DE PRUEBA PARA LA FRASE PRUEBA PRUEBA"
	palabras := make(map[string]int)
	palabras = contarPalabras(frase)
	palabraMasRepetida(palabras)
	
}

func contarPalabras(frase string)map [string]int{

		palabras := make(map[string]int)
		resultado := strings.Split(frase, " ")
	for i := 0; i < len(resultado); i++ {
		if(palabras[resultado[i]] != 0){
			palabras[resultado[i]] = palabras[resultado[i]] + 1;
		}else{
			palabras[resultado[i]] = 1;
		}
				
	}

	return palabras
}

func palabraMasRepetida(palabras map [string]int){
	
	mayor := 0
	palabra := ""
	for key, valor := range palabras {

		if palabras[key] > mayor {
			mayor = valor
			palabra = key
		}
	}

	fmt.Println("Palabra mas repetida")
	fmt.Println(palabra)
	fmt.Println(mayor)


}