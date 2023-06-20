package main

import (
	"fmt"
)

type myType int

func main() {
	/*
		Ejercicio Ninja L06
		Quiz
		1. El elemento mas pequenio de un programa que expresa alguna accion a ser ejecutada
			a. declaracion
		2. Combinacion de uno o mas valores explicitos, constantes, variables y funciones que el
			lenguaje de programacion interpreta y computa para generar otro valor. i.i 2+3 = 5
			a. expresion
		3. Cual de estos es parentesis
			b. ()
		4. Cual de estos son llaves?
			b. {}
		5. Cual de estos son corchetes?
			a. []
		6. El scope de una variable se refiere a que parte del codigo puede tener acceso a esa
		variable, por ejemplo, escribir o leer sobre o desde ella:
			a. verdadero
		7. Un tipo de dato primitivo es uno que es construiido internamenteen el lenguaje o tipo de
			dato Basico que es construido internamente en el lenguaje
			a. verdadero
		8. En Go, un int no es un tipo de dato primitivo
			a. falso
		9. En go, un string es un tipo de dato primitivo
			a. verdadero
		10. En go, un tipo de dato compuesto te permite componer valores a partir de otros
			tipos de datos.
			a. verdadero
		11. Cuando una variable es declarada en Ho usando la palabra clave "Var" y ningun valor se le es
			asignado, el compilador le asigna un valor por defecto a esa variable. Esto es conocido
			como el valor cero.
			a. verdadero
		12. Las palabras claves son palabras reservadas por el lenguaje Go, tienen que ser usadas de cierta
			manera para cierto proposito especifico
			a. verdadero
		13. No puedes usar una palabra clave para algo diferente que para lo qeu esta hecha. Es decir
			darle un proposito diferente.
			a. verdadero
		14. En 2 + 2 el + es el operador
			a. verdadero
		15. en 2 + 2 los 2 son los operandos
			a. verdadero
		16. Para leer la doc, cual es la diferencia entre la documentacion encontrada en golang.org y en
			godoc.org
			a. una es la doc oficial golang.org , luego estan la documentacion de la comunidad con
				la lista de paquetes de la comunidad
		17. package no es una palabra reservada
			b. falso
		18. var es una palabra reservada
			a. verdadero
		19. el punto de entrada para todos los programas en en func main() el cual necesita estar dentro
			del paquete main.
			a. verdadero
		20. El operador de declaracion corta puede ser usado en cualquier parte del programa, incluyendo ambos,
			nivel de paquete y a nivel de bloque.
			b. falso
		21. Cuales son las 3 palabras usadas para describir buenos nombres de paquetes, buscar en la doc
			effective go
			Descriptivo, corto, conciso
		22. Cual es el nombre del sitio web donde puedes escribir codigo en linea
			playground
		23. Uno de los mejores lugares para hacer preguntas es el golang Bridge
			a. verdadero
		24. Cuando ves algo como fmt.Println se esa llamado a la funcion PrintLn del paquete fmt
			a. verdadero
		25. Un identificador es el nombre asignado a una variable, func o constante
			a. verdadero
		26. Para llamar a una funcion, var, o const de un paquete se utiliza la sintaxis paquete-punto-ident
			a. verdadero
		27. Que es codigo de go idiomatico
			a. codigo de go escrito con los estandares de la comunidad
		28. Cual caracter te permite invalidad o anular o enviar al vacio dicho de otra manera, cual caracter
			te permite decirle al compilador que no va a usar el valor retornado por una funcion


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
