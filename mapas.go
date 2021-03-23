package main

import "fmt"

func main7(){

	monedas := make(map[string]string)

	monedas["colombia"] = "Peso Colombiano"
	monedas["EEUU"] = "Dolar"
	monedas["España"] = "Euro"

	delete(monedas,"España")

	fmt.Println(monedas)
	fmt.Print(len(monedas))

	for key, value:= range monedas {
		fmt.Println("Clave; ", key, "Valor", value)
	}

	arreglo :=[]int{10,20,30,40,50}

	//Range devuelve dos objetos, el primero es el indice el segundo el valor, con el guion bajo 
	//decimos que no necesitamos esa variable
	for i, _ := range arreglo {
		fmt.Println("Elemento:", i)
	}
}