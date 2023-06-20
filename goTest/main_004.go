package main

import (
	"fmt"
)

/* package level scope */

var z = 41


func main(){
	/* 
	Tipos de datos
	los datos y sus tipos son estaticos en go, esto acelera el procesamiento ya que no existen 
	operaciones intermedias para determinar si un dato tiene o no un tipo y ademas no 
	existen tampoco las operaciones tipo parse. La asignacion del tipo puede realizarse de forma 
	automatica como en la primera expresion o se puede especificar el tipo como en la 
	segunda expresion:
	Es mala practica utlizar variables con scope a nivel de paquete, en lo posible hacerlo 
	en el scope a nivel de cuerpo de la funcion
	la palabra reservada var solo puede emplearse con el operador de asignacion convencional, no con 
	el corto.

	tipo de datos primitivos:
		tipos de datos basicos: 

		tipos de datos internos:

	tipo de dato compuesto:
		tipo de dato construido a traves de los tipos de datos primitivos.

		construir datos compuestos se conoce como composicion

	*/

	k := 90
	var l int = 93
	fmt.Println(k,l)
	
}

