package utils

import (
	"database/sql"
	"log"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xuri/excelize/v2"
)

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitEnvFile(){
	err := godotenv.Load(".env"); CheckForNil(err)
}

func InitDatabase() *sql.DB{
	database, err := sql.Open("sqlite3", "./Sqlite/dogtag.db"); CheckForNil(err)

	return database
}

func OpenExcelFile() *excelize.File{
	f, err := excelize.OpenFile("test.xlsx"); CheckForNil(err)
	return f
}

func ExtractNameFromARNS3() []string{

	var resourceName []string

	f := OpenExcelFile()
	defer f.Close() 

	rows, err := f.GetRows("Update"); CheckForNil(err)
	cols, err := f.GetCols("Update"); CheckForNil(err)
	
	for i, row := range rows[0]{
		if row == "ARN"{
			for j := 0; j < len(rows)-1; j++{
				if strings.Contains(cols[i][j+1], "s3"){
					ARN         := (cols[i][j+1])
					res         := strings.Split(ARN, "arn:aws:s3:::")
					resourceName = append(resourceName, res[1])
				}
			}
		}
	}

	return resourceName
}


