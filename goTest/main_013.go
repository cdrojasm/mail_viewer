package main

import (
	"fmt"
)

type myType int

func main() {
	/*
		Ejercicio Ninja L05
		crear tipos propios 2
		A partir del codigo del ejercicio anterior.
		1. a nivel de paquete usando la palabra clave var crear una variable con el identificador y.
			La variable debe ser del mismo tipo subyacente que el tipo subyacente utilizado para crear
			el nuevo tipo empleado por de x.

		2. en el func main
			a. lo mismo del ejercicio anterior
			b. usar una conversion para convertir el tipo del valor almacenado en x al valor implicito
			c. usar el operador = para saignar el valor a y
	*/
	var x myType
	var y int
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", y)

}
