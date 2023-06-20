package main

import (
	"fmt"
)

/* package level scope */
// declarando la variable con el identificador z del tipo int
// la declaracion es valida, la asignacion no dud
var z int

func main(){
	/* 
	El valor 0, o valores por defecto de un tipo
	bools => false
	int => 0
	float => 0.0
	strings => ""
	nil para => pointers, funciones, interfaces, slices, channels, maps
	Usar el operador de declaracion corta de variables mientras mas sea posible
	*/
	//z = 985
	fmt.Println(z)
}

