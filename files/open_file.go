package main

import (
	"fmt"
	"os"
	"log"
)

func main() {

	fmt.Println("Abriendo Fichero")
	file , err := os.Open("./file.txt") // For read access.
	if err != nil {
		log.Fatal(err.Error())
	}
	//Stat returns the FileInfo structure describing file. 
	//If there is an error, it will be of type *PathError.
	fi, err := file.Stat()
	if err != nil {
	  // Could not obtain stat, handle error
	}
	size := fi.Size()
	data := make([]byte, size)
	fmt.Printf("The file is %d bytes long\n", size)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes:\n %q \n", count, data[:count])
}