package main

import (
	"fmt"
)

/* package level scope */
var a int
var b string = "programa"
var c bool 
var d bool = true


func main(){
	/* 
	paquete fmt
	"" -> string literal interpretado
	`` -> string literal row 
	emplea funcionalidades de lectura de archivos, respuesta a solicitudes y en cualquier caso 
	func de tipo IO

	println agrega un \n al final de la cadena de texto 
	printf toma una cadena de texto que define el formato de la cadena de texto que va a la 
		salida de texto estandar. Se emplean verbos en la cadena que define el formato, tambien 
		se conoce como especificador de formato. i.e. verbos
		%v: verbo que especifica el formato por defecto 
		%#v: representacion del valor en go sintax
		%T: representacion del tipo en sintaxis de go
		%%: imprime el simbolo de %
		
	Sprint retorna todos los valores como un string, recibe un conjunto de parametros variable
	las funciones con el prefijo F del paquete F son empleadas para enviar y recibir data de archivos 
	y request a servidor. 
	print 
	*/
	e := 42
	f := "hola mundo de perros"
	g := `El doctor perro dice que comer vegetales es 
		saludable para los otros perros.`
	fmt.Println("Hello, playground")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("Mi",b, f)
	fmt.Println(g)
	a = 392
	fmt.Printf("%T es el tipo de la variable a con valor %d \n", a, a)
	fmt.Printf("%T es el tipo del valor de retorno de la expresion a+b con valor %d \n", a+e, a+e)

	s1 := fmt.Sprint("El programa \"" , f, "\" es una implemetacion basica." )
	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T", s1)

}
