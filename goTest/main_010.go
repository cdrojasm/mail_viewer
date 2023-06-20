package main
import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	/* 
	Caraceteres de escape
	\n salto de linea
	\b backspace
	\' single quote
	\" double quote
	\t tab

	Ejercicio Ninja L02
	1. usa var para declarar 3 variables en scope de paquete usar los isguiente identificadores 
		con los siguientes tipos
		a. x identificador con tipo int
		b. identificador y con tipo string
		c. identificador z con tipo bool
	2. en el main:
		a. imprimir los valores de cada identificador
		b. El compilador asigna valores a estas variables. Como se llaman estos valores?

	*/
	
	fmt.Printf("el valor del identificador x es %d, el valor del identificador y es %s y el valor del identificador z es %t\n", x, y, z)
	fmt.Printf(`El valor por defecto que toma un identificador sin asignacion es conocido como valor cero`)
}