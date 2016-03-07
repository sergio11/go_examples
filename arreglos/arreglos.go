package main

import (
	"fmt"
)

func main() {

	var arreglo [10]int;
	//expresión para declaración e inicialización de un arreglo
	arreglo2 := [3]int{1,2,3}
	//expresión para declarar e inicializar arreglo de dos dimensiones
	arreglo3 := [3][3]int{{4,5},{7,8,2},{6,5,3}}

	fmt.Println("********** Arreglo 1 ********");
	fmt.Println(arreglo);
	fmt.Println("********** Arreglo 2 ********");
	for i := 0; i < len(arreglo2); i++ {
		fmt.Printf("arreglo2[%d] =>  %d \n",i,arreglo2[i]);
	}
	fmt.Println("********** Arreglo 2 ********");
	for i := 0; i < len(arreglo3); i++ {
		for j := 0; j < len(arreglo3[i]); j++ {
			fmt.Printf("arreglo3[%d][%d] =>  %d \n",i,j,arreglo3[i][j]);
		}
	}

	
}