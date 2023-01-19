package main

// ---------------------------------- Imports
import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"

	utils "github.com/DevopsGuyXD/Route53-Checker/Utils"
	"github.com/xuri/excelize/v2"
)

// ---------------------------------- END: Imports

// --------------------------------- Record count
type RecordCount struct{
	HostedZoneCount int `json:"HostedZoneCount,omitempty"`
}
// --------------------------------- END: Record count


// --------------------------------- Route53 records
type Route53Records struct {
	HostedZones []HostedZones `json:"HostedZones,omitempty"`
}

type HostedZones struct {
	Id                     string  `json:"Id,omitempty"`
	Name                   string  `json:"Name,omitempty"`
	CallerReference        string  `json:"CallerReference,omitempty"`
	Config                 *Config `json:"Config,omitempty"`
	ResourceRecordSetCount int     `json:"ResourceRecordSetCount,omitempty"`
}

type Config struct {
	PrivateZone bool `json:"PrivateZone,omitempty"`
}

type CheckDomainAvailability struct{
	Availability string `json:"Availability,omitempty"`
}
// --------------------------------- END: Route53 records


// --------------------------------- Main function
func main() {

	fmt.Println("Welcome to Domain-checker")
	fmt.Println("")

	GetRoute53RecordDetails()
}
// --------------------------------- END: Main function


//---------------------------------- Get DNS records and record count
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
//---------------------------------- END: Get DNS records and record count


//---------------------------------- Unmarshall DNS records and record count data sent back at the response
func UnmarshalData(route53_list []byte, record_count []byte){

	var count RecordCount
	var route53records Route53Records

	err_record_count := json.Unmarshal(record_count, &count); utils.CheckForNil(err_record_count)

	err_route53_list := json.Unmarshal(route53_list, &route53records); utils.CheckForNil(err_route53_list)

	fmt.Println("Decoding your request \u2714")

	CheckIfDomaisAvailable(count, route53records)
	//CheckDomainResponse(count, route53records)
}
//----------------------------------- END: Unmarshall DNS records and record count data sent back at the response


//----------------------------------- Check if domain is available
func CheckIfDomaisAvailable(count RecordCount, route53records Route53Records){

	fmt.Println("Checking for domain availability \u29D7")
	fmt.Println("")
	defer fmt.Println("Check in dir 'Domain_Availability_Output'")
	defer fmt.Println("Completed successfully \u2B50")
	defer fmt.Println("")

	// unavailable, err := os.Create("Domain_Availability_Output/unavailable.xlsx"); utils.CheckForNil(err)
	// available, err := os.Create("Domain_Availability_Output/available.xlsx"); utils.CheckForNil(err)
	// pending, err := os.Create("Domain_Availability_Output/pending.xlsx"); utils.CheckForNil(err)
	// error, err := os.Create("Domain_Availability_Output/error.xlsx"); utils.CheckForNil(err)

	file := excelize.NewFile()
	sheet, err := file.NewSheet("Sheet1"); utils.CheckForNil(err)

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

			file.SetCellValue("Sheet1","A"+strconv.Itoa(i+1),route53records.HostedZones[i].Name)
			file.SetCellValue("Sheet1","B"+strconv.Itoa(i+1), check_domain_availability.Availability)
			file.SetActiveSheet(sheet)
			err = file.SaveAs("./Domain_Availability_Output/Domain_status.xlsx"); utils.CheckForNil(err)

		}else if check_domain_availability.Availability == "AVAILABLE"{

			file.SetCellValue("Sheet1","A"+strconv.Itoa(i+1),route53records.HostedZones[i].Name)
			file.SetCellValue("Sheet1","B"+strconv.Itoa(i+1),check_domain_availability.Availability)
			file.SetActiveSheet(sheet)
			err = file.SaveAs("./Domain_Availability_Output/Domain_status.xlsx"); utils.CheckForNil(err)

		}else if check_domain_availability.Availability == "PENDING"{

			file.SetCellValue("Sheet1","A"+strconv.Itoa(i+1),route53records.HostedZones[i].Name)
			file.SetCellValue("Sheet1","B"+strconv.Itoa(i+1),check_domain_availability.Availability)
			file.SetActiveSheet(sheet)
			err = file.SaveAs("./Domain_Availability_Output/Domain_status.xlsx"); utils.CheckForNil(err)

		}else{

			file.SetCellValue("Sheet1","A"+strconv.Itoa(i+1),route53records.HostedZones[i].Name)
			file.SetCellValue("Sheet1","B"+strconv.Itoa(i+1),"ERROR")
			file.SetActiveSheet(sheet)
			err = file.SaveAs("./Domain_Availability_Output/Domain_status.xlsx"); utils.CheckForNil(err)
		}
	}
}
//----------------------------------- End: Check if domain si available


// // ---------------------------------- Check domain response
// func CheckDomainResponse(count RecordCount, route53records Route53Records){

// 	fmt.Println("")
// 	fmt.Println("6. Check domain response \u29D7")
// 	fmt.Println("")


// 	for i := 0; i < count.HostedZoneCount; i++{
// 		protocal_http := "http://"
// 		//protocal_https := "https://"

// 		add_protocal_http := protocal_http +  route53records.HostedZones[i].Name
// 		//add_protocal_https := protocal_https +  route53records.HostedZones[i].Name

// 		http_res, _ := http.Get(add_protocal_http)

		
// 		fmt.Println(http_res.StatusCode)
// 	}
// }
// // ---------------------------------- END: Check domain response