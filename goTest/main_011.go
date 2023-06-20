package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	/*
		Ejercicio Ninja L03
		Usando el codigo del ejercicio anterior:
		1. asignar los siguientes valores a las variables
			a. x asignar 42
			b. y asignar James Bond
			c. asignarle true
		2. en el main:
			a. usar fmt.Sprintg para imprimir todos los valores en un solo string. Asignar
			el valor retornado de tipo string usando el operador de declaracion corta a la
			variable con el identificador s
			b. Imprimir el valor almacenado por la variable s
	*/
	s := fmt.Sprintf("%d\t%s\t%t", x, y, z)
	fmt.Println(s)

}
