package main

import (
	"fmt"
	"net/http"
	"strconv"

	utils "github.com/DevopsGuyXD/Route53-Checker/Utils"
	"github.com/xuri/excelize/v2"
)

// ------------------------------------- HTTP Method ----------------------------------
func MakeDomainNamesAbsoluteHTTP(domain string) (string){
	return "http://" + domain
}

// ------------------------------------ HTTPS Method ----------------------------------
func MakeDomainNamesAbsoluteHTTPS(domain string) (string){
	return "https://" + domain
}

// --------------------------------- Excel sheet heading ------------------------------
func ExcelSheetsHeading(file *excelize.File, sheet_name string, sheet_title string){
	file.SetCellValue(sheet_name, "A1", sheet_title)
	file.SetCellValue(sheet_name, "B1", "HTTP:")
	file.SetCellValue(sheet_name, "C1", "HTTPS:")
}

// --------------------------------- HTTP sheet Function ------------------------------
func AddToExcelSheetHTTP(file *excelize.File, sheet_name string, sheet int, i int, custom_domain string, available_domains []string, status_response string){

	file.SetCellValue(sheet_name,"A"+strconv.Itoa(i+2), custom_domain + available_domains[i])
	file.SetCellValue(sheet_name,"B"+strconv.Itoa(i+2), status_response)
	file.SetActiveSheet(sheet)
	err := file.SaveAs("./Domain_status.xlsx"); utils.CheckForNil(err)

	//fmt.Printf("%v\n",MakeDomainNamesAbsoluteHTTP(available_domains[i]))
}

// -------------------------------- HTTPS sheet Function -----------------------------
func AddToExcelSheetHTTPS(file *excelize.File, sheet_name string, sheet int, i int, available_domains []string, status_response string){

	file.SetCellValue(sheet_name,"C"+strconv.Itoa(i+2), status_response)
	file.SetActiveSheet(sheet)
	err := file.SaveAs("./Domain_status.xlsx"); utils.CheckForNil(err)

	//fmt.Printf("%v\n",MakeDomainNamesAbsoluteHTTP(available_domains[i]))
}

// -------------------------------- Site response checker -----------------------------
func CheckWebSiteStatus(file *excelize.File, available_count int, available_domains []string) {

	fmt.Println("")
	fmt.Println("Checking status of registered domains...")
	fmt.Println("")

	naked_domain_sheet_name := "Naked domain status"
	custom_domain_sheet_name := "Custom domain status"

	sheet_naked, err := file.NewSheet("Naked domain status"); utils.CheckForNil(err)
	ExcelSheetsHeading(file, naked_domain_sheet_name, "Naked domain:")

	sheet_custom, err := file.NewSheet("Custom domain status"); utils.CheckForNil(err)
	ExcelSheetsHeading(file, custom_domain_sheet_name,"Custom domain:")

	if available_count > 0{
		for i := 0 ; i < available_count; i++{
			
			//------------------------------------------------------ HTTP NAKED -----------------------------------------------------
			res_http_naked, err := http.Get(MakeDomainNamesAbsoluteHTTP(available_domains[i])); if err != nil{
				custom_domain := ""
				AddToExcelSheetHTTP(file, naked_domain_sheet_name, sheet_naked, i, custom_domain, available_domains  ,"ERROR")

			}else{
				status_response := strconv.Itoa(res_http_naked.StatusCode)
				custom_domain := ""
				AddToExcelSheetHTTP(file, naked_domain_sheet_name, sheet_naked, i, custom_domain, available_domains, status_response)
			}

			//------------------------------------------------------ HTTP CUSTOM ------------------------------------------------------
			res_http_custom, err := http.Get(MakeDomainNamesAbsoluteHTTP("www." + available_domains[i])); if err != nil{
				custom_domain := "www."
				AddToExcelSheetHTTP(file, custom_domain_sheet_name, sheet_custom, i, custom_domain, available_domains  ,"ERROR")

			}else{
				status_response := strconv.Itoa(res_http_custom.StatusCode)
				custom_domain := "www."
				AddToExcelSheetHTTP(file, custom_domain_sheet_name, sheet_custom, i, custom_domain, available_domains, status_response)
			}

			//------------------------------------------------------ HTTPS NAKED -----------------------------------------------------
			res_https_naked, err := http.Get(MakeDomainNamesAbsoluteHTTPS(available_domains[i])); if err != nil{
				AddToExcelSheetHTTPS(file, naked_domain_sheet_name, sheet_naked, i, available_domains  ,"ERROR")

			}else{
				status_response := strconv.Itoa(res_https_naked.StatusCode)
				AddToExcelSheetHTTPS(file, naked_domain_sheet_name, sheet_naked, i, available_domains, status_response)
			}
			
			//------------------------------------------------------ HTTPS CUSTOM -----------------------------------------------------
			res_https_custom, err := http.Get(MakeDomainNamesAbsoluteHTTPS(available_domains[i])); if err != nil{
				AddToExcelSheetHTTPS(file, custom_domain_sheet_name, sheet_naked, i, available_domains  ,"ERROR")

			}else{
				status_response := strconv.Itoa(res_https_custom.StatusCode)
				AddToExcelSheetHTTPS(file, custom_domain_sheet_name, sheet_naked, i, available_domains, status_response)
			}

		}
	}else{
		fmt.Println("No registered records to check")
	}

	fmt.Println("")
	fmt.Println("Completed successfully \u2B50")
	fmt.Println("Check in dir 'Domain_Availability_Output'")
	//----------------------------------------------------------- FINISH ----------------------------------------------------------------
}