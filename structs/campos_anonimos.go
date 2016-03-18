package main

import (
	"fmt"
)

type Human struct{
	name, firstname string
	age int
}

func (this *Human) getName() string{
	return this.name;
}

func (this *Human) getFullName() string{
	return this.name + " " + this.firstname;
}

func (this *Human) getAge() int{
	return this.age;
}

type Teacher struct{
	Human
	level string
}

func (this *Teacher) teach(){
	fmt.Println("Enseñando ....")
}

func (this *Teacher) getLevel() string{
	return this.level;
}

type Tutor struct{
	Teacher //anonymous field
}


func main() {
	david := Tutor{Teacher{Human{"David","Martín",34},"Diplomado"}}

	fmt.Printf("Hola soy %s y tengo %d años \n",david.getFullName(), david.getAge());
	david.teach();
}