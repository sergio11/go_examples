package main

import (
	"fmt"
)

func main() {
	//Arreglo1 longitud 3, capacidad 5
	slice := make([]int,3,5)
	fmt.Printf("Capacidad %d , Longitud %d \n", cap(slice) ,len(slice))
	fmt.Println(slice)

	//No compila
	//slice2 := make([]string,7,3)
	//len larger than cap in make([]string)

	slice = append(slice,45)
	fmt.Printf("Capacidad %d , Longitud %d \n", cap(slice) ,len(slice))
	fmt.Println(slice)

	/*
		Si se sobrepasa la capacidad del arreglo, se crea otro con el doble 
		de su capacidad inicial.

		Capacidad 10 , Longitud 6
		[0 0 0 45 45 45]
	*/


}