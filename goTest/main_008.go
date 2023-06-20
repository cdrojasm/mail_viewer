package main
import (
	"fmt"
)

var a int
type dinero int
var b dinero 

func main() {
	/* 
	creacion de tipos de datos =>  composicion
	go es un lenguaje con tipos static, eso significa que las varibales una vez declaradas
	solamente puede almacenar valores de ese tipo, EL PUTO TIPO NO ES MUTABLE, ENTIENDELO MARYCARMEN.

	para crear un tipo utilizamos la keyword typeluego colocamos el nombre del tipo y le asignamos el tipo
	del que hereda. Se tienen las mismas limitaciones frente a la inmutabilidad, esto debido a que se debe
	respetar el tipo subyacente o padre del que se hereda.

	Type is life perras.... bill K.

	*/
	b = 1000
	a = 392
	fmt.Printf("%T es el tipo de la variable a con valor %d \n", a, a)
	fmt.Printf("%T es el tipo de la variable a con valor %d \n", b, b)
	/* a = b -> tipos incompatibles en la asignacion de esta expresion*/
	fmt.Printf("%T es el tipo de la variable a con valor %d \n", a, a)
	

}