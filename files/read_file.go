package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {
	
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fi,err := file.Stat()
    scanner := bufio.NewScanner(file)

    fmt.Println("Leyendo Fichero")
    fmt.Printf("Fichero : %s, Bytes : %d, Mode: %s \n",fi.Name(),fi.Size(), fi.Mode())
    fmt.Println("Contenido")
    /*
		Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method. 
		It returns false when the scan stops, either by reaching the end of the input or an error. 
		After Scan returns false, the Err method will return any error that occurred during scanning, except that if it was io.
		EOF, Err will return nil. 
		Scan panics if the split function returns 100 empty tokens without advancing the input. 
		This is a common error mode for scanners.
    */
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}
