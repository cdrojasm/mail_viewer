package main

import "fmt"

func main() {

	/*
		creacion de array
		se emplean principalmente como bloques constructores de slices
		siempre emplear slices
		los arrays son secuencias numeradas

	*/
	var x [10]int
	x[3] = 32
	var y [6]int
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(len(x))
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
