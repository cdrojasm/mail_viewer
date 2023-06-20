package main
import (
	"fmt"
)

/* var a int
type dinero int
var b dinero  */

func main() {
	/* 
	Ejercicio Ninja L01 - L03
	1. Usando el operador de declaracion corta de variables ASIGNA los siguientes valores
	a las variables con los identificadores x, y, z.
		x = 42, 
		y = "james Bond" 
		z = true
	2. Luego imprimir los valores almacenados por estas variables usando:
		a. una solo declaracion de la funcion Println
		b. multiples declaraciones de la funcion PrintLn

	*/
	x := 42
	y := "James Bond"
	z := true


	fmt.Println("Primera impresion de las variables x, y, z:")
	fmt.Println(x, y, z)
	fmt.Println("impresiones consecutivas de las variables x, y, z con multiples declaraciones de la funcion Println")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	/* fmt.Printf("%T es el tipo de la variable a con valor %d \n", b, b) */
	

}