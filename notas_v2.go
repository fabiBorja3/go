package main

import (
	"fmt"
)

//Definiendo tipo de la nota
//type Nota float32

/*Definiendo slice notas
type Notas struct {
notas map[string]Nota
*/

//Definiendo estudiante
type Estudiante struct{
nombre string
edad int
materia string
notas map[int] float32
promedio float32
}

/*
func (nota Nota) aFloat() float32 {
return float32(nota)
}
*/

func (estudiante Estudiante) CalcularPromedio() float32 {

	var promedioTemp float32
	for i := 1; i < len(estudiante.notas)+1; i++ {
		promedioTemp = promedioTemp + estudiante.notas[i]
	}

	 return promedioTemp/3

	}



func main(){

	var cantidadEstudiantes int
	estudiantesMap := make(map[int]Estudiante)
	fmt.Println("Ingresa la cantidad del estudiantes:")
	fmt.Scanf("%d\n",&cantidadEstudiantes)

	for i := 0; i < cantidadEstudiantes; i++ {
		
			
			estudiante := new(Estudiante)
			fmt.Println("Ingresa nombre del estudiante:")
			fmt.Scanf("%s\n",&estudiante.nombre)
			//fmt.Println(estudiante.nombre)

			fmt.Println("Ingresa edad del estudiante:")
			fmt.Scanf("%d\n",&estudiante.edad)
			//fmt.Println(estudiante.edad)

			var materia string
			fmt.Println("Ingrese materia:")
			fmt.Scanf("%s\n",&materia)
			estudiante.materia = materia
			//fmt.Println(estudiante.materia)

			notasMap := make(map[int]float32)
			for j := 1; j < 4; j++ {
			var notasTemp float32
			
			fmt.Println("Ingrese notas:")
			fmt.Scanf("%g\n",&notasTemp)
			//fmt.Println(notasTemp)
			notasMap[j] = float32(notasTemp)
			//fmt.Println(estudiante.notas[j])
			}
			estudiante.notas = notasMap
			


			estudiante.promedio = estudiante.CalcularPromedio()
			estudiantesMap[i] = *estudiante
		}
		
		for k := 0; k < len(estudiantesMap); k++ {
			println(estudiantesMap[k].nombre)
			println(estudiantesMap[k].edad)
			println(estudiantesMap[k].materia)
			println(estudiantesMap[k].promedio)
			println("-----------------------")
		}
		


}




			/*

			notasMap := make(map[int]float32)
			for i := 1; i < 4; i++ {
				fmt.Println("Ingresa nota:")
				notasMap[i] = rand.Float32()*5
				fmt.Println(notasMap[i])
			}
			*/

			//estudiante.notas = notasMap