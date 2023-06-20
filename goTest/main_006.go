package main

import (
	"fmt"
)

/* package level scope */
var a int
type dinero int
var b dinero

func main(){
	/* 
	Conversion de tipo
	el lexico en go cambia
	cast -> conversion o asercion
	*/
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}

