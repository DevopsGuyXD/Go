package main

import (
	"fmt"

	configs "github.com/DevopsGuyXD/dogtag/Configs"
	"github.com/DevopsGuyXD/dogtag/utils"
)

func main() {
	fmt.Println("Welcome to tagging master")

	utils.InitEnvFile()

	configs.CreateTable()
}