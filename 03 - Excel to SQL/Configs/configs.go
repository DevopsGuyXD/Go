package configs

import (
	"fmt"
	"strconv"
	"strings"

	models "github.com/DevopsGuyXD/dogtag/Models"
	"github.com/DevopsGuyXD/dogtag/utils"
	_ "github.com/mattn/go-sqlite3"
)


var counter = 2
var keysToUpdate []string


func CreateTable() {

	database := utils.InitDatabase()
	GetKeysToBeUpdated()

	for i, key := range keysToUpdate{
		var createDomainsTableStatement string

		if i == 0{ 
			createDomainsTableStatement = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS domains(
				%v VARCHAR(255)
			)`, key)
		}else{
			createDomainsTableStatement = fmt.Sprintf(`ALTER TABLE domains ADD COLUMN %v VARCHAR(255)`, key)
		}

		createDomainsTable, err := database.Prepare(createDomainsTableStatement); utils.CheckForNil(err)
		createDomainsTable.Exec()
	}

	GetValuesToBeUpdated()
}


func GetKeysToBeUpdated() []string{

	f := utils.OpenExcelFile()
	defer f.Close()

	for i := 'A'; i <= 'Z'; i++ {
		styleID, err := f.GetCellStyle("Update", string(i)+strconv.Itoa(1)); utils.CheckForNil(err)
		style, err := f.GetStyle(styleID); utils.CheckForNil(err)

		if len(style.Fill.Color) < 1{
			continue
		}else{
			if style.Fill.Color[0] != "FFFFFF"{
				res, err := f.GetCellValue("Update", string(i)+strconv.Itoa(1)); utils.CheckForNil(err)
				if strings.Contains(res, "Tag:"){
					res_split := strings.Split(res,"Tag:")
					keysToUpdate = append(keysToUpdate, res_split[1])
				}else{
					keysToUpdate = append(keysToUpdate, res)
				}
			}
		}
	}

	return keysToUpdate
}


func GetValuesToBeUpdated(){

	var valuesToUpdate []string

	f := utils.OpenExcelFile()
	defer f.Close()

	rows, err := f.GetRows("Update"); utils.CheckForNil(err)

	for i := 'A'; i <= 'Z'; i++ {
		styleID, err := f.GetCellStyle("Update", string(i)+strconv.Itoa(1)); utils.CheckForNil(err)
		style, err := f.GetStyle(styleID); utils.CheckForNil(err)

		if len(style.Fill.Color) < 1{
			continue
		}else{
			if style.Fill.Color[0] != "FFFFFF"{
				res2, err := f.GetCellValue("Update", string(i)+strconv.Itoa(counter)); utils.CheckForNil(err)
				valuesToUpdate = append(valuesToUpdate, res2)
			}
		}
	}

	if counter <= len(rows){
		if len(valuesToUpdate) == len(keysToUpdate){
			models.AddAllRecords(keysToUpdate, valuesToUpdate)
		}

		counter += 1
		GetValuesToBeUpdated()
	}
}