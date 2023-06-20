package main

import (
	"fmt"
)

func main(){
	/* se emplea _ en las variables para indicar el no 
	uso del retorno de una funcion
	*/
	nb, _ := fmt.Println("hola mundo perros")
	fmt.Println("la cantidad de bytes escritos fue", nb)
}