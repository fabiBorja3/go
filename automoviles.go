package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	leerArchivoVehiculos()
}

type vehiculo struct{
	modelo string
	marca string
	ano string
	precio string
	color string
	impuesto float32
	estadoVehiculo string
}

type Vehiculo interface{
	Encender() string 
	CalcularImpuesto() float32
}

type Moto struct {
	vehiculo
	Cilindrage string
	impuesto float32
}

type Auto struct {
	vehiculo
	puertas string
	Transmision string
}



func (m Moto) Encender() string{
	return "Encendiendo Moto..."
}

func (a Auto) Encender() string{
	return "Encendiendo Auto..."
}

func (a Auto) EstadoVehiculo(i int) string{
	if i % 2 == 0 {
	return "Funcional"
	}else{
	return "Fallando"
	}
}

func (m Moto) EstadoVehiculo(i int) string{
	if i % 2 == 0 {
	return "Funcional"
	}else{
	return "Fallando"
	}
}

func (m Moto) calcularImpuesto(){
		value, err := strconv.ParseFloat(m.precio,32)
	if err != nil {
		fmt.Println(err)
	}
	m.impuesto = float32(value) * (float32(46.5)/100)
}

func (a Auto) calcularImpuesto() {
		value, err := strconv.ParseFloat(a.precio,32)
	if err != nil {
		fmt.Println(err)
	}
	a.impuesto = float32(value) * (float32(46.5)/100)
}


func leerArchivoVehiculos() {
	archivo, err := os.Open("./vehiculos.txt")

	defer func() {
		fmt.Print("Se cierra con defer funcion que se lanza cuando termina el metodo donde fue definido")
		archivo.Close()
	}()


	if err != nil {
		fmt.Println("Algo salio mal")
		panic(err)
	}

	scanner := bufio.NewScanner(archivo)
	autosMap := make(map[int]Auto)
    motosMap := make(map[int]Moto)
	a := new(Auto)
	m :=new(Moto)
	for scanner.Scan() {
		linea := scanner.Text()
		if strings.Contains(linea,"Carros") == true {
			j := 0
			for scanner.Scan() {
			linea := scanner.Text()
			if strings.Contains(linea,"Motos") != true {
				a := a.separararInfoAutos(linea)
				a.estadoVehiculo = a.EstadoVehiculo(j)
				autosMap[j] = a
				j++
			}else{
		      if strings.Contains(linea,"Motos") == true {
				k :=0
				for scanner.Scan() {
					linea := scanner.Text()
					m := m.separararInfoMotos(linea)
					m.estadoVehiculo = m.EstadoVehiculo(k)
					motosMap[k] = m
					k++
				}
				fmt.Println(motosMap)
			}
			}
		}
	}
	//fmt.Println(autosMap)
	//fmt.Println(motosMap)
	break
 }
}

func (a Auto) separararInfoAutos(frase string) Auto{
	fmt.Println("carros -> %s",frase)
	
	resultado := strings.Split(frase, ",")
	a.marca = resultado[0]
	a.modelo = resultado[1]
	a.ano = resultado[2]
	a.precio =strings.TrimSpace(resultado[3])
	a.puertas = resultado[4]
	a.color = resultado[5]
	a.Transmision = resultado[6]
	a.calcularImpuesto()

	return a
}

func (m Moto)separararInfoMotos(frase string) Moto{
	fmt.Println("motos -> %s",frase)

	resultado := strings.Split(frase, ",")
	m.marca = resultado[0]
	m.modelo = resultado[1]
	m.ano = resultado[2]
	m.Cilindrage = resultado[3]
	m.precio = strings.TrimSpace(resultado[4])
	m.color = resultado[5]
	m.calcularImpuesto()

	return m

}




