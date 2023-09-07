package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	res, err := exec.Command("nslookup", "myip.opendns.com", "resolver1.opendns.com").Output(); if err != nil{
		fmt.Println(err)
	}

	res1 := string(res)
	res2 := strings.Split(res1, " ")

	byte := []byte(res2[10])

	res3 := bytes.TrimSpace(byte)

	fmt.Println(string(res3))
}