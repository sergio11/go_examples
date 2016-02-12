package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Introduce tu nombre: ")
	name,err := reader.ReadString('\n')
	if err == nil{
		name = strings.TrimSpace(name);
		if len(name) > 0 {
			fmt.Println("Introduce tu edad: ")
			if edad,err := reader.ReadString('\n'); err == nil{
				fmt.Printf("Tu nombre es %s y tienes %s años",name,edad)
			}
			
		}else{
			fmt.Println("Debes introducir una cadena")
		}

	}else{
		fmt.Println("Error al introducir datos !!!")
	}
}