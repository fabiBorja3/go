package main

import (
	"fmt"
	"strconv"
)


var var1 bool
type Persona struct {
	nombre, apellido string
	edad int
}

type empleado struct {
	idEmpleado string
	fecha string
}

func (p *Persona) toString() string {
	return "Nombre: " + p.nombre + p.apellido+ ", Edad:" + strconv.Itoa(p.edad)
}

func (e *empleado) registrarIngreso(){
	e.fecha ="12/12/2021"
}

type Usuario struct {
	Persona
	empleado
	email string
	password string
}

func (u *Usuario) imprimiDatos() string{
	return "Nombre: " + u.nombre + " Apellido: "+ u.apellido+", Fecha: "+ u.fecha
}

func(u *Usuario) iniciarSesion() bool {
	return true
}

func main() {
	usuario := Usuario{Persona{nombre: "Juan"}, empleado{"KDIFDS","21/03/2021"},"miborja@gmail.com",""}
	fmt.Println(usuario)
	fmt.Println(usuario.nombre)
	fmt.Println(usuario.idEmpleado)
	usuario.registrarIngreso()
	fmt.Println(usuario.imprimiDatos())
}