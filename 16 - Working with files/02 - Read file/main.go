package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Welcome to the lesson on reading files")

	readFile("./createFileTest.txt")
}

func readFile(fileName string){
	data, err := ioutil.ReadFile(fileName); if err!= nil{
		panic(err)
	}

	fmt.Println(string(data))
}