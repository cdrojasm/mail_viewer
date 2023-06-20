package main

import "fmt"

type myType int

func main() {
	/*
		Tipos booleanos
		it can take only two values, true or false. Due this, is called binary
		This type is so used due to comparison capabilities through comparators operators
		== > < != <= >= ......
	*/

	x := 7
	y := 42
	fmt.Println(x == y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)

}
