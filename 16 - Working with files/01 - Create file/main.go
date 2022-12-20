package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This is content that needs to be put in the test file"

	file, err := os.Create("./createFileTest.txt"); if err != nil{
		panic(err)
	}

	length, err := io.WriteString(file, content); if err != nil{
		panic(err)
	}

	fmt.Println(length)
	defer file.Close()
}