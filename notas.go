package main

import (
	"fmt"
	"strconv"
)

func main(){
	var arreglo [5][3] string
	for i := 0; i < 5; i++ {
	var nombre string
	var edad string
	var nota string
	var x int
    fmt.Println("Ingresa nombre del estudiante:")
	fmt.Scanf("%s\n",&nombre)
	arreglo  [i][x] = nombre
	x++
	fmt.Println("Ingresa edad del estudiante:")
	fmt.Scanf("%s\n",&edad)
	arreglo  [i][x] = edad
	x++
	fmt.Println("Ingresa nota del estudiante:")
	fmt.Scanf("%s\n",&nota)
	arreglo  [i][x] = nota
	x++
	fmt.Println(arreglo)
	}

	CalcularPromedio(arreglo)
}

	func CalcularPromedio(valor[5][3] string){

		var edadPromedioTemp float32
		var notaPromedioTemp float32
		for i := 0; i < 5; i++ {
			
			edad, err := strconv.Atoi(valor[i][1])
			if err != nil {
				//fmt.Println("Error en edad del estudiante Nro %s", err)
				fmt.Println(err)
			}
			edadPromedioTemp = edadPromedioTemp + float32(edad)

			nota, err := strconv.Atoi(valor[i][2])
			if err != nil {
				//fmt.Println("Error en nota del estudiante Nro %d", )
				fmt.Println(err)
			}
			notaPromedioTemp = notaPromedioTemp + float32(nota)

		}

		fmt.Println(notaPromedioTemp)

		fmt.Println(edadPromedioTemp)

		fmt.Println(float32(edadPromedioTemp/float32(5)))
		fmt.Println(float32(notaPromedioTemp/float32(5)))
		
	}