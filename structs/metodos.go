package main

import (
	"fmt"
	"strconv"
)

type User struct {
	age int
	name,firstname string
}

//Métodos

func (this *User) setName(name string){
	this.name = name
}

func (this *User) setAge(age int){

	if age > 0 && age < 100 {
       this.age = age
	}else{
	 	fmt.Printf("Age %d debe estar entre 0 y 100\n",age)
	}
	
}

func (this *User) setFirstName(firstname string){
	this.firstname = firstname
}

func (this *User) fullInfo() string{
	return this.name + " - " + this.firstname + " - " + strconv.Itoa(this.age)
}

func main() {

	alumno := new(User)

	alumno.name = "Jose"
	alumno.firstname = "Martín"
	alumno.age = 28

	fmt.Println(alumno.fullInfo())

	alumno.setName("Pedro")
	alumno.setFirstName("Fidalgo")
	alumno.setAge(222)

	fmt.Println(alumno.fullInfo())
}		