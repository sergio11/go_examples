package main

import (

	"fmt"
	"os"
	"log"

)

func main() {

	var offset int64
	saludo := []byte{72,79,76,65,32,83,79,89,32,83,69,82,71,73,79}
	saludo2 := []byte("Hola Mundo!!!")
	saludo3 := "Esto es una cadena"

	file, err_open := os.OpenFile("./file.txt",os.O_RDWR, 0666)
	fi, err_stat := file.Stat()
	if err_open != nil || err_stat != nil {
		log.Fatal(err_open)
		log.Fatal(err_stat)
	}

	fmt.Println("func (*File) Write")
	count, err := file.Write(saludo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Se han escrito %d bytes\n",count)

	fmt.Println("func (*File) WriteAt")
	offset = fi.Size()
	count2, err := file.WriteAt(saludo2,offset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Se han escrito %d bytes",count2)

	fmt.Println("func (*File) WriteString")
	count3, err := file.WriteString(saludo3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Se han escrito %d bytes",count3)
}