package main

import (
	"fmt"
	"runtime"
)

type myType int

func main() {
	/*
		Tipos numericos
		Go have multiplke representations to use according to the tyupe of number to represent. It can be
		represented floats, integers and others.
		accordign to wordsize, go choose was is the optimal type for a variable.

		numeric Types::::
		uint8			(0,255)
		uint16			(0,65535)
		uint32
		uint64
		int8
		int16
		int32
		int64
		float32
		float64
		complex64
		complex128
		byte 	-> alias for uint8
		rune -> alias for int32
		numerical boundaries should be followed, in another case, compiler will generate an error
		due to invalid assignation
		uintptr an unsigned integer large to enough to store the uninterpreted bits of a pointer value

		numerical assignations on variable will bbehave as a constant after assignation.

	*/

	x := 42
	y := 42.5
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(x <= int(y))
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
