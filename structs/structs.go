package main

import (
	"fmt"
)


type User struct{
	id,age int
	name,firstname string
}

func main() {

	var juan User
	pedro := User{id:12345, name: "Pedro", firstname: "Martín", age:40}
	marcos := new(User)
	(*marcos).name = "Marcos"
	marcos.age = 29
	fmt.Println(juan)
	fmt.Println(pedro)
	fmt.Println(*marcos)
	
}