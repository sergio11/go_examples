package main

import (
	"fmt"
)

type User interface{
	Nombre() string
	Permisos() int //1 - 5
}

type Admin struct{
	name string
}

func (this Admin) Permisos() int{
	return 5
}

func (this Admin) Nombre() string{
	return this.name
}

type Editor struct{
	name string
}

func (this Editor) Permisos() int{
	return 4
}

func (this Editor) Nombre() string{
	return this.name
}

func isAdmin(user User) bool{
	auth := false
	if user.Permisos() == 5 {
		auth = true
	}
	return auth
}


func main() {
	users := []User{Admin{"David"},Editor{"Fulano"}}

	for _,user := range users {
		if isAdmin(user) {
			fmt.Printf(" %s tiene permisos de administrador \n",user.Nombre())
		}else{
			fmt.Printf(" %s no tiene permisos de administrador \n",user.Nombre())
		}
	}
	
}