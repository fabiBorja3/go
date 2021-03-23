package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	data, err := ioutil.ReadFile("./hola.go")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(data)
		fmt.Println(string(data))
	}
}