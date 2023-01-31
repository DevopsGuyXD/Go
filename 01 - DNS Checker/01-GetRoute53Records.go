package main

import (
	"fmt"
	"log"
	"os/exec"

	utils "github.com/DevopsGuyXD/Route53-Checker/Utils"
)

//------------------------------------ Get DNS records and record count --------------------------------
func GetRoute53RecordDetails(){

	count, err := exec.Command("aws","route53","get-hosted-zone-count").Output(); utils.CheckForNil(err)
	fmt.Println("Getting Route53 record count \u2714")

	res, err := exec.Command("aws","route53","list-hosted-zones").Output(); utils.CheckForNil(err)
	fmt.Println("Getting Route53 records \u2714")

	if !utils.CheckIfValidJson(res){
		log.Fatal("Invalid Json")
	}

	UnmarshalData(res, count)
}