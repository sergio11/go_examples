package main

import (
	"fmt"
)

func main() {
	// <nil>
	var x,y *int;
	entero := 5
	//x = 5 -> cannot use 5 (type int) as type *int in assignment
	//x = &5 -> cannot take the address of 5

	x = &entero
	y = x
	fmt.Println("x = &entero")
	fmt.Println(x)
	fmt.Println("y = x")
	fmt.Println(y)
	//Cambiamos valor en la dirección de memoria
	*x = 6
	fmt.Println("*x = 6")
	fmt.Println(x)
	fmt.Println(*y)

	//var z *int
	//*z = 7 -> panic: runtime error: invalid memory address or nil pointer dereference 


}