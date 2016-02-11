package main;

import (
	"fmt"
	"strings"
)

func main() {
	var name string;
	var age int;

	fmt.Printf("Escriba su nombre: ");
	fmt.Scanf("%s\n",&name);
	name = strings.TrimSpace(name);
	if len(name) > 0 {
		fmt.Printf("Escriba %s dime tu edad: ",name);
		fmt.Scanf("%d\n",&age);
		if age <= 0 || age >= 100{
			fmt.Printf("La edad especificada no es correcta");
		}else{
			fmt.Printf("Te llamas %s y tienes %d años\n",name,age);
		}
	}else{
		fmt.Printf("Debes especificar una cadena");
	}
}
