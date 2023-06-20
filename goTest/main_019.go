package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

func main() {

	/*
		IOTA
		const
		IOTA is a keyword that permits make sequenced declaration where data
		will take a value with a increment of 1 unit.

	*/
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
