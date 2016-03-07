package main

import (
	"fmt"
)

func main() {

	var arreglo []int
	arreglo2 := []int{1,2}
	arreglo3 := []int{1,2,3,4,5,6}

	if arreglo == nil {
		fmt.Println("El arreglo está vacío\n")
	}else{
		fmt.Println(arreglo);
	}

	if arreglo2 == nil {
		fmt.Println("El arreglo2 está vacío\n")
	}else{
		fmt.Printf("El arreglo2 tiene %d elementos \n",len(arreglo2));
	}

	//slice
	// 3 -> Inicio , 5 -> índice final sin incluir 
	slice := arreglo3[3:5]
	fmt.Println("Slice 1 => arreglo3[3:5]")
	fmt.Println(slice)
	slice2 := arreglo3[3:]
	fmt.Println("Slice 2  => arreglo3[3:]")
	fmt.Println(slice2)


	
}