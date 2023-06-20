package main

import (
	"fmt"
)

type myType int

func main() {
	/*
		tipo subyacente raiz o implicito
		el tipo subyacente es el tipo sobre el cual se construye el nuevo tipo

		Ejercicio Ninja L04
		crear tipos propios

		1. crear t tipo propio, hacer que el tipo subyacente sea un int
		2. crear una variable utilizando el nuevo tipo usando el identificador x, usando la palabre var.
		3. en el func main
			a. imprimir el valor de la variable x
			b. imprimir el tipo de la variable x
			c. asigna 42 a x usando el operador de asignacion
			d. imprimir el valor de la variable x
	*/
	var x myType

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

}
