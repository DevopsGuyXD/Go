package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the lesson on checking for NIL errors")
	WriteToFile("./Testfile", "This is a test file that needs to be written")
}

func WriteToFile(fileName string, content string){
	file ,err := os.Create(fileName)
	checkNilError(err)

	length, err := io.WriteString(file, content)

	fmt.Printf(`File written successfully`)
	fmt.Printf(`Total characters written: %v`, length)

	defer file.Close()
}

func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}

// Note: You can also use error.New("This is a custom error message") to generate custom error messages.
// Do not use fmt.Println() for this