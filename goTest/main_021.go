package main

import "fmt"

func main() {

	/*
		a composite datatype is any type of datatype that can be builded using
		the exissent datatypes in the language.
		The act of build a composed type is also knowed as composittion

		composite literals

		slices has dynamic len


	*/
	// tipo{elementos}
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	for i, v := range x {
		fmt.Println(i, v)
	}
}
