package main

import "fmt"

type myType int

func main() {

	/*
		Tipos string
		An string is a possibly empty bytes sequence. Strings are not mutables.
		An string
		slice y cliclo for en strings
		inmutability refers to the block in memory of initial position
		of memory reservation.
		strings are read_only, they can not be modified, but it can be reassigned
		string literal interpretado reconoce automaticamente los saltos de linea
		escritos en el editor.

	*/
	s1 := "hola mundo de perros"
	s2 := `Este es un string declarado con 
	row string.`
	fmt.Printf("El tipo de S1 es: %T\n", s1)
	fmt.Printf("El tipo de S2 es: %T\n", s2)
	fmt.Println(s1)
	fmt.Println(s2)
	bs := []byte(s1) /* convertion from string to bytes slice */
	fmt.Printf("\n")
	fmt.Printf("%T\n", bs)
	fmt.Printf("%v\n", bs)

}
