package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"

	utils "github.com/DevopsGuyXD/Route53-Checker/Utils"
	"github.com/xuri/excelize/v2"
)

//-------------------------------------------- Add to Excel sheet -------------------------------------------
func AddToExcelSheet(file *excelize.File, sheet int, i int, route53records Route53Records, check_domain_availability CheckDomainAvailability, check_domain_availability_custom string){
	file.SetCellValue("Sheet1","A"+strconv.Itoa(i+2),route53records.HostedZones[i].Name)
	file.SetCellValue("Sheet1","B"+strconv.Itoa(i+2), check_domain_availability.Availability)
	file.SetCellValue("Sheet1","B"+strconv.Itoa(i+2), check_domain_availability_custom)
	file.SetActiveSheet(sheet)
	err := file.SaveAs("./Domain_status.xlsx"); utils.CheckForNil(err)
}

//--------------------------------------- Check if domain is available ---------------------------------------
func CheckIfDomaisAvailable(count RecordCount, route53records Route53Records)(*excelize.File , int, []string){

	fmt.Println("Checking for domain availability \u29D7")
	fmt.Println("")

	var available_domain_name []string
	available_domain_name_total_count := 0

	file := excelize.NewFile()
	sheet, err := file.NewSheet("Sheet1"); utils.CheckForNil(err)
	file.SetCellValue("Sheet1","A1", "Domain name:")
	file.SetCellValue("Sheet1","B1", "Availability:")

	for  i := 0; i < count.HostedZoneCount; i++{

		var check_domain_availability CheckDomainAvailability

		domaincheck, err := exec.Command("aws","route53domains","check-domain-availability","--region", "us-east-1","--domain-name", route53records.HostedZones[i].Name).Output()
		json.Unmarshal(domaincheck, &check_domain_availability)
		
		var err_response string

		if err != nil{
			err_response = "ERROR"
		}else{
			err_response = ""
		}

		fmt.Printf("(%v/%v)", i+1, count.HostedZoneCount )
		fmt.Printf(" %v : %v %v\n",route53records.HostedZones[i].Name ,check_domain_availability.Availability, err_response)

		if check_domain_availability.Availability == "UNAVAILABLE"{
			check_domain_availability_custom := "REGISTERED"
			AddToExcelSheet(file, sheet, i ,route53records, check_domain_availability, check_domain_availability_custom)

			available_domain_name_total_count = available_domain_name_total_count + 1
			available_domain_name = append(available_domain_name, route53records.HostedZones[i].Name)

		}else if check_domain_availability.Availability == "AVAILABLE"{
			check_domain_availability_custom := ""
			AddToExcelSheet(file, sheet, i ,route53records, check_domain_availability, check_domain_availability_custom)

		}else if check_domain_availability.Availability == "PENDING"{
			check_domain_availability_custom := "PENDING"
			AddToExcelSheet(file, sheet, i ,route53records, check_domain_availability, check_domain_availability_custom)

		}else{
			check_domain_availability_custom := "ERROR"
			AddToExcelSheet(file, sheet, i ,route53records, check_domain_availability, check_domain_availability_custom)
		}
	}

	return file, available_domain_name_total_count, available_domain_name
}