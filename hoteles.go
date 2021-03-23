package main

import (
	"fmt"
	"strconv"
)


type habitacion struct {
	numero string
	piso string
	estado bool
	precio int
}

type hotel struct {
	nombre string
	habitaciones map[string] habitacion
	ingresos int
}

type hoteles struct {
	hoteles map[string]hotel
	
}



func main() {
	h := new(hoteles)
	listarOpciones(h)
}


func listarOpciones(h *hoteles) {
	
	/*
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
    cmd.Stdout = os.Stdout
    cmd.Run()
	*/

	var ops string 

	fmt.Println("--------------------------------")
	fmt.Println("Linea de hoteles:")
	fmt.Println("1.Ver hoteles")
	fmt.Println("2.Registrar hotel")
	fmt.Println("3.Ocupar habitacion")
	fmt.Println("4.Calcular total ingresos")
	fmt.Println("5.Salir")
	fmt.Scanf("%s\n", &ops)
	fmt.Println("--------------------------------")


	fmt.Println(ops)

	switch ops {
	case "1":
		fmt.Println("Ver hoteles")
	    listarHoteles(h)
	case "2":
		fmt.Println("Registar hotel")
		registrarHoteles(h)
	case "3":
		fmt.Println("Ocupar habitacion")
		ocuparHabitacion(h)
	case "4":
		fmt.Println("Calcular total ingreso")
	case "5":
		fmt.Println("Salir")
		
	}
}

func listarHoteles(hoteles *hoteles){
	fmt.Println(hoteles)
	listarOpciones(hoteles)
}

func registrarHoteles(h *hoteles){
	var nombre string
	var pisos, habitaciones, precios int
	hotelMap := make(map[string]habitacion)
	hotelesMap := make(map[string]hotel)
	hotel := new(hotel)
	fmt.Println("Digite el nombre")
	fmt.Scanf("%s\n", &nombre)

	fmt.Println("Digite numero de pisos")
	fmt.Scanf("%d\n", &pisos)

	fmt.Println("Digite habitaciones por piso")
	fmt.Scanf("%d\n", &habitaciones)

	//fmt.Println("PISOS %d ",pisos)
	//fmt.Println("HABITACIONES %d", habitaciones)


	for i := 1; i < pisos+1; i++ {
		for j := 1; j < habitaciones+1; j++ {

			//fmt.Println("CANTIDAD ITERACIONES" +strconv.Itoa(i)+strconv.Itoa(j))
			var comNum string
			if habitaciones < 9 {
				comNum = "0"
			}
			nombreHabitacion := "Habitación "+strconv.Itoa(i)+comNum+strconv.Itoa(j)

			fmt.Println("Digite el precio de la habitacion")
	        fmt.Scanf("%d\n", &precios)

			habitacion := new(habitacion)
			habitacion.piso = strconv.Itoa(i)
			habitacion.numero = nombreHabitacion
			habitacion.precio = precios
			hotelMap[strconv.Itoa(i)+comNum+strconv.Itoa(j)] = *habitacion
			
		}
	}
	hotel.nombre = nombre
	hotel.habitaciones = hotelMap
	hotelesMap[nombre] = *hotel
	h.hoteles = hotelesMap
	
	fmt.Println("----------------")
	fmt.Println("Hotel registrado")
	fmt.Println("----------------")
	listarOpciones(h)
}



func ocuparHabitacion(h *hoteles){
	var nombreHotel string

	fmt.Println("Seleccione el hotel")
	fmt.Println(h.hoteles)
	fmt.Scanf("%s\n", &nombreHotel)
	habitaciones := h.hoteles
	fmt.Println(habitaciones)
}


