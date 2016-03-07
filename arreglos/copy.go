package main

import (
	"fmt"
)

func main() {

	source := []int{1,2,3,4,5,6}
	target := make([]int,4,8)

	fmt.Println("Arreglo Fuente")
	fmt.Println(source)
	fmt.Println("Arreglo Destino (antes de la copia)")
	fmt.Println(target)

	copy(target,source)
	fmt.Println("Arreglo Destino (después de la copia)")
	fmt.Println(target)
	
}