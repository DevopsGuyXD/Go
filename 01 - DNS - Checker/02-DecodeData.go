package main

import (
	"encoding/json"
	"fmt"

	utils "github.com/DevopsGuyXD/Route53-Checker/Utils"
)

//------------------------------ Unmarshall DNS records and record count data ------------------------------
func UnmarshalData(route53_list []byte, record_count []byte){

	var count RecordCount
	var route53records Route53Records

	err_record_count := json.Unmarshal(record_count, &count); utils.CheckForNil(err_record_count)

	err_route53_list := json.Unmarshal(route53_list, &route53records); utils.CheckForNil(err_route53_list)

	fmt.Println("Decoding your request \u2714")

	file, available_count, available_domains := CheckIfDomaisAvailable(count, route53records)
	CheckWebSiteStatus(file, available_count, available_domains)
}