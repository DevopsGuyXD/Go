package main

import (
	"fmt"
	"strings"
)

func getinitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s," ")

	var initials []string

	for _, v := range names{
		initials = append(initials, v[:1])
	}

	if len(initials) > 1{
		return initials[0], initials[1]
	}else{
		return initials[0], "_"
	}
}

func main() {
	fn1,sn1 := getinitials("tifa lockhat")
	fmt.Println(fn1, sn1)

	fn2,sn2 := getinitials("cloud Strife")
	fmt.Println(fn2, sn2)

	fn3,sn3 := getinitials("barret")
	fmt.Println(fn3, sn3)
}