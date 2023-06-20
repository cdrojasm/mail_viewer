package main

import (
	"fmt"
)

func main(){
	/* 
	Operacion de declaracion corta
	...interface{}
	...tipo -> spread operator + tipo
	i.e. 
	int, boolean, ... etc.

	operador declaraciones cortas
	x := 42 -> statement
	x := 2 + 4 -> expression
	2 and 4 are operands
	+ is the operator

	este operador solo es empleable dentro de los cuerpos de las 
	funciones. quiere decir que su scope esta limitado al cuerpo
	de la funcion donde es declarada. 

	:= is only used in statement, can not be used to re asign value
	to a previous declared variable. To re asign value to a identifier
	should be used operator =.

	keywords resserved and can not be used as identifiers. thats
	mean that can be used in statements

	break, default, func, go, map, type, var, return ......

	*/

	x := 42 
	x = 5
	y := 2 + 4
	fmt.Println(x, y)

}