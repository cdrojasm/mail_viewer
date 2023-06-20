package main

import (
	"fmt"
)

/* package level scope */

var z = 41

func main(){
	/* 
	body function level scope
	palabra clave var es usada para declarar variables con scope
	de paquete. Esto quiere decir que puede ser accedida desde
	cualquier funcion o punto desde el paquete. var puede emplearse
	dentro de el cuerpo de la funcion como en el scope general..

	*/
	var w int
	x := 42 + 7
	y := "James Bond"
	fmt.Println(x)
	fmt.Println(y)
	x = 50
	fmt.Println(x)
	fmt.Println(z)
	numero()
}

func numero(){
	fmt.Println(z)
}