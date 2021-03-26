package main

import (
	"fmt"
	"sync"
	"time"
)




func main(){
	
	pedido()
}

func pedido(){

	fmt.Println("Por Favor Realice su pedido")
	var ops,conf, cantPedidos int 
	var confirmar bool
	fmt.Println("Por Favor Realice su pedido")
	for !confirmar {

	for i := 0; i < len(productos); i++ {
		fmt.Println(productos[i].id,".Producto: ",productos[i].nombre, "|Precio",productos[i].precio,"|Tiempo de espera: ",productos[i].tiempo)
	}
	
	fmt.Scanf("%d\n", &ops)
	fmt.Println("--------------------------------")
	fmt.Println(ops)

	fmt.Println("Confirmar Pedido?")
	fmt.Println("1.Si")
	mapProdPedido[cantPedidos] = productos[ops-1]
	//mapProdPedido[0] = productos[0]
	//mapProdPedido[1] = productos[1]
	fmt.Scanf("%d\n", &conf)
	if(conf == 1){
		confirmar = true
	}
	//confirmar = true
	cantPedidos++
}
    mostrarElementosProductos(mapProdPedido)
   //wg.Wait()
}


func mostrarElementosProductos(productosPedidos map[int]Producto){


	ingredientes.elementos = make(map[int]Elemento)
	productosVendidos := make(map[int]Producto) 

	for i := 0; i < len(elementos); i++ {
		ingredientes.elementos[i] = elementos[i]
	}

	fmt.Println("ELEMENTOS INICIAL ",elementos)
	canalPedidoPrueba := make(chan Respuesta)
	var respuesta = new(Respuesta)

			for i := 0; i < len(productosPedidos); i++ {
				producto :=productosPedidos[i]
				go prepararProducto(ingredientes , producto, canalPedidoPrueba)
				*respuesta = <- canalPedidoPrueba
			    fmt.Println("RESPUESTA ",respuesta)
				if respuesta.ok == 0 {
		        fmt.Println("Preparando Pedido...")
				wg := &sync.WaitGroup{}
	            wg.Add(1)
				go prepararPedido2(wg, producto.tiempo)
				wg.Wait()
				productosVendidos[i] = producto
				}else{
				fmt.Print(respuesta.mensaje)
				}
			}
			
			var costo float32
			for i := 0; i < len(productosVendidos); i++ {
				costo = costo + productosVendidos[i].precio
			}
			fmt.Println("El valor del pedido es: ",costo)
			





		

}

func prepararPedido2(wg *sync.WaitGroup, tiempo int){
	time.Sleep(time.Duration(tiempo) * time.Second)
	wg.Done()
}

func prepararProducto(ingredientes *Ingredientes, producto Producto, canalPedidoPrueba chan Respuesta){


	//defer wg.Done()
	            respuesta := new(Respuesta)
				for k, elemento := range ingredientes.elementos {
				//fmt.Println("ELEMENTO ...",elemento)


				if elemento.cantidad < producto.elementos[elemento.id] {
			        respuesta.ok = 1
				} else {
					ingredientes.mutext.Lock()
					elemento.cantidad = elemento.cantidad - producto.elementos[elemento.id]
					ingredientes.elementos[k] = elemento
					ingredientes.mutext.Unlock()
					//fmt.Println("ELEMENTO NOMBRE",elemento.nombre)
					//fmt.Println("CANTIDAD SOLICITADA",producto.elementos[elemento.id]) 
					//fmt.Println("CANTIDAD EXISTENCIA",elemento.cantidad)
				}


			}

			if respuesta.ok == 0 {
			        respuesta.producto = producto
			        respuesta.mensaje = "Ok"
			        canalPedidoPrueba <- *respuesta
			}else{
				respuesta.producto = producto
			    respuesta.mensaje = "No existen suficientes elementos para el producto"
				canalPedidoPrueba <- *respuesta
			}

					
		

}


type Elemento struct {
  id int
  nombre string
  cantidad int
  mutext sync.Mutex
}



type Producto struct {
  id int
  nombre string
  precio float32
  tiempo int
  elementos map[int]int
}

type Respuesta struct {
  ok int
  mensaje string
  producto Producto
}

type Ingredientes struct {
	elementos map[int]Elemento
	mutext sync.Mutex
}


var carne = Elemento{
	id:1,
	nombre:  "Carne",
	cantidad: 5,
}

var pan = Elemento{
	id:2, 
	nombre:"Pan",
	 cantidad: 10,
}

 var salchicha = Elemento{
	id:3, 
	nombre: "Salchicha", 
	cantidad:  3,
}
var tomate = Elemento{
	id:4,
	nombre:"Tomate",
	cantidad:5,
}
var lechuga = Elemento{
	id: 5,
	nombre: "Lechuga",
	cantidad: 5,
}
var papa = Elemento{
	id: 6,
	nombre: "Papa",
	cantidad: 10,
}
var arepa = Elemento{
	id: 7,
	nombre: "Arepa",
	cantidad:  3,
}

var hamburguesa = Producto{
  id: 1,
  nombre: "Hamburguesa",
  precio: 10000,
  tiempo: 10,
  elementos: map[int]int{
    1: 1,
    2: 2,
    4: 1,
    5: 1,
    6: 3,
  },
}

var perro = Producto{
  id: 2,
  nombre: "Perro Caliente",
  precio: 7000,
  tiempo: 6,
  elementos: map[int]int{
    3: 1,
    2: 1,
    4: 1,
    5: 1,
    6: 3,
  },
}

var arepaB = Producto{
  id: 3,
  nombre: "Arepa burger",
  precio: 8000,
  tiempo: 7,
  elementos: map[int]int{
    1: 1,
    4: 1,
    5: 1,
    6: 3,
  },
}

var elementos = [7]Elemento{carne,pan,salchicha,tomate,lechuga,papa,arepa}
var productos = [3]Producto{hamburguesa, perro, arepaB}

var mapProdPedido = make(map[int]Producto)
var ingredientes = new(Ingredientes)

