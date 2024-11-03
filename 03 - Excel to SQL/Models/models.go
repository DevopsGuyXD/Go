package models

import (
	"fmt"

	"github.com/DevopsGuyXD/dogtag/utils"
)


func AddAllRecords(keysToUpdate, valuesToUpdate []string) {

	database := utils.InitDatabase()
	statement := `INSERT INTO domains(`

	for i, key := range keysToUpdate {
		if i == 0 {
			statement += key
		} else {
			statement += "," + key
		}
	}
	statement += `) VALUES(`

	for j, value := range valuesToUpdate {
		if j == 0 {
			statement += fmt.Sprintf(`"%v"`, value)
		} else {
			statement += "," + fmt.Sprintf(`"%v"`, value)
		}
	}
	statement += `)`

	createDomainsTable, err := database.Prepare(statement)
	utils.CheckForNil(err)
	createDomainsTable.Exec()
}