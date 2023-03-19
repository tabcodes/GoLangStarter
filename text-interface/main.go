package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No filename supplied at program execution.")
		os.Exit(1)
	}
	fileName := os.Args[1]

	fmt.Println("Filename supplied: ", fileName)

	fileHandler, error := os.Open(fileName)

	if error != nil {
		fmt.Println("Couldn't open", fileName, "for reading- check that the file exists.")
		os.Exit(1)
	}

	b := make([]byte, 999)
	fileHandler.Read(b)
	fmt.Println(string(b))

}
